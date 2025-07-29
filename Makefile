lint:
	go mod verify
	go vet ./...
	staticcheck ./...

format:
	goimports -local github.com/TimBerk/gophKeeper -v -w .

run:
	go run ./cmd/server

# Work with DB container
dbu:
	docker-compose up -d
dbd:
	docker-compose -f docker-compose.yml down
	docker volume rm goph-keeper_db
