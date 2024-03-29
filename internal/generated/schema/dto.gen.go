// Package dto provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package dto

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	AccessTokenScopes = "AccessToken.Scopes"
)

// Defines values for GetLLMCheckStatusValueStatus.
const (
	COMPLETED  GetLLMCheckStatusValueStatus = "COMPLETED"
	ERROR      GetLLMCheckStatusValueStatus = "ERROR"
	INPROGRESS GetLLMCheckStatusValueStatus = "IN_PROGRESS"
	NOTSTARTED GetLLMCheckStatusValueStatus = "NOT_STARTED"
	UNDEFINED  GetLLMCheckStatusValueStatus = "UNDEFINED"
)

// Defines values for LaunchLLMCheckRequestLlmSlug.
const (
	Dummy    LaunchLLMCheckRequestLlmSlug = "dummy"
	Gigachat LaunchLLMCheckRequestLlmSlug = "gigachat"
	Gpt4     LaunchLLMCheckRequestLlmSlug = "gpt4"
)

// ApiError defines model for ApiError.
type ApiError struct {
	// ErrorCode Error unique code
	ErrorCode int `json:"error_code"`

	// ErrorMessage Error description
	ErrorMessage string `json:"error_message"`
}

// CreateTestQuestionAnswerPayload defines model for CreateTestQuestionAnswerPayload.
type CreateTestQuestionAnswerPayload struct {
	// IsCorrect Is answer correct
	IsCorrect bool `json:"is_correct"`

	// Number Answer number in the list
	Number int `json:"number"`

	// Text Answer text
	Text string `json:"text"`
}

// CreateTestQuestionPayload defines model for CreateTestQuestionPayload.
type CreateTestQuestionPayload struct {
	// Answers Question answers
	Answers []CreateTestQuestionAnswerPayload `json:"answers"`

	// Number Question number in the list
	Number int `json:"number"`

	// Text Question text
	Text string `json:"text"`
}

// CreateTestRequest defines model for CreateTestRequest.
type CreateTestRequest struct {
	// Description Test description
	Description *string `json:"description,omitempty"`

	// Name Test unique name
	Name string `json:"name"`

	// Questions Test questions
	Questions []CreateTestQuestionPayload `json:"questions"`
}

// GetLLMCheckResultLLMAnswer defines model for GetLLMCheckResultLLMAnswer.
type GetLLMCheckResultLLMAnswer struct {
	// QuestionNumber Question number in test
	QuestionNumber int `json:"question_number"`

	// SelectedAnswerNumber LLM selected answer number in questions list
	SelectedAnswerNumber int `json:"selected_answer_number"`
}

// GetLLMCheckResultResponse defines model for GetLLMCheckResultResponse.
type GetLLMCheckResultResponse struct {
	// Results Results of all LLM analyses
	Results []GetLLMCheckResultValue `json:"results"`
}

// GetLLMCheckResultValue defines model for GetLLMCheckResultValue.
type GetLLMCheckResultValue struct {
	// Answers LLM answers of the test questions
	Answers []GetLLMCheckResultLLMAnswer `json:"answers"`

	// LlmSlug Large language model unique name
	LlmSlug string `json:"llm_slug"`
}

// GetLLMCheckStatusResponse defines model for GetLLMCheckStatusResponse.
type GetLLMCheckStatusResponse struct {
	// Statuses Statuses of all LLM analyses
	Statuses []GetLLMCheckStatusValue `json:"statuses"`
}

// GetLLMCheckStatusValue defines model for GetLLMCheckStatusValue.
type GetLLMCheckStatusValue struct {
	// LlmSlug Large language model unique name
	LlmSlug string `json:"llm_slug"`

	// Status LLM analysis status
	Status GetLLMCheckStatusValueStatus `json:"status"`
}

// GetLLMCheckStatusValueStatus LLM analysis status
type GetLLMCheckStatusValueStatus string

// LaunchLLMCheckRequest defines model for LaunchLLMCheckRequest.
type LaunchLLMCheckRequest struct {
	// LlmSlug Large language model unique name
	LlmSlug LaunchLLMCheckRequestLlmSlug `json:"llm_slug"`
}

// LaunchLLMCheckRequestLlmSlug Large language model unique name
type LaunchLLMCheckRequestLlmSlug string

// LaunchLLMCheckResponse defines model for LaunchLLMCheckResponse.
type LaunchLLMCheckResponse struct {
	// LaunchIdentifier Launch identifier UUID v4
	LaunchIdentifier openapi_types.UUID `json:"launch_identifier"`
}

// QuestionAnswer defines model for QuestionAnswer.
type QuestionAnswer struct {
	// IsCorrect Is answer correct
	IsCorrect bool `json:"is_correct"`

	// Number Answer number in the list
	Number int `json:"number"`

	// Text Answer text
	Text string `json:"text"`
}

// SignInRequest defines model for SignInRequest.
type SignInRequest struct {
	// UserLogin User unique login
	UserLogin string `json:"user_login"`

	// UserPasswordHash User password hash in SHA512
	UserPasswordHash string `json:"user_password_hash"`
}

// SignUpRequest defines model for SignUpRequest.
type SignUpRequest struct {
	// UserLogin User unique login
	UserLogin string `json:"user_login"`

	// UserPasswordHash User password hash in SHA512
	UserPasswordHash string `json:"user_password_hash"`
}

// Test defines model for Test.
type Test struct {
	// Description Test description
	Description *string `json:"description,omitempty"`

	// Identifier Test identifier UUID v4
	Identifier openapi_types.UUID `json:"identifier"`

	// Name Test unique name
	Name string `json:"name"`

	// Questions Test questions
	Questions []TestQuestion `json:"questions"`
}

// TestQuestion defines model for TestQuestion.
type TestQuestion struct {
	// Answers Question answers
	Answers []QuestionAnswer `json:"answers"`

	// Number Question number in the list
	Number int `json:"number"`

	// Text Question text
	Text string `json:"text"`
}

// RefreshToken defines model for RefreshToken.
type RefreshToken = string

// TestId defines model for TestId.
type TestId = openapi_types.UUID

// AuthRefreshTokenParams defines parameters for AuthRefreshToken.
type AuthRefreshTokenParams struct {
	XLLMCheckerRefreshToken RefreshToken `json:"X-LLM-Checker-Refresh-Token"`
}

// TestsMyParams defines parameters for TestsMy.
type TestsMyParams struct {
	// PageNumber Pagination page number
	PageNumber *int `form:"page-number,omitempty" json:"page-number,omitempty"`

	// PageSize Pagination page size
	PageSize *int `form:"page-size,omitempty" json:"page-size,omitempty"`
}

// AuthSignInJSONRequestBody defines body for AuthSignIn for application/json ContentType.
type AuthSignInJSONRequestBody = SignInRequest

// AuthSignUpJSONRequestBody defines body for AuthSignUp for application/json ContentType.
type AuthSignUpJSONRequestBody = SignUpRequest

// TestCreateJSONRequestBody defines body for TestCreate for application/json ContentType.
type TestCreateJSONRequestBody = CreateTestRequest

// LlmLaunchJSONRequestBody defines body for LlmLaunch for application/json ContentType.
type LlmLaunchJSONRequestBody = LaunchLLMCheckRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Updating system access token using refresh token
	// (POST /api/v1/auth/refresh-token)
	AuthRefreshToken(w http.ResponseWriter, r *http.Request, params AuthRefreshTokenParams)
	// User login to the system with login and password
	// (POST /api/v1/auth/sign-in)
	AuthSignIn(w http.ResponseWriter, r *http.Request)
	// User registration in the system
	// (POST /api/v1/auth/sign-up)
	AuthSignUp(w http.ResponseWriter, r *http.Request)
	// Creating a new test
	// (PUT /api/v1/test/create)
	TestCreate(w http.ResponseWriter, r *http.Request)
	// Deleting an existing test
	// (DELETE /api/v1/test/{testId}/delete)
	TestDelete(w http.ResponseWriter, r *http.Request, testId TestId)
	// Getting complete information of a specific test
	// (GET /api/v1/test/{testId}/get)
	TestById(w http.ResponseWriter, r *http.Request, testId TestId)
	// Launching test analysis using a generative language model
	// (POST /api/v1/test/{testId}/llm/launch)
	LlmLaunch(w http.ResponseWriter, r *http.Request, testId TestId)
	// Getting all the test analysis results
	// (GET /api/v1/test/{testId}/llm/result)
	LlmResult(w http.ResponseWriter, r *http.Request, testId TestId)
	// Getting the current status of test analysis
	// (GET /api/v1/test/{testId}/llm/status)
	LlmStatus(w http.ResponseWriter, r *http.Request, testId TestId)
	// Getting "my" created tests
	// (GET /api/v1/tests/my)
	TestsMy(w http.ResponseWriter, r *http.Request, params TestsMyParams)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Updating system access token using refresh token
// (POST /api/v1/auth/refresh-token)
func (_ Unimplemented) AuthRefreshToken(w http.ResponseWriter, r *http.Request, params AuthRefreshTokenParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// User login to the system with login and password
// (POST /api/v1/auth/sign-in)
func (_ Unimplemented) AuthSignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// User registration in the system
// (POST /api/v1/auth/sign-up)
func (_ Unimplemented) AuthSignUp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Creating a new test
// (PUT /api/v1/test/create)
func (_ Unimplemented) TestCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Deleting an existing test
// (DELETE /api/v1/test/{testId}/delete)
func (_ Unimplemented) TestDelete(w http.ResponseWriter, r *http.Request, testId TestId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Getting complete information of a specific test
// (GET /api/v1/test/{testId}/get)
func (_ Unimplemented) TestById(w http.ResponseWriter, r *http.Request, testId TestId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Launching test analysis using a generative language model
// (POST /api/v1/test/{testId}/llm/launch)
func (_ Unimplemented) LlmLaunch(w http.ResponseWriter, r *http.Request, testId TestId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Getting all the test analysis results
// (GET /api/v1/test/{testId}/llm/result)
func (_ Unimplemented) LlmResult(w http.ResponseWriter, r *http.Request, testId TestId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Getting the current status of test analysis
// (GET /api/v1/test/{testId}/llm/status)
func (_ Unimplemented) LlmStatus(w http.ResponseWriter, r *http.Request, testId TestId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Getting "my" created tests
// (GET /api/v1/tests/my)
func (_ Unimplemented) TestsMy(w http.ResponseWriter, r *http.Request, params TestsMyParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// AuthRefreshToken operation middleware
func (siw *ServerInterfaceWrapper) AuthRefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params AuthRefreshTokenParams

	headers := r.Header

	// ------------- Required header parameter "X-LLM-Checker-Refresh-Token" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-LLM-Checker-Refresh-Token")]; found {
		var XLLMCheckerRefreshToken RefreshToken
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "X-LLM-Checker-Refresh-Token", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-LLM-Checker-Refresh-Token", runtime.ParamLocationHeader, valueList[0], &XLLMCheckerRefreshToken)
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "X-LLM-Checker-Refresh-Token", Err: err})
			return
		}

		params.XLLMCheckerRefreshToken = XLLMCheckerRefreshToken

	} else {
		err := fmt.Errorf("Header parameter X-LLM-Checker-Refresh-Token is required, but not found")
		siw.ErrorHandlerFunc(w, r, &RequiredHeaderError{ParamName: "X-LLM-Checker-Refresh-Token", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AuthRefreshToken(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AuthSignIn operation middleware
func (siw *ServerInterfaceWrapper) AuthSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AuthSignIn(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AuthSignUp operation middleware
func (siw *ServerInterfaceWrapper) AuthSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AuthSignUp(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// TestCreate operation middleware
func (siw *ServerInterfaceWrapper) TestCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, AccessTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.TestCreate(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// TestDelete operation middleware
func (siw *ServerInterfaceWrapper) TestDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "testId" -------------
	var testId TestId

	err = runtime.BindStyledParameterWithLocation("simple", false, "testId", runtime.ParamLocationPath, chi.URLParam(r, "testId"), &testId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "testId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, AccessTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.TestDelete(w, r, testId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// TestById operation middleware
func (siw *ServerInterfaceWrapper) TestById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "testId" -------------
	var testId TestId

	err = runtime.BindStyledParameterWithLocation("simple", false, "testId", runtime.ParamLocationPath, chi.URLParam(r, "testId"), &testId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "testId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, AccessTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.TestById(w, r, testId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LlmLaunch operation middleware
func (siw *ServerInterfaceWrapper) LlmLaunch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "testId" -------------
	var testId TestId

	err = runtime.BindStyledParameterWithLocation("simple", false, "testId", runtime.ParamLocationPath, chi.URLParam(r, "testId"), &testId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "testId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, AccessTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LlmLaunch(w, r, testId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LlmResult operation middleware
func (siw *ServerInterfaceWrapper) LlmResult(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "testId" -------------
	var testId TestId

	err = runtime.BindStyledParameterWithLocation("simple", false, "testId", runtime.ParamLocationPath, chi.URLParam(r, "testId"), &testId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "testId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, AccessTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LlmResult(w, r, testId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LlmStatus operation middleware
func (siw *ServerInterfaceWrapper) LlmStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "testId" -------------
	var testId TestId

	err = runtime.BindStyledParameterWithLocation("simple", false, "testId", runtime.ParamLocationPath, chi.URLParam(r, "testId"), &testId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "testId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, AccessTokenScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LlmStatus(w, r, testId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// TestsMy operation middleware
func (siw *ServerInterfaceWrapper) TestsMy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, AccessTokenScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params TestsMyParams

	// ------------- Optional query parameter "page-number" -------------

	err = runtime.BindQueryParameter("form", true, false, "page-number", r.URL.Query(), &params.PageNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page-number", Err: err})
		return
	}

	// ------------- Optional query parameter "page-size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page-size", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page-size", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.TestsMy(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
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
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
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

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/auth/refresh-token", wrapper.AuthRefreshToken)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/auth/sign-in", wrapper.AuthSignIn)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/auth/sign-up", wrapper.AuthSignUp)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/v1/test/create", wrapper.TestCreate)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/test/{testId}/delete", wrapper.TestDelete)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/test/{testId}/get", wrapper.TestById)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/test/{testId}/llm/launch", wrapper.LlmLaunch)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/test/{testId}/llm/result", wrapper.LlmResult)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/test/{testId}/llm/status", wrapper.LlmStatus)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/tests/my", wrapper.TestsMy)
	})

	return r
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
