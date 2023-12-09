package llm

import "fmt"

type Launcher interface {
	LaunchLLMCheck(currentUserId int32, testId, llmSlug string) (error, *LaunchResult)
}

type LaunchResult struct {
	Identifier string
}

func NewLauncher() Launcher {
	return defaultLauncher{}
}

type defaultLauncher struct {
}

func (launcher defaultLauncher) LaunchLLMCheck(currentUserId int32, testId, llmSlug string) (error, *LaunchResult) {
	return fmt.Errorf("not implemented yet"), nil
}
