echo "SQLC install..."
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

echo "Generating SQLC..."
sqlc generate

echo "Be happy!"