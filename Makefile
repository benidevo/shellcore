fmt:
	go fmt ./...

run-shell:
	go run cmd/myshell/main.go

test:
	codecrafters test
