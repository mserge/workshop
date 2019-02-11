APP?=tracking-workshop

.PHONY: env-up
env-up:
	@CURRENT_UID=$(id -u):$(id -g) docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down

.PHONY: clean
clean:
	@echo "+ $@"
	rm -rf bin/${APP}

.PHONY: build
build: clean
	@echo "+ $@"
	go build -o bin/${APP} ./cmd/tracking