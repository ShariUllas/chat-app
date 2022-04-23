BUILD_ROOT=${PWD}/build

chat-app: 
	@mkdir -p $(BUILD_ROOT)
	CGO_ENABLED=0 GOOS=linux go build \
		-a \
		-installsuffix cgo \
		-o $(BUILD_ROOT)/chat-app \
		github.com/ShariUllas/chat-app/cmd/chat-app

image:
	docker build -f Dockerfile -t chat-app:latest .

deploy:
	kubectl apply -f ./resource-manifests
