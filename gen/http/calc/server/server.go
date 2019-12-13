// Code generated by goa v3.0.8, DO NOT EDIT.
//
// calc HTTP server
//
// Command:
// $ goa gen calc/design

package server

import (
	calc "calc/gen/calc"
	"context"
	"net/http"
	"path"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the calc service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Add    http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the calc service endpoints.
func New(
	e *calc.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Add", "GET", "/add/{a}/{b}"},
			{"./gen/http/openapi.json", "GET", "/ui/swagger.json"},
			{"ui/", "GET", "/ui"},
		},
		Add: NewAddHandler(e.Add, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "calc" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Add = m(s.Add)
}

// Mount configures the mux to serve the calc endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAddHandler(mux, h.Add)
	MountGenHTTPOpenapiJSON(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/http/openapi.json")
	}))
	MountUI(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upath := path.Clean(r.URL.Path)
		rpath := upath
		if strings.HasPrefix(upath, "/ui") {
			rpath = upath[3:]
		}
		http.ServeFile(w, r, path.Join("ui/", rpath))
	}))
}

// MountAddHandler configures the mux to serve the "calc" service "add"
// endpoint.
func MountAddHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/add/{a}/{b}", f)
}

// NewAddHandler creates a HTTP handler which loads the HTTP request and calls
// the "calc" service "add" endpoint.
func NewAddHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeAddRequest(mux, dec)
		encodeResponse = EncodeAddResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add")
		ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/ui/swagger.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/ui/swagger.json", h.ServeHTTP)
}

// MountUI configures the mux to serve GET request made to "/ui".
func MountUI(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/ui/", h.ServeHTTP)
	mux.Handle("GET", "/ui/*filepath", h.ServeHTTP)
}
