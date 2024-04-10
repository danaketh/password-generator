ARTIFACT_NAME := password-generator

build:
	@go build -o bin/${ARTIFACT_NAME}/${ARTIFACT_NAME}.exe cmd/${ARTIFACT_NAME}/main.go

run:
	@go run cmd/${ARTIFACT_NAME}/main.go
