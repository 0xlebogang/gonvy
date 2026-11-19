package httpapi

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type healthOutput struct {
	Body struct {
		Status string `json:"status"`
	}
}

func registerHealthz(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "server:healthz",
		Method:      http.MethodGet,
		Path:        "/healthz",
		Summary:     "Liveness check",
		Tags:        []string{"Probe"},
	}, func(ctx context.Context, _ *struct{}) (*healthOutput, error) {
		output := &healthOutput{}
		output.Body.Status = "ok"
		return output, nil
	})
}
