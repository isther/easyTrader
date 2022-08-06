build:
	@echo "Start Building..."
	@go build
	@echo "End Building..."

run: build
	@echo "Start Run..."
	@./easyTrader

clear:
	@echo "Clear..."
	@rm easyTrader
	@rm -rf logs

docker-compose-up: 
	@docker-compose up -d

docker-compose-down:
	@docker-compse down
