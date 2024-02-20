echo "Getting Postgres driver..."
go get github.com/lib/pq

echo "Installing SQLC..."
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

echo "Getting SQL package..."
go get github.com/jackc/pgx/v5

echo "Generating SQLC..."
sqlc generate

echo "Be happy!"