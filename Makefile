all: myecho helloworld http-hello

myecho: myecho.go
	gccgo myecho.go -o myecho

helloworld: helloworld.go
	gccgo helloworld.go -o helloworld

http-hello: http-hello.go
	gccgo http-hello.go  -o http-hello

setup:
	sudo apt install -y golang
