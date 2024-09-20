# run program without compile
dev:
	go run ./cmd/golps/main.go

# compile program
build:
	rm -rf ./build
	go build -o ./build/golps ./cmd/golps

# run compiled program
start:
	./build/golps

# clean compiled program
clean:
	rm -rf ./build