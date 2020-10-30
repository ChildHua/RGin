package RGin

import (
	"context"
	"net/http"
)

type RContext struct {
	In  http.Request
	Out http.ResponseWriter
	Ctx context.Context
}
