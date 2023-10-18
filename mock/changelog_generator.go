package mock

import "github.com/go-semantic-release/semantic-release/v2/pkg/generator"

var _ generator.ChangelogGenerator = &Mock{}

func (m *Mock) Generate(gc *generator.ChangelogGeneratorConfig) string {
	m.Log.Printf("generating changelog for %s", gc.NewVersion)
	return "mocked changelog output"
}
