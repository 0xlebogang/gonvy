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

type registerInput struct {
	Body struct {
		Name     string `json:"name" minLength:"1" maxLength:"100"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}

type registerOutput struct {
	Body RegisteredUser
}

func (h *handler) register(ctx context.Context, in *registerInput) (*registerOutput, error) {
	out, err := h.svc.RegisterUserAccount(ctx, in)
	if err != nil {
		if errors.Is(err, ErrDuplicateEmail) {
			return nil, httperr.Conflict("DUPLICATE_EMAIL", "The provided email is in use")
		}
		return nil, err
	}
	resp := &registerOutput{}
	resp.Body = RegisteredUser{
		UserID:    out.UserID,
		AccountID: out.AccountID,
		Name:      out.Name,
		Email:     out.Email,
	}
	return resp, nil
}
