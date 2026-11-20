package users

import (
	"context"
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gonvy.internal/internal/httperr"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	if svc == nil {
		panic("service is nil")
	}
	return &Handler{svc: svc}
}

func (h *Handler) Mount(api huma.API) {
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

func (h *Handler) register(ctx context.Context, in *registerInput) (*registerOutput, error) {
	out, err := h.svc.RegisterUserAccount(ctx, in)
	if err != nil {
		if errors.Is(err, ErrDuplicateEmail) {
			return nil, httperr.Conflict("DUPLICATE_EMAIL", "The provided email is in use")
		}
		return nil, err
	}
	return &registerOutput{Body: *out}, nil
}
