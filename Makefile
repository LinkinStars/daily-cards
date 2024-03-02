.PHONY: ui go all

rollback:
	@mv ui.go ./static/

ui:
	@mv static/ui.go .
	@cd ui && yarn && yarn build --emptyOutDir && cd ..
	@mv ./ui.go static/ui.go

move-dist:
	@rm -rf ./ui/public/dist
	@cp -r ./ui/node_modules/vditor/dist ./ui/public

go:
	@go build -o dc cmd/dc/main.go

linux:
	@GOOS=linux GOARCH=amd64 go build -o dc cmd/dc/main.go

all: ui go
