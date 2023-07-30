swag: swag-fmt
	swag init --parseDependency --parseInternal --parseDepth 1 -d cmd/,pkg/api/feed/

swag-fmt:
	swag fmt -d cmd/,pkg/api/feed/

build:
	go build -C cmd -o newsfeed

run:
	go run cmd/main.go