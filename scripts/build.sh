#!/bin/sh
BASE_DIR="$(dirname "$(dirname "$0")")"
DIST_DIR="$BASE_DIR/dist"
echo "Building to $DIST_DIR"
PLATFORMS="darwin-arm64 linux-amd64 windows-amd64"
for platform in $PLATFORMS; do
	echo "Building for $platform"
	goos="$(echo "$platform" | cut -d - -f 1)"
	goarch="$(echo "$platform" | cut -d - -f 2)"
	output="$DIST_DIR/boop-$platform"
	if [ "$goos" = "windows" ]; then
		output="$output.exe"
	fi
	GOOS="$goos" GOARCH="$goarch" go build -o "$output" "$BASE_DIR/boop.go"
done
