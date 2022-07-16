# Build

help:
	@echo "See Makefile"
tidy:
	@bash -c "go mod tidy"
run:
	@bash -c "go run example/main.go"
scan:
	@script/gosec.sh