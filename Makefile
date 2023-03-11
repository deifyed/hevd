.PHONY: run build test clean

run:
	echo "This target must always start a development execution of the software"

build:
	echo "This target must always build a production version of the software"

test:
	find . -name '*.go' | entr -r go run run main.go docs/examples/minimal-curl.yaml
test-watch:
	find . -name '*.go' | entr -r make test

clean:
	echo "This target must always clean up any temporary files created by the build process"

fmt:
	echo "This target must always format the code according to project standards. This target must modify the source code in place."

fmt-ci:
	echo "This target must return an exit code of 0 if there is nothing to format, and 1 if there is something to format. This target must not modify the source code in place."

lint:
	echo "This target must always lint the code according to project standards. This target must modify the source code in place."

lint-ci:
	echo "This target must return an exit code of 0 if there is nothing to lint, and 1 if there is something to lint. This target must not modify the source code in place."

ci: fmt-ci lint-ci test
