generate:
	@templ generate
	@go generate ./...
    
run:
	@if [ -n "$$(lsof -ti:1323)" ]; then kill -9 $$(lsof -ti:1323); fi
	@wgo -file=.go -file=.templ -xfile=_templ.go make generate :: go run main.go

build: generate 
	@go build -ldflags="-s -w" -o build/bin main.go
	@ls -lah build/bin | awk '{print "Location:" $$9, "Size:" $$5}' | column -t