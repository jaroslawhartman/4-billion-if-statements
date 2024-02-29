all: template build

template:
	go run cmd/generator/generator.go 75000 > cmd/oddeven/oddeven.go

build:
	time go build cmd/oddeven/oddeven.go