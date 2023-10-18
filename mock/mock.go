package mock

import (
	"log"
)

type Mock struct {
	PluginVersion string
	Log           *log.Logger
}

func (m *Mock) Init(_ map[string]string) error {
	m.Log.Println("init")
	return nil
}

func (m *Mock) Name() string {
	return "mock"
}

func (m *Mock) Version() string {
	return m.PluginVersion
}
