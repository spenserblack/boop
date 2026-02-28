#!/bin/sh
SCRIPT_DIR="$(dirname "$0")"
BASE_DIR="$(dirname "$SCRIPT_DIR")"
DIST_DIR="$BASE_DIR/dist"
if ! command -v gh > /dev/null 2>&1; then
	echo "GitHub CLI (gh) required" >&2
	exit 1
fi

# NOTE: Ensure that assets are built
# NOTE: Does not ensure that the `dist/` directory is clean. It may contain outdated
#		assets from an old build.
echo "Running build script..."
"$SCRIPT_DIR/build.sh"

echo "Creating draft release..."
gh release create --draft --generate-notes "v0.0.0" "$DIST_DIR/*.tar.gz"
