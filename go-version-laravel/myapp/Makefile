BINARY_NAME=celeritasApp

build:
	@go mod vendor
	@echo "Building Cele..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Cele built!"

run: build
	@echo "Starting Cele..."
	@./tmp/${BINARY_NAME} &
	@echo "Cele started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Cele ..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Cele ...!"


restart: stop start