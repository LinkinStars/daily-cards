.PHONY: ui go all

rollback:
	@mv ui.go ./static/

ui:
	@mv static/ui.go .
	@cd ui && yarn && yarn build --emptyOutDir && cd ..
	@mv ./ui.go static/ui.go

go:
	@go build -o dc cmd/dc/main.go

linux:
	@GOOS=linux GOARCH=amd64 go build -o dc cmd/dc/main.go

all: ui go
