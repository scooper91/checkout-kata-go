.SILENT:

test:
	docker run --rm -t \
		-v $$PWD/pkg:/pkg -w /pkg \
		golang:alpine \
		sh -c 'go build -v && go test'
.PHONY: test
