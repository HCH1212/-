package rpc

import (
	"context"
	userhoster "rpc-server/rpc/kitex_gen/userhoster"
	"rpc-server/service"
)

// UserhosterImpl implements the last service interface defined in the IDL.
type UserhosterImpl struct{}

// Register implements the UserhosterImpl interface.
func (s *UserhosterImpl) Register(ctx context.Context, u *userhoster.User, h *userhoster.Hoster) (err error) {
	// TODO: Your code here...
	return service.Register(u, h)
}

// Login implements the UserhosterImpl interface.
func (s *UserhosterImpl) Login(ctx context.Context, u *userhoster.User, h *userhoster.Hoster) (err error) {
	// TODO: Your code here...
	return service.Login(u, h)
}

// SetToken implements the UserhosterImpl interface.
func (s *UserhosterImpl) SetToken(ctx context.Context, u *userhoster.User, h *userhoster.Hoster) (resp string, err error) {
	// TODO: Your code here...
	resp, err = service.SetToken(u, h)
	if err != nil {
		return
	}
	err = service.WriteToken(u, h, resp)
	return
}
