// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

// Comment defines model for Comment.
type Comment struct {
	// ArticleId ID of the article the comment belongs to
	ArticleId int `json:"article_id"`

	// CreatedAt Timestamp when the comment was created
	CreatedAt time.Time `json:"created_at"`

	// Id Comment ID
	Id int `json:"id"`

	// Message The comment text
	Message string `json:"message"`
}

// CommentRequest defines model for CommentRequest.
type CommentRequest struct {
	// ArticleId ID of the article to comment on
	ArticleId int `json:"article_id"`

	// Message The comment text
	Message string `json:"message"`
}

// Error defines model for Error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Todo defines model for Todo.
type Todo struct {
	Contents string `json:"contents"`
	Id       int    `json:"id"`
	NiceNum  int    `json:"nice_num"`
	Title    string `json:"title"`
	Username string `json:"username"`
}

// TodoNiceRequest defines model for TodoNiceRequest.
type TodoNiceRequest struct {
	Id int `json:"id"`
}

// TodoRequest defines model for TodoRequest.
type TodoRequest struct {
	Contents string `json:"contents"`
	Title    string `json:"title"`
	Username string `json:"username"`
}

// GetTodoListParams defines parameters for GetTodoList.
type GetTodoListParams struct {
	// Page The page number to retrieve
	Page *int `form:"page,omitempty" json:"page,omitempty"`

	// PageSize The number of items per page
	PageSize *int `form:"pageSize,omitempty" json:"pageSize,omitempty"`
}

// CreateCommentJSONRequestBody defines body for CreateComment for application/json ContentType.
type CreateCommentJSONRequestBody = CommentRequest

// CreateTodoJSONRequestBody defines body for CreateTodo for application/json ContentType.
type CreateTodoJSONRequestBody = TodoRequest

// IncrementNiceJSONRequestBody defines body for IncrementNice for application/json ContentType.
type IncrementNiceJSONRequestBody = TodoNiceRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new comment
	// (POST /comment)
	CreateComment(w http.ResponseWriter, r *http.Request)
	// Returns hello message
	// (GET /hello)
	GetHello(w http.ResponseWriter, r *http.Request)
	// Create a new todo
	// (POST /todo)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	// Get a list of todos
	// (GET /todo/list)
	GetTodoList(w http.ResponseWriter, r *http.Request, params GetTodoListParams)
	// Increment nice count for a todo
	// (POST /todo/nice)
	IncrementNice(w http.ResponseWriter, r *http.Request)
	// Delete a todo
	// (DELETE /todo/{id})
	DeleteTodo(w http.ResponseWriter, r *http.Request, id int)
	// Get a todo by ID
	// (GET /todo/{id})
	GetTodoById(w http.ResponseWriter, r *http.Request, id int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// CreateComment operation middleware
func (siw *ServerInterfaceWrapper) CreateComment(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateComment(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetHello operation middleware
func (siw *ServerInterfaceWrapper) GetHello(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetHello(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateTodo operation middleware
func (siw *ServerInterfaceWrapper) CreateTodo(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTodo(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetTodoList operation middleware
func (siw *ServerInterfaceWrapper) GetTodoList(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTodoListParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "pageSize", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTodoList(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// IncrementNice operation middleware
func (siw *ServerInterfaceWrapper) IncrementNice(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.IncrementNice(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteTodo operation middleware
func (siw *ServerInterfaceWrapper) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", mux.Vars(r)["id"], &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteTodo(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetTodoById operation middleware
func (siw *ServerInterfaceWrapper) GetTodoById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", mux.Vars(r)["id"], &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTodoById(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
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

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
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

	r.HandleFunc(options.BaseURL+"/comment", wrapper.CreateComment).Methods("POST")

	r.HandleFunc(options.BaseURL+"/hello", wrapper.GetHello).Methods("GET")

	r.HandleFunc(options.BaseURL+"/todo", wrapper.CreateTodo).Methods("POST")

	r.HandleFunc(options.BaseURL+"/todo/list", wrapper.GetTodoList).Methods("GET")

	r.HandleFunc(options.BaseURL+"/todo/nice", wrapper.IncrementNice).Methods("POST")

	r.HandleFunc(options.BaseURL+"/todo/{id}", wrapper.DeleteTodo).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/todo/{id}", wrapper.GetTodoById).Methods("GET")

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RX32/bNhD+Vwhuj2rlbHko9NYsQ2dsSIu2b0MQMNRZYiH+CHlK5gX+34cjJVu2JCdA",
	"42zA3mSZvPvuu7vvTo9cWu2sAYOBF488yBq0iI+/WK3BID06bx14VBD/EB6VbOBGlfSrhCC9cqis4QVf",
	"XjK7YlgD607FZ5lMsVtorKkCQ8szjmsHvODKIFTg+Sbj0oNAKG8Ejg1/VRoCCu3YQw1mz+qDCKy7yjO+",
	"sl6TAV4KhDeoNOx8BfTKVORqCnsXMFteTqLTEIKoYALaAAvCXzj2t8m4h7tWeSh58Sc5z4Ys7mzvcXC9",
	"tWNvv4FEAtFh/Ax3LYTvz43dArfmxEFPxjsV4q/eWz+OTNoywjiK8TiEActkbMr5V1vaKd8G+w6ZKaUx",
	"LKMk3JhWT/+LChuYtNcG8EboZ8QTqUyGsh3GgYUBiLlgr5SE2WqaDm2MYs74rOGjhL4ENcdYGaOl28qs",
	"bA9NyIg5ueL1m1o9CBShVgRivw/ef1qylfVMCyMqZSqGtrRMIeiwTU4RyWDvPy15xu/Bh3T17O3i7YIs",
	"WgdGOMUL/nN8lXEnsI7k5HKgwjZRSUQKcr8sSbWiYPRinViAgBe2XA+YjtrgXKNkvJl/CwSh13t6+tHD",
	"ihf8h3w3EPJuGuQHqrPZZxt9C/FFcNaElOCfFouX9p7c7rP/8Xfi7/wFnSX1mXB1IUrmtwRkPLRaC7/e",
	"ZoAJZuCh10TKvagClWIP/5pu5TU0TayzCiaS+QHwt3jgSTpJc3PXCHUQ22FbjOL40koJIRzE8Bmw9Saw",
	"CI/thLIPIupiigB7iTxSjvH4aWpxqCuvXIgxqpkqnK8ITFzMMJk3KtE4Vw90/g86Q6rghQYET4amBrIT",
	"FTDT6lvwNNs9oFdwT3mkMuF3Lfg1DYWkay6leBd7CSvRNsiLs4xrZZSm4XWWTQyAKeedX7tK6scceNa5",
	"mPP+Rf09h2DxBITr78xzUuhnJXzrXXgv1s8qgA+ATDBKbVy5bGnDkRKgIT3fUUsjPZCC0Kg+YVMNN4H/",
	"UGORvJ+fXt6vLLKVbU15kMkt+4yyxKRtDcaBL55q7EdVblJNN4AwzutlfN8p5UFrx4ahJWDXL3HZ209J",
	"Ntb9Yw1yPt7ir+gDINH6bxOd6JilNTsqkBfrZflKLP4vqj7pV1xlb9fdB/FBmdN58PfTw+gS7qGxLvZN",
	"OkULuG9om0Z0RZ43VoqmtgGLd4t3C05Edx5Gs4VQfOyzvktlJ85zX/ETN7aL5PXmnwAAAP//LQjEnPgQ",
	"AAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
