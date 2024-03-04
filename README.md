# Golang REST APIs

## Instructions 

### Initialize Golang package
go mod init github.com/tipper-truong/go-rest-apis

### Run Postgres DB Container Locally
docker run --name comments-api-db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

NOTE: To remove existing database container, run `docker rm <container-id>`

### Run setup_env_variables.sh to export database environment variables
source ./set_env_variables.sh
