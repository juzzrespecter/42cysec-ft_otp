.PHONY: build run clean test

NAME	= ft_otp

build:
	go build -o $(NAME) .

run:
	go run .

clean:
	go clean
	rm -f $(NAME)

test:
	go test
