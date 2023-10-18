#!/usr/bin/env bash

set -euo pipefail

[[ ! -f ./semantic-release ]] && {
  echo "downloading semantic-release..."
  curl -SL "https://get-release.xyz/semantic-release/$(go env GOOS)/$(go env GOARCH)" -o ./semantic-release
  chmod +x ./semantic-release
}

echo "running build-local script..."
scriptsDir="$(dirname "${BASH_SOURCE[0]}")"
bash "${scriptsDir}/build-local.sh"

echo "running semantic-release..."
./semantic-release \
  --changelog-generator mock \
  --commit-analyzer mock \
  --provider mock \
  --ci-condition mock \
  --hooks mock \
  --files-updater mock \
  --update test.mock
