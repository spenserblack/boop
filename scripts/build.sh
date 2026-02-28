#!/bin/sh
SCRIPT_DIR="$(dirname "$0")"
BASE_DIR="$(dirname "$SCRIPT_DIR")"
DIST_DIR="$BASE_DIR/dist"
echo "Building to $DIST_DIR"
PLATFORMS="darwin-arm64 linux-amd64 windows-amd64"
for platform in $PLATFORMS; do
	echo "Building for $platform"
	goos="$(echo "$platform" | cut -d - -f 1)"
	goarch="$(echo "$platform" | cut -d - -f 2)"
	filename="boop-$platform"
	if [ "$goos" = "windows" ]; then
		filename="$filename.exe"
	fi
	target="$DIST_DIR/$filename"
	GOOS="$goos" GOARCH="$goarch" go build -o "$target" "$BASE_DIR/boop.go"

	tar -C "$DIST_DIR" --remove-files -czf "$DIST_DIR/$filename.tar.gz" "$filename"
done
