BINARY_NAME=swiftnessApp

build:
	@go mod vendor
	@echo "Building swiftness..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Swiftness built!"

run: build
	@echo "Starting swiftness..."
	@./tmp/${BINARY_NAME} &
	@echo "swiftness started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping swiftness..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped swiftness!"

restart: stop start