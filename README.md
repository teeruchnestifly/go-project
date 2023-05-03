#Initial project
go get
go mod tidy
go run main.go

test--> open http://localhost:8080/books


#Run test project
go test -v ./services_test/service_test.go


#Run docker image on localhost
docker build -t go-rest-api-image .
docker run -p 8080:8080 go-rest-api-image

gcloud services enable cloudbuild.googleapis.com

gcloud builds submit --tag gcr.io/poc-nestifly/go-rest-api-image:1.0.0 .

kubectl create deployment go-rest --image=gcr.io/poc-nestifly/go-rest-api-image:1.0.0

kubectl expose deployment go-rest --type=LoadBalancer --port 80 --target-port 8080

kubectl scale deployment go-rest --replicas=3