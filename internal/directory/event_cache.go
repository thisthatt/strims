// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package directory

import (
	"sync"

	networkv1 "github.com/MemeLabs/strims/pkg/apis/network/v1"
	networkv1directory "github.com/MemeLabs/strims/pkg/apis/network/v1/directory"
)

func newEventCache(network *networkv1.Network) *eventCache {
	return &eventCache{
		Network: network,

		listingChangeEvents:      map[uint64]*networkv1directory.Event_ListingChange{},
		userCountChangeEvents:    map[uint64]*networkv1directory.Event_UserCountChange{},
		userPresenceChangeEvents: map[uint64]*networkv1directory.Event_UserPresenceChange{},
	}
}

type eventCache struct {
	Network *networkv1.Network

	mu                       sync.Mutex
	listingChangeEvents      map[uint64]*networkv1directory.Event_ListingChange
	userCountChangeEvents    map[uint64]*networkv1directory.Event_UserCountChange
	userPresenceChangeEvents map[uint64]*networkv1directory.Event_UserPresenceChange
}

func (d *eventCache) Events() *networkv1directory.EventBroadcast {
	d.mu.Lock()
	defer d.mu.Unlock()

	b := &networkv1directory.EventBroadcast{}
	for _, e := range d.listingChangeEvents {
		b.Events = append(b.Events, &networkv1directory.Event{
			Body: &networkv1directory.Event_ListingChange_{
				ListingChange: e,
			},
		})
	}
	for _, e := range d.userCountChangeEvents {
		b.Events = append(b.Events, &networkv1directory.Event{
			Body: &networkv1directory.Event_UserCountChange_{
				UserCountChange: e,
			},
		})
	}
	for _, e := range d.userPresenceChangeEvents {
		b.Events = append(b.Events, &networkv1directory.Event{
			Body: &networkv1directory.Event_UserPresenceChange_{
				UserPresenceChange: e,
			},
		})
	}
	return b
}

func (d *eventCache) StoreEvent(b *networkv1directory.EventBroadcast) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, e := range b.Events {
		switch b := e.Body.(type) {
		case *networkv1directory.Event_ListingChange_:
			d.handleListingChange(b.ListingChange)
		case *networkv1directory.Event_Unpublish_:
			d.handleUnpublish(b.Unpublish)
		case *networkv1directory.Event_UserCountChange_:
			d.handleUserCountChange(b.UserCountChange)
		case *networkv1directory.Event_UserPresenceChange_:
			d.handleUserPresenceChange(b.UserPresenceChange)
		}
	}
}

func (d *eventCache) handleListingChange(e *networkv1directory.Event_ListingChange) {
	d.listingChangeEvents[e.Id] = e
}

func (d *eventCache) handleUnpublish(e *networkv1directory.Event_Unpublish) {
	delete(d.listingChangeEvents, e.Id)
	delete(d.userCountChangeEvents, e.Id)
}

func (d *eventCache) handleUserCountChange(e *networkv1directory.Event_UserCountChange) {
	d.userCountChangeEvents[e.Id] = e
}

func (d *eventCache) handleUserPresenceChange(e *networkv1directory.Event_UserPresenceChange) {
	if e.Online {
		d.userPresenceChangeEvents[e.Id] = e
	} else {
		delete(d.userPresenceChangeEvents, e.Id)
	}
}
