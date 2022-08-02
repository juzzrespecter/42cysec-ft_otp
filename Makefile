.PHONY: build run clean test

# ~~ aesthetica ~~ #
RED="\e[31m"
BLUE="\e[36m"
GREEN="\e[32m"
END="\e[0m"
# ~~     **     ~~ #

NAME	= ft_otp

build:
	@go build -o $(NAME) .
	@echo ${GREEN} "[build] " ${END} "✨ ${NAME} built successfully! ✨"

clean:
	@go clean
	@echo ${RED} "[clean] " ${END} "cleaning cache..."
	@rm -f $(NAME)
	@echo ${RED} "[clean] " ${END} "removed ${NAME}"

test:
	@echo ${BLUE} "[test] " ${END} "running tests for ${NAME}..."
	@go test
