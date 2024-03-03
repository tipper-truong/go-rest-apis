# Golang REST APIs

## Instructions 

### Initialize Golang package
go mod init 
### Run Postgres DB Container Locally
docker run --name comments-api-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
