

run: build
	@./bin/website_go-htmx-temple

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/air-verse/air@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@go clean --cache
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

css:
	npx tailwindcss -i view/css/app.css -o public/styles.css --watch
templ:
	@templ generate --watch --proxy=http://localhost:3000

build:
	@templ generate view
	@go build -tags prod -o /website_go-htmx-temple .

