all: myecho helloworld

myecho: myecho.go
	gccgo myecho.go -o myecho

helloworld: helloworld.go
	gccgo helloworld.go -o helloworld


