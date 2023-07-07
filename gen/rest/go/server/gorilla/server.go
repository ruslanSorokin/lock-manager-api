// Package grlapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package grlapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
)

// ResourceId defines model for resource_id.
type ResourceId = UuidV4

// Token defines model for token.
type Token = UuidV4

// UuidV4 defines model for uuid_v4.
type UuidV4 = openapi_types.UUID

// PostLocksParams defines parameters for PostLocks.
type PostLocksParams struct {
	ResourceId ResourceId `form:"resource_id" json:"resource_id"`
}

// DeleteLocksResourceIdParams defines parameters for DeleteLocksResourceId.
type DeleteLocksResourceIdParams struct {
	Token Token `form:"token" json:"token"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Lock a resource.
	// (POST /locks)
	PostLocks(w http.ResponseWriter, r *http.Request, params PostLocksParams)
	// Unlock a resource.
	// (DELETE /locks/{resource_id})
	DeleteLocksResourceId(w http.ResponseWriter, r *http.Request, resourceId ResourceId, params DeleteLocksResourceIdParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// PostLocks operation middleware
func (siw *ServerInterfaceWrapper) PostLocks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostLocksParams

	// ------------- Required query parameter "resource_id" -------------

	if paramValue := r.URL.Query().Get("resource_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "resource_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "resource_id", r.URL.Query(), &params.ResourceId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "resource_id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostLocks(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteLocksResourceId operation middleware
func (siw *ServerInterfaceWrapper) DeleteLocksResourceId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "resource_id" -------------
	var resourceId ResourceId

	err = runtime.BindStyledParameter("simple", false, "resource_id", mux.Vars(r)["resource_id"], &resourceId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "resource_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteLocksResourceIdParams

	// ------------- Required query parameter "token" -------------

	if paramValue := r.URL.Query().Get("token"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "token"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "token", r.URL.Query(), &params.Token)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "token", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteLocksResourceId(w, r, resourceId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/locks", wrapper.PostLocks).Methods("POST")

	r.HandleFunc(options.BaseURL+"/locks/{resource_id}", wrapper.DeleteLocksResourceId).Methods("DELETE")

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RW72/bNhD9V4jbgH5RLNnJgEXfjMTbjLl24STbii5wGfFks5FIlTw5MQz97wNJ/4zd",
	"bVjWIvlkSyLvju+9e8clZLqstEJFFtIl2GyGJfd/DVpdmwwnUrhHXhSjHNIPS/jeYA4pfBdvt8arfXFd",
	"SzGZn0FzG4FAmxlZkdQKUhivwjEpUJHMJZrWnwqaCEjfo/o/UlxhZpCYj8doxolJa2sUjOeEhmUGOUk1",
	"ZZwVOrtnWjHO1qdssQuu2B2y2qJgpJnAAgk3K2mG27Wh7nUh6RLwkZdVgZDCmfih3TkV2Yk4//H8pN1G",
	"ccJz3j5JOmcdnrU7SZJ0IIJcm5ITpD4IRFDyxwGqKc0gPe1EUEq1+0iLysW2ZKSaQtM0EUiVa5eZJPm8",
	"A1fnW674FA3rvutDBHM0NgCTtNqtxJWsK1S8kpDCaStpJRBBxWnm6Y7dSf2/Sltyv/vgXjj00DLOFD4E",
	"WB4kzdhUzlGxjztq+ci4Eswg1Ua59Ss+NpA6KN1+B6MvyXCXoy8ghXfa0sAX4kozvERCY70mpKvic41m",
	"AREoXrpT72o0AoOfa2lQQEqmxmgl5n+vrN1ozW3jBGbQVlpZ9MB0krb7ybQiVB4iXlWFzHz58Ser1baF",
	"PJDGHY5k2D0ppAoAS8LSHi6olUPl8P3Ml+z4yHldONHMiKo0dozxYqYtpefJeRLzSsbzdhzCxMud0zR7",
	"ijMFHEpq80bffcKMjr3ZvODG8MV/6Nyw3CF7mK152syDI43HHrhlts4ytDavi2IRehpFi73X9ZuiYFOk",
	"reKcA2RHu3ovqFPrL93r3qh7xQJKzFPF5Ca71wC702Kx6v2zJHmGFNAYbSaFPNZnXQcu0zlDVZfrevwG",
	"Gxpmo579fb2ny8NqFwXSD9Af/tYd9C8n497V6GZ80Zv0L+F2RxV+3YEsnlL+LOKGo+uj5EXsDjNeW3Sn",
	"drukmvNCCrY1gPUX1+FoaZ/uPcACUlEgv6wtOfZz+RjoL/k97gZiPtGG0/PXxekOl5PuYNzrXr6f9P7o",
	"X11fvTxqd+doGBnb/isMcrFg+Cgt2a/BbROBrcuSm8X6EHxvmEdAfOqmjP8I3qDCPHxio54h5yCHXI2x",
	"1HO0/zAZtSoWTAYxB5fKJdljk/DS5/GzcH116osvTEU3xL/VUIyOT+Lg7c9Pt50Rh9M3OUTdk3kwFQJH",
	"4rV79fXo194QIvh9PBr+vHl6iU6+AvzlOPnZ63Xy4eh68tPoZvhNWRUarXpDwYO/ugXf+CvqF004fHY2",
	"7Hahma/Nbv8gb7lULHyGyF9r//5m7FzlrwAAAP//AD8pIvAOAAA=",
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
