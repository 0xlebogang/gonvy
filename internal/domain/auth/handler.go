package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gonvy.internal/internal/domain/users"
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
	grp := huma.NewGroup(api, "/auth")
	huma.Register(grp, huma.Operation{
		OperationID: "auth:register",
		Method:      http.MethodPost,
		Path:        "/register",
		Summary:     "Register a user account",
		Tags:        []string{"Auth"},
	}, h.register)
}

func (h *handler) register(ctx context.Context, in *users.RegisterInput) (*users.RegisterOutput, error) {
	out, err := h.svc.RegisterUserAccount(ctx, in)
	if err != nil {
		if errors.Is(err, users.ErrDuplicateEmail) {
			return nil, httperr.Conflict("DUPLICATE_EMAIL", "The provided email address is already in use")
		}
		return nil, err
	}
	resp := &users.RegisterOutput{Body: *out}
	return resp, nil
}
