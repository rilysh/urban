compile:
	go build -ldflags "-w" urban.go

cleanup:
	rm -rf urban
