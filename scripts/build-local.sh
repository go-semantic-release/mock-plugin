#!/usr/bin/env bash

set -euo pipefail

echo "building..."
rootPluginDir=".semrel/$(go env GOOS)_$(go env GOARCH)"
pluginBinFile="${rootPluginDir}/mock-plugin"
go build -o "$pluginBinFile" ./

for pluginType in changelog-generator commit-analyzer condition files-updater hooks provider; do
  pluginDir="${rootPluginDir}/${pluginType}-mock/0.0.0-dev/"
  [[ ! -d "$pluginDir" ]] && {
    echo "creating $pluginDir"
    mkdir -p "$pluginDir"
  }
  echo "installing plugin to $pluginDir"
  cp "$pluginBinFile" "$pluginDir"
done
