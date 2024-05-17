builddb:
					docker-compose up --build

start:
					go run cmd/blog/main.go

test:		
					go test -v -cover ./...
						