database:
	echo 'CREATE DATABASE zone_six;' | docker exec -i infrastructure_db_1 psql -h 127.0.1.1 -U postgres    

drop_database:
	echo 'DROP DATABASE zone_six;' | docker exec -i infrastructure_db_1 psql -h 127.0.1.1 -U postgres 

run_server: database
	go run internal/cmd/server/server.go