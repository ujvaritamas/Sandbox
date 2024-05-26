Steps:

create a dev container and mount this dir (target /usr/src/app)
docker run -d --name <name> -v <srcpath>:<targetpath> <imagename>

cd webserver

go run yourmodule/main
go run example.com/webserver/main

go build -o webserver ./main