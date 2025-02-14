// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package frontend

import (
	"context"
	"io"

	"github.com/MemeLabs/protobuf/pkg/rpc"
	"github.com/MemeLabs/strims/internal/app"
	videov1 "github.com/MemeLabs/strims/pkg/apis/video/v1"
	"github.com/MemeLabs/strims/pkg/ppspp"
	"github.com/MemeLabs/strims/pkg/ppspp/store"
	"go.uber.org/zap"
)

func init() {
	RegisterService(func(server *rpc.Server, params ServiceParams) {
		videov1.RegisterEgressService(server, &videoEgressService{
			logger: params.Logger,
			app:    params.App,
		})
	})
}

// videoEgressService ...
type videoEgressService struct {
	logger *zap.Logger
	app    app.Control
}

// OpenStream ...
func (s *videoEgressService) OpenStream(ctx context.Context, r *videov1.EgressOpenStreamRequest) (<-chan *videov1.EgressOpenStreamResponse, error) {
	uri, err := ppspp.ParseURI(r.SwarmUri)
	if err != nil {
		return nil, err
	}
	logger := s.logger.With(zap.Stringer("swarm", uri.ID))

	ch := make(chan *videov1.EgressOpenStreamResponse)
	go func() {
		defer close(ch)

		transferID, r, err := s.app.VideoEgress().OpenStream(ctx, r.SwarmUri, r.NetworkKeys)
		if err != nil {
			logger.Debug("opening stream failed", zap.Error(err))

			ch <- &videov1.EgressOpenStreamResponse{
				Body: &videov1.EgressOpenStreamResponse_Error_{
					Error: &videov1.EgressOpenStreamResponse_Error{
						Message: err.Error(),
					},
				},
			}
			return
		}

		go func() {
			<-ctx.Done()
			r.Close()
		}()

		ch <- &videov1.EgressOpenStreamResponse{
			Body: &videov1.EgressOpenStreamResponse_Open_{
				Open: &videov1.EgressOpenStreamResponse_Open{
					TransferId: transferID[:],
				},
			},
		}

		var i int
		var resPool [3]*videov1.EgressOpenStreamResponse
		for i := range resPool {
			resPool[i] = &videov1.EgressOpenStreamResponse{
				Body: &videov1.EgressOpenStreamResponse_Data_{
					Data: &videov1.EgressOpenStreamResponse_Data{
						Data: make([]byte, 32*1024),
					},
				},
			}
		}
		for {
			if i++; i == len(resPool) {
				i = 0
			}
			res := resPool[i]

			d := res.GetData()
			d.SegmentEnd = false
			d.Discontinuity = false
			d.Data = d.Data[:cap(d.Data)]

			var n int
		ReadLoop:
			for n < len(d.Data) {
				nn, err := r.Read(d.Data[n:])
				n += nn

				switch err {
				case nil:
				case io.EOF:
					d.SegmentEnd = true
					break ReadLoop
				case store.ErrStreamReset:
					fallthrough
				case store.ErrBufferUnderrun:
					d.Discontinuity = true
					n = 0
				default:
					logger.Debug("stream closed", zap.Error(err))

					ch <- &videov1.EgressOpenStreamResponse{
						Body: &videov1.EgressOpenStreamResponse_Error_{
							Error: &videov1.EgressOpenStreamResponse_Error{
								Message: err.Error(),
							},
						},
					}
					return
				}
			}

			d.Data = d.Data[:n]
			ch <- res
		}
	}()
	return ch, nil
}
