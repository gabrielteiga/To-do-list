echo "Installing Postgres driver..."
go get github.com/lib/pq

echo "Installing SQLC..."
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

echo "Generating SQLC..."
sqlc generate

echo "Be happy!"