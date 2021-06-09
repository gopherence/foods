# Build
build-all: build build-static build-stripped build-static-stripped build-linux-arm build-linux-arm64

build:
	go build -o bin/foods cmd/main.go

build-static:
	CGO_ENABLED=0 go build -o bin/foods-static cmd/main.go

build-stripped:
	go build -ldflags="-s -w" -o bin/foods-stripped cmd/main.go

build-static-stripped:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/foods-static-stripped cmd/main.go

build-linux-arm:
	GOOS=linux GOARCH=arm go build -o bin/foods-linux-arm cmd/main.go

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/foods-linux-arm64 cmd/main.go

# Run
run: build
	bin/foods

run-linux-arm: build-linux-arm
	bin/foods-linux-arm

run-linux-arm64: build-linux-arm64
	bin/foods-linux-arm64

# Cleanup
reset:
	echo "[]" > var/data/foods/content

clean: reset
	rm -f bin/foods*

# Docker
build-docker:
	docker build -t gopherence.foods .

run-docker:
	docker run --rm -p 8080:8080 gopherence.foods

# Deploy
deploy-run:
	ssh ubuntu@foods.gopherence.org 'killall foods'
	scp bin/foods-stripped ubuntu@foods.gopherence.org:foods/bin/foods
	scp var/www/index.aws.html ubuntu@foods.gopherence.org:foods/var/www/index.html
	ssh ubuntu@foods.gopherence.org 'cd foods && bin/foods'