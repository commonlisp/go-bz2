GOPATH=$(PWD)
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

cmd: 
	$(MAKE) -C src/cmd

all: cmd

clean:
	$(GOCLEAN) 
