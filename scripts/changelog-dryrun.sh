#!/usr/bin/env bash
# Local dry-run script to simulate changelog generation used by CI
# Usage: ./scripts/changelog-dryrun.sh [since-tag]
# If no tag is provided, it will use the most recent tag as the range start.

set -euo pipefail

SINCE_TAG=${1:-}
if [ -z "$SINCE_TAG" ]; then
  SINCE_TAG=$(git describe --tags --abbrev=0 2>/dev/null || true)
fi

RANGE="${SINCE_TAG}..HEAD"
if [ -z "$SINCE_TAG" ]; then
  RANGE="HEAD"
fi

echo "Collecting commits for range: $RANGE"

git log --pretty=format:'%h %s' $RANGE | sed 's/^/- /'

# Provide a preview of entries that would be prepended
echo
echo "Preview of entries to be added to CHANGELOG.md"

git log --pretty=format:'%s%n%b---END---' $RANGE | sed '/^chore: add initial CHANGELOG.md$/d' | tr '\n' '|' | sed 's/---END---/|/g' | tr '|' '\n' | awk 'NF' | sed 's/^/- /' | sed -n '1,100p'
