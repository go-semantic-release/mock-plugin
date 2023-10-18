package mock

import "github.com/go-semantic-release/semantic-release/v2/pkg/updater"

var _ updater.FilesUpdater = &Mock{}

func (m *Mock) ForFiles() string {
	m.Log.Println("listing files to update")
	return ".*\\.mock"
}

func (m *Mock) Apply(file, newVersion string) error {
	m.Log.Printf("updating version in file %s to %s", file, newVersion)
	return nil
}
