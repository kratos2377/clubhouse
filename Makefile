run:
	go run main.go

schema:
	@read -p "Enter Schema Name: " name; \
		go run entgo.io/ent/cmd/ent init $$name

generate:
	go generate ./ent

mac:
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/backend

linux:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/linux/backend

windows:
	env GOOS=windows GOARCH=amd64 go build -o ./bin/windows/backend