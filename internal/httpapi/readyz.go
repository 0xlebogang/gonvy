package httpapi

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gonvy.internal/internal/httperr"
	"gonvy.internal/internal/persistance/postgres"
)

type readyzOutput struct {
	Body struct {
		Status string `json:"status"`
	}
}

func registerReadyz(api huma.API, db postgres.DB) {
	huma.Register(api, huma.Operation{
		OperationID: "server:readyz",
		Method:      http.MethodGet,
		Path:        "/readyz",
		Summary:     "Readiness check",
		Tags:        []string{"Probe"},
	}, func(ctx context.Context, _ *struct{}) (*readyzOutput, error) {
		if err := db.Ping(); err != nil {
			return nil, httperr.ServiceUnavailable(err)
		}
		output := &readyzOutput{}
		output.Body.Status = "ready"
		return output, nil
	})
}
