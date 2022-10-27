compile:
	go build -ldflags "-w" urban.go

install:
	cp -f ./urban /usr/bin/urban
	chmod 755 /usr/bin/urban

cleanup:
	rm -rf urban

