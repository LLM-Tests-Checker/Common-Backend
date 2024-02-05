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
	Dummy      LaunchLLMCheckRequestLlmSlug = "Dummy"
	GPT3       LaunchLLMCheckRequestLlmSlug = "GPT3"
	YandexGPT2 LaunchLLMCheckRequestLlmSlug = "YandexGPT2"
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

	// UserName User friendly name
	UserName string `json:"user_name"`

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
	// User login to the system with username and password
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

// User login to the system with username and password
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
	"d1zb5rU3G9vtnwEAAP//mmZ77d5OAAA=",
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
