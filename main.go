package main

import (
	"log"
	"os"

	"github.com/go-semantic-release/mock-plugin/mock"
	"github.com/go-semantic-release/semantic-release/v2/pkg/analyzer"
	"github.com/go-semantic-release/semantic-release/v2/pkg/condition"
	"github.com/go-semantic-release/semantic-release/v2/pkg/generator"
	"github.com/go-semantic-release/semantic-release/v2/pkg/hooks"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"
	"github.com/go-semantic-release/semantic-release/v2/pkg/provider"
	"github.com/go-semantic-release/semantic-release/v2/pkg/updater"
)

var version = "dev"

func main() {
	mockPlugin := &mock.Mock{
		PluginVersion: version,
		Log:           log.New(os.Stderr, "", 0),
	}
	plugin.Serve(&plugin.ServeOpts{
		CommitAnalyzer: func() analyzer.CommitAnalyzer {
			return mockPlugin
		},
		CICondition: func() condition.CICondition {
			return mockPlugin
		},
		ChangelogGenerator: func() generator.ChangelogGenerator {
			return mockPlugin
		},
		Provider: func() provider.Provider {
			return mockPlugin
		},
		FilesUpdater: func() updater.FilesUpdater {
			return mockPlugin
		},
		Hooks: func() hooks.Hooks {
			return mockPlugin
		},
	})
}
