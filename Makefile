BUILD_ROOT=${PWD}/build

chat-app: 
	@mkdir -p $(BUILD_ROOT)
	CGO_ENABLED=0 GOOS=linux go build \
		-a \
		-installsuffix cgo \
		-o $(BUILD_ROOT)/chat-app \
		github.com/ShariUllas/chat-app/cmd/chat-app

chat-app-mac:
	@mkdir -p $(BUILD_ROOT)
	go build \
		-a \
		-installsuffix cgo \
		-o $(BUILD_ROOT)/chat-app \
		github.com/ShariUllas/chat-app/cmd/chat-app

image:
	docker build -f Dockerfile -t chat-app:latest .

deploy:
	kubectl apply -f postgres-configmap.yaml
	kubectl apply -f postgres-storage.yaml
	kubectl apply -f postgres-deployment.yaml
	kubectl apply -f postgres-service.yaml

	kubectl exec -it [pod-name] --  psql -h localhost -U admin --password -p [port] postgresdb