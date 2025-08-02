user=alice

lint:
	go mod verify
	go vet ./...
	staticcheck ./...
	golangci-lint run ./...

format:
	goimports -local github.com/TimBerk/gophKeeper -v -w .
	golangci-lint fmt .

run:
	go run ./cmd/server

login:
	go run ./cmd/cli login -u $(user)
list:
	go run ./cmd/cli list
sync:
	go run ./cmd/cli sync
version:
	go run ./cmd/cli version

# Work with DB container
dbu:
	docker-compose up -d
dbd:
	docker-compose -f docker-compose.yml down
	docker volume rm goph-keeper_db
