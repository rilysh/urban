all:
	rm -rf ./build
	mkdir build
	make compile
	make tarpkg

compile:
	go build -ldflags "-w" urban.go
	mv urban ./build

cleanup:
	rm -rf ./build

tarpkg:
	tar cf ./build/urban.tar ./build/urban
	xz -z ./build/urban.tar