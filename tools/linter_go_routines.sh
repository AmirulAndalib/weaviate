#!/usr/bin/env bash

set -euo pipefail
# Search for all tracked non-test .go files in the Git repository
all_files=$(git ls-files | grep -E '\.go$' | grep -vE '_test\.go$')

# Get all files with 'go ' in them, ignoring lines that start with a comment
files=$(grep -E -l '^[[:space:]]*[^/]*go ' ${all_files})

found_error=false

# Define array of excluded files with their explanations
excluded_files=(
    "entities/errors/go_wrapper.go"              # wrapper
    "adapters/handlers/rest/server.go"           # autogenerated
    "test/helper/sample-schema/books/books.go"   # test file
    # race condition in classification when replacing the direct goroutine with a wrapper. Not important enough to investigate as nobody is using classification.
    "usecases/classification/classifier_run.go"
    "usecases/auth/authorization/docs/generator.go" # docs generator
    "tools/dev/generate_release_notes/main.go" # generate release notes tool
)

# Check if file is in excluded list
for file in $files; do
    if [[ " ${excluded_files[@]} " =~ " ${file} " ]]; then
        continue
    fi
    
    echo "Error: $file uses direct goroutines. Please use entities/errors/go_wrapper.go instead."
    found_error=true
done

if [ "$found_error" = true ]; then
    exit 1
fi
