// Package experimental provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package experimental

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns OK if experimental API is enabled.
	// (GET /v2/experimental)
	ExperimentalCheck(ctx echo.Context) error
	// Fast track for broadcasting a raw transaction or transaction group to the network through the tx handler without performing most of the checks and reporting detailed errors. Should be only used for development and performance testing.
	// (POST /v2/transactions/async)
	RawTransactionAsync(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ExperimentalCheck converts echo context to params.
func (w *ServerInterfaceWrapper) ExperimentalCheck(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ExperimentalCheck(ctx)
	return err
}

// RawTransactionAsync converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransactionAsync(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransactionAsync(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/v2/experimental", wrapper.ExperimentalCheck, m...)
	router.POST(baseURL+"/v2/transactions/async", wrapper.RawTransactionAsync, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
