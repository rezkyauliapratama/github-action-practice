registryURL = asia.gcr.io/$(project_id)/public-ip-service
tag = latest

install-dep:
	go mod tidy

run:
	go run main.go

build-docker:
	docker build -t $(registryURL):$(tag) .

push-docker:
	docker push $(registryURL):$(tag)

clear-docker:
	docker rmi $(registryURL):$(tag)