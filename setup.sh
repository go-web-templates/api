#! /usr/bin/env bash

# replace all module references
CURR_MODULE='github.com/go-web-templates/api'
NEW_MODULE=$1

echo "Replacing the module name"
find . -type f -name "*.go" -print0 | \
  xargs -0 sed -i '' -e "s#${CURR_MODULE}#${NEW_MODULE}#g" \
  2>/dev/null

echo "Deleting old go module files"
rm -rf go.mod go.sum

echo "Creating a new go module"
go mod init ${NEW_MODULE}
go mod tidy
