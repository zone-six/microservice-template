database:
	echo 'CREATE DATABASE zone_six;' | docker exec -i infrastructure_db_1 psql -h 127.0.1.1 -U postgres    

drop_database:
	echo 'DROP DATABASE zone_six;' | docker exec -i infrastructure_db_1 psql -h 127.0.1.1 -U postgres 

run_server: database
	go run internal/cmd/server/server.go

generate_rest_client:
	rm -rf ./internal/clients/rest/restapi ./internal/clients/rest/models 
	swagger generate server -A identity-api -f ./swagger.yml -t ./internal/clients/rest --template=stratoscale
	go get -u -f ./internal/clients/rest/...
	go mod tidy