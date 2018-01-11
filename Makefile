GO_FLAGS=-a -ldflags '-extldflags "-static"'



go-template: go-template.go
	CGO_ENABLED=0 GOOS=linux go build $(GO_FLAGS) -o $@ $^

.PHONY: clean
clean:
	rm go-template

