// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package chat

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/MemeLabs/strims/internal/dao"
	"github.com/MemeLabs/strims/internal/network/dialer"
	chatv1 "github.com/MemeLabs/strims/pkg/apis/chat/v1"
	"github.com/MemeLabs/strims/pkg/kv"
	"github.com/MemeLabs/strims/pkg/logutil"
	"github.com/MemeLabs/strims/pkg/timeutil"
	"go.uber.org/zap"
)

const (
	whisperRateLimitMaxCount = 15
	whisperRateLimitPeriod   = 30 * time.Second
)

var (
	errWhisperRateLimit     = errors.New("whisper rate limit exceeded")
	errWhisperInternalError = errors.New("unable to receive whisper")
)

func newWhisperService(
	logger *zap.Logger,
	store *dao.ProfileStore,
) *whisperService {
	return &whisperService{
		logger: logger,
		store:  store,
		done:   make(chan struct{}),
	}
}

type whisperService struct {
	logger    *zap.Logger
	store     *dao.ProfileStore
	closeOnce sync.Once
	done      chan struct{}
	lock      sync.Mutex
}

func (d *whisperService) SendMessage(ctx context.Context, req *chatv1.WhisperSendMessageRequest) (*chatv1.WhisperSendMessageResponse, error) {
	now := timeutil.Now()
	peerCert := dialer.VPNCertificate(ctx).GetParent()

	if ignore, err := dao.ChatUIConfigIgnoresByPeerKey.Get(d.store, peerCert.Key); err == nil {
		if ignore.Deadline == 0 || ignore.Deadline > now.Unix() {
			return nil, errWhisperInternalError
		}
	} else if !errors.Is(err, kv.ErrRecordNotFound) {
		d.logger.Error(
			"error loading ignore",
			logutil.ByteHex("peer", peerCert.Key),
			zap.Error(err),
		)
		return nil, errWhisperInternalError
	}

	record, err := dao.NewChatWhisperRecord(
		d.store,
		dao.CertificateNetworkKey(peerCert),
		req.ServerKey,
		peerCert.Key,
		peerCert,
		req.Body,
		ExtractMessageEntities(req.Body),
	)
	if err != nil {
		return nil, errWhisperInternalError
	}

	err = d.store.Update(func(tx kv.RWTx) error {
		thread, err := dao.ChatWhisperThreadsByPeerKey.Get(tx, peerCert.Key)
		if err != nil && !errors.Is(err, kv.ErrRecordNotFound) {
			return err
		}
		if thread == nil {
			thread, err = dao.NewChatWhisperThread(d.store, peerCert)
			if err != nil {
				return err
			}
		}

		thread.Alias = peerCert.Subject
		thread.UnreadCount++
		thread.LastReceiveTimes = updateWhisperLastReceiveTimes(thread.LastReceiveTimes, now)
		thread.LastMessageTime = now.UnixNano() / int64(timeutil.Precision)
		thread.LastMessageId = record.Id

		if len(thread.LastReceiveTimes) > whisperRateLimitMaxCount {
			return errWhisperRateLimit
		}

		record.ThreadId = thread.Id
		if err := dao.ChatWhisperRecords.Insert(tx, record); err != nil {
			return err
		}
		return dao.ChatWhisperThreads.Upsert(tx, thread)
	})
	if err != nil {
		return nil, errWhisperInternalError
	}

	return &chatv1.WhisperSendMessageResponse{}, nil
}

func updateWhisperLastReceiveTimes(src []int64, now timeutil.Time) []int64 {
	threshold := now.Add(-whisperRateLimitPeriod).UnixNano() / int64(timeutil.Precision)
	dst := make([]int64, 0, len(src)+1)
	for _, t := range src {
		if t > threshold {
			dst = append(dst, t)
		}
	}
	return append(dst, now.UnixNano()/int64(timeutil.Precision))
}
