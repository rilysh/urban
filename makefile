compile:
	go build -ldflags "-w" urban.go
	chmod -R 755 /usr/bin/urban

install:
	cp -f ./urban /usr/bin/urban
	chmod +x /usr/bin/urban

cleanup:
	rm -rf urban

