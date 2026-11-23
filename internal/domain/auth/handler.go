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

func newhandler(svc Service) Handler {
	if svc == nil {
		panic("service is nil")
	}
	return &handler{svc: svc}
}

func (h *handler) Mount(api huma.API) {
	grp := huma.NewGroup(api, "/auth")
	huma.Register(grp, huma.Operation{
		OperationID:   "auth:register",
		Method:        http.MethodPost,
		Path:          "/register",
		Summary:       "Register a user account",
		Tags:          []string{"Auth"},
		DefaultStatus: http.StatusCreated,
	}, h.register)
	huma.Register(grp, huma.Operation{
		OperationID:   "auth:login",
		Method:        http.MethodPost,
		Path:          "/login",
		Summary:       "Authenticate an existing user and create a session",
		Tags:          []string{"Auth"},
		DefaultStatus: http.StatusOK,
	}, h.login)
}

func (h *handler) register(ctx context.Context, in *users.RegisterInput) (*registerOutput, error) {
	out, err := h.svc.RegisterUserAccount(ctx, in)
	if err != nil {
		if errors.Is(err, users.ErrDuplicateEmail) {
			return nil, httperr.Conflict("DUPLICATE_EMAIL", "The provided email address is already in use")
		}
		return nil, err
	}
	resp := &registerOutput{}
	resp.Body.UserID = out.User.UserID
	resp.Body.AccountID = out.User.AccountID
	resp.Body.Name = out.User.Name
	resp.Body.Email = out.User.Email
	resp.Body.Token = out.Session.Token
	resp.Body.ExpiresAt = out.Session.ExpiresAt
	return resp, nil
}

func (h *handler) login(ctx context.Context, in *loginInput) (*loginOutput, error) {
	out, err := h.svc.Login(ctx, in)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return nil, httperr.Unauthorized("")
		}
		return nil, err
	}
	resp := &loginOutput{}
	resp.Body.Token = out.Session.Token
	resp.Body.User = out.User
	resp.Body.Session = out.Session
	return resp, nil
}
