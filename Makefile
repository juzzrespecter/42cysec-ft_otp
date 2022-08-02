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
	@echo ${GREEN} "[build] ${NAME} built successfully!" ${END}

run:
	@echo ${BLUE} "[run] running ${NAME}" ${END}
	@go run .

clean:
	@go clean
	@echo ${RED} "[clean] cleaning cache..." ${END}
	@rm -f $(NAME)
	@echo ${RED} "[clean] removed ${NAME}" ${END}

test:
	@echo ${BLUE} "[test] running tests for ${NAME}..." ${END}
	@go test
