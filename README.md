#Initial project
go get
go mod tidy


#Run test project
go test -v ./services_test/service_test.go


#Run docker image on localhost
docker build -t go-rest-api-image .
docker run -p 8080:8080 go-rest-api-image