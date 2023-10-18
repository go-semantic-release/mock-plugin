package mock

import "github.com/go-semantic-release/semantic-release/v2/pkg/condition"

var _ condition.CICondition = &Mock{}

func (m *Mock) RunCondition(_ map[string]string) error {
	m.Log.Println("running ci condition")
	return nil
}

func (m *Mock) GetCurrentBranch() string {
	m.Log.Println("getting current branch")
	return "main"
}

func (m *Mock) GetCurrentSHA() string {
	m.Log.Println("getting current sha")
	return "HEAD"
}
