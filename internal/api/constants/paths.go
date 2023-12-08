package constants

const (
	SignInPath             = "/api/v1/auth/sign-in"
	SignUpPath             = "/api/v1/auth/sign-up"
	RefreshAccessTokenPath = "/api/v1/auth/refresh-token"
)

const (
	GetMyTestsPath     = "/api/v1/tests/my"
	GetTestByIdPath    = "/api/v1/test/{" + TestIdPathParameter + "}/get"
	CreateTestPath     = "/api/v1/test/create"
	DeleteTestByIdPath = "/api/v1/test/{" + TestIdPathParameter + "}/delete"
)

const (
	LaunchLLMCheckPath    = "/api/v1/test/{" + TestIdPathParameter + "}/llm/launch"
	GetLLMCheckStatusPath = "/api/v1/test/{" + TestIdPathParameter + "}/llm/status"
	GetLLMCheckResultPath = "/api/v1/test/{" + TestIdPathParameter + "}/llm/result"
)

const (
	TestIdPathParameter = "testId"
)
