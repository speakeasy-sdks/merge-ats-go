.PHONY: *
all: speakeasy

speakeasy: check-speakeasy
	speakeasy generate sdk --lang go -o . -s ./openapi.json

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }
