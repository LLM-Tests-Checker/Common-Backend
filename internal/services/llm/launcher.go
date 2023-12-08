package llm

type LauncherService interface {
	LaunchLLMCheck(currentUserId int32, testId int32, llmSlug string) (error, LaunchResult)
}

type LaunchResult struct {
	Identifier string
}
