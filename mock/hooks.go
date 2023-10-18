package mock

import "github.com/go-semantic-release/semantic-release/v2/pkg/hooks"

var _ hooks.Hooks = &Mock{}

func (m *Mock) Success(_ *hooks.SuccessHookConfig) error {
	m.Log.Println("success")
	return nil
}

func (m *Mock) NoRelease(_ *hooks.NoReleaseConfig) error {
	m.Log.Println("no release")
	return nil
}
