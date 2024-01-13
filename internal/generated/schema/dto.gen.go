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
	PageNumber int `form:"page-number" json:"page-number"`

	// PageSize Pagination page size
	PageSize int `form:"page-size" json:"page-size"`
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

	// ------------- Required query parameter "page-number" -------------

	if paramValue := r.URL.Query().Get("page-number"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "page-number"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page-number", r.URL.Query(), &params.PageNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page-number", Err: err})
		return
	}

	// ------------- Required query parameter "page-size" -------------

	if paramValue := r.URL.Query().Get("page-size"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "page-size"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page-size", r.URL.Query(), &params.PageSize)
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
	"aPdOAAA=",
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
