package users

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gonvy.internal/internal/httperr"
)

type Handler interface {
	Mount(api huma.API)
}

type handler struct {
	svc Service
}

func Newhandler(svc Service) Handler {
	if svc == nil {
		panic("service is nil")
	}
	return &handler{svc: svc}
}

func (h *handler) Mount(api huma.API) {
	grp := huma.NewGroup(api, "/users")
	huma.Register(grp, huma.Operation{
		OperationID: "users:register",
		Method:      http.MethodPost,
		Path:        "",
		Summary:     "Create a user account",
		Tags:        []string{"Users"},
	}, h.register)
}

func (h *handler) register(ctx context.Context, in *RegisterInput) (*RegisterOutput, error) {
	out, err := h.svc.RegisterUserAccount(ctx, in)
	if err != nil {
		if errors.Is(err, ErrDuplicateEmail) {
			return nil, httperr.Conflict("DUPLICATE_EMAIL", "The provided email is in use")
		}
		return nil, err
	}
	resp := &RegisterOutput{}
	resp.Body.UserID = out.User.ID
	resp.Body.AccountID = out.Account.ID
	resp.Body.Name = out.User.Name
	resp.Body.Email = out.User.Email
	return resp, nil
}
