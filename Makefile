PRODUCT=fuse
GOOS=linux
GOARCH=amd64
NAME=$(PRODUCT)-$(GOOS)-$(GOARCH)$(EXT)
EXT=
ifeq ($(GOOS),windows)
	override EXT=.exe
endif

# alpine image doesn't have git = no vcs info
IMAGE=golang:1.24.5
DOCKER=docker run -t --rm \
		-u $$(id -u):$$(id -g) \
		-v $$(pwd):$$(pwd) \
		-w $$(pwd) \
		-e GOCACHE=/tmp \
		-e CGO_ENABLED=0 \
		-e GOOS=$(GOOS)\
		-e GOARCH=$(GOARCH) \
		$(IMAGE)

clean:
	-rm handlers/testdb.dat*

test: clean
	 $(DOCKER) go test -v ./...

build:
	$(DOCKER) go build -trimpath \
				-buildvcs=true \
				-o $(NAME)

gen:
	-cd seedgen && go run main.go

release: test gen
	$(MAKE) GOOS=linux build
	$(MAKE) GOOS=windows build

.DEFAULT_GOAL := release
