.PHONY: build run clean test

# ~~ aesthetica ~~ #
RED="\e[31m"
BLUE="\e[36m"
GREEN="\e[32m"
END="\e[0m"
# ~~     **     ~~ #

DIR	= ./totp/
NAME	= ft_otp

all: test vet fmt build

test:
	@echo ${BLUE} "[test] " ${END} "running tests for ${NAME}..."
	@go mod tidy
	@go test -v ./...

vet:
	@echo ${BLUE} "[vet] " ${END} "vetting..." 

fmt:
	@echo ${BLUE} "[fmt] " ${END} "formatting files..."
	@go list -f '{{.Dir}}' ./... | xargs -L1 gofmt -l

build:
	@go build -o $(NAME) ${DIR}
	@echo ${GREEN} "[build] " ${END} "✨ ${NAME} built successfully! ✨"

clean:
	@go clean
	@echo ${RED} "[clean] " ${END} "cleaning cache..."
	@rm -f $(NAME)
	@echo ${RED} "[clean] " ${END} "removed ${NAME}"
