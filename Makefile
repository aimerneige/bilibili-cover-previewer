GO				:= go
CC				?= gcc
GO_SOURCES		:= $(shell find . -name "*.go")
WEB_SOURCES		:= $(shell find ./web/app -name "*") $(shell find ./web/static -name "*")
WEB_ROOT		?= ./web/
WEB_PORT		?= :54612
GOOS			?= linux
GOARCH			?= amd64

.PHONY: all clean

all: bin/bilibili-cover-previewer-$(GOOS)-$(GOARCH)

web/build/: $(WEB_SOURCES)
	echo $(WEB_SOURCES)
	# cd web && npm run build

bin/bilibili-cover-previewer-$(GOOS)-$(GOARCH): $(GO_SOURCES) web/build/
	-mkdir -p bin
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=1 CC=$(CC) \
	$(GO) build -o bin/bilibili-cover-previewer-$(GOOS)-$(GOARCH) \
	-ldflags "-X main.root=$(WEB_ROOT) -X main.port=$(WEB_PORT)" \
	cmd/bilibili-cover-previewer/*

# Rerun target every time a file change is detected in the current directory.
watch:
	while true; do \
		clear; \
        make --no-print-directory $(TARGET); \
        inotifywait -qre close_write .; \
    done

clean:
	-rm -rvf bin
