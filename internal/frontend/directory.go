package frontend

import (
	"context"

	control "github.com/MemeLabs/go-ppspp/internal"
	"github.com/MemeLabs/go-ppspp/internal/dao"
	"github.com/MemeLabs/go-ppspp/internal/directory"
	"github.com/MemeLabs/go-ppspp/internal/event"
	networkv1directory "github.com/MemeLabs/go-ppspp/pkg/apis/network/v1/directory"
	"github.com/MemeLabs/protobuf/pkg/rpc"
	"golang.org/x/sync/errgroup"
)

func init() {
	RegisterService(func(server *rpc.Server, params *ServiceParams) {
		networkv1directory.RegisterDirectoryFrontendService(server, &directoryService{
			app: params.App,
		})
	})
}

// directoryService ...
type directoryService struct {
	app control.AppControl
}

// Open ...
func (s *directoryService) Open(ctx context.Context, r *networkv1directory.FrontendOpenRequest) (<-chan *networkv1directory.FrontendOpenResponse, error) {
	ch := make(chan *networkv1directory.FrontendOpenResponse, 128)

	go func() {
		raw := make(chan interface{}, 8)
		s.app.Events().Notify(raw)

		for {
			select {
			case e := <-raw:
				switch e := e.(type) {
				case event.DirectoryEvent:
					ch <- &networkv1directory.FrontendOpenResponse{
						NetworkId:  e.NetworkID,
						NetworkKey: e.NetworkKey,
						Body: &networkv1directory.FrontendOpenResponse_Broadcast{
							Broadcast: e.Broadcast,
						},
					}
				case event.NetworkStop:
					ch <- &networkv1directory.FrontendOpenResponse{
						NetworkId:  e.Network.Id,
						NetworkKey: dao.CertificateRoot(e.Network.Certificate).Key,
						Body: &networkv1directory.FrontendOpenResponse_Close_{
							Close: &networkv1directory.FrontendOpenResponse_Close{},
						},
					}
				}
			case <-ctx.Done():
				s.app.Events().StopNotifying(raw)
				close(raw)
				close(ch)
				return
			}
		}
	}()

	return ch, nil
}

// Test ...
func (s *directoryService) Test(ctx context.Context, r *networkv1directory.FrontendTestRequest) (*networkv1directory.FrontendTestResponse, error) {
	client, err := s.app.Dialer().Client(r.NetworkKey, r.NetworkKey, directory.AddressSalt)
	if err != nil {
		return nil, err
	}
	directoryClient := networkv1directory.NewDirectoryClient(client)

	sets := []struct {
		service networkv1directory.Listing_Embed_Service
		ids     []string
	}{
		{
			networkv1directory.Listing_Embed_DIRECTORY_LISTING_EMBED_SERVICE_YOUTUBE,
			[]string{"DDTGlyJVNVI", "SocSlBubzwA", "c5Z-0hkyxgo", "-IO6fpjDJY8", "nKrznsPB5t8", "PMesD2l6viA", "GTYJd2qfx5g", "i6zaVYWLTkU", "9VE7afYWzYo", "JQ4Jx8XfP_w", "oEgZeTr3Vs4", "pRpeEdMmmQ0"},
		},
		{
			networkv1directory.Listing_Embed_DIRECTORY_LISTING_EMBED_SERVICE_ANGELTHUMP,
			[]string{"psrngafk", "bliutwo", "t4tv", "somuchforsubtlety", "erik", "windowsmoviehouse", "keyno", "eastmancolor", "feenamabob", "harkdan"},
		},
		{
			networkv1directory.Listing_Embed_DIRECTORY_LISTING_EMBED_SERVICE_TWITCH_STREAM,
			[]string{"namemannen", "not0like0this", "buddha", "shroud", "purgegamers", "sweetdreams", "destiny", "hannapig_", "ibabyrainbow", "tastelesstv"},
		},
		{
			networkv1directory.Listing_Embed_DIRECTORY_LISTING_EMBED_SERVICE_TWITCH_VOD,
			[]string{"1159927987", "1157956746", "1160400711"},
		},
	}

	eg, ctx := errgroup.WithContext(ctx)

	for _, set := range sets {
		for _, id := range set.ids {
			listing := &networkv1directory.Listing{
				Content: &networkv1directory.Listing_Embed_{
					Embed: &networkv1directory.Listing_Embed{
						Service: set.service,
						Id:      id,
					},
				},
			}

			eg.Go(func() error {
				return directoryClient.Publish(
					ctx,
					&networkv1directory.PublishRequest{Listing: listing},
					&networkv1directory.PublishResponse{},
				)
			})
		}
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// listing := &networkv1directory.Listing{
	// 	// Content: &networkv1directory.Listing_Media_{
	// 	// 	Media: &networkv1directory.Listing_Media{
	// 	// 		MimeType: rtmpingress.TranscoderMimeType,
	// 	// 		// SwarmUri:  s.swarm.URI().String(),
	// 	// 	},
	// 	// },
	// 	Content: &networkv1directory.Listing_Embed_{
	// 		Embed: &networkv1directory.Listing_Embed{
	// 			Service: networkv1directory.Listing_Embed_DIRECTORY_LISTING_EMBED_SERVICE_ANGELTHUMP,
	// 			Id:      "psrngafk",
	// 		},
	// 	},
	// }

	// res := &networkv1directory.PublishResponse{}
	// err = directoryClient.Publish(ctx, &networkv1directory.PublishRequest{Listing: listing}, res)
	// if err != nil {
	// 	return nil, err
	// }

	// err = directoryClient.Join(
	// 	ctx,
	// 	&networkv1directory.JoinRequest{Id: res.Id},
	// 	&networkv1directory.JoinResponse{},
	// )
	// if err != nil {
	// 	return nil, err
	// }

	return &networkv1directory.FrontendTestResponse{}, nil
}
