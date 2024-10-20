# windows/x64 fallback because that is my only box that doesn't have uname
OS := $(shell uname -s 2>/dev/null || echo windows)
ARCH := $(shell uname -m 2>/dev/null || echo x64)

# account for windows with git shell extensions
ifeq ($(findstring MSYS_NT,$(OS)),MSYS_NT)
	OS := windows
endif

ifneq ($(OS),windows)
	OS := $(shell echo $(OS) | tr '[:upper:]' '[:lower:]' | sed 's/darwin/macos/')
	ARCH := $(shell echo $(ARCH) | tr '[:upper:]' '[:lower:]' | sed 's/x86_64/x64/')
endif

TAILWINDCSS_PKG := $(OS)-$(ARCH)
ifeq ($(OS),windows)
	TAILWINDCSS_PKG := $(TAILWINDCSS_PKG).exe
endif

help:
	@echo "The following commands are available:"
	@echo "  dev:          Run the web application in development mode using the Air tool for live reloading"
	@echo "  watch-css:    Watch CSS and template files for changes and rebuild using Tailwind CSS CLI"
	@echo "  build:        Build the web application for deployment and output the binary and required assets to the 'dist' directory"
	@echo "  build-css:    Build the production CSS using Tailwind CSS CLI"
	@echo "  tailwindcss:  Download the Tailwind CSS CLI based on the OS and architecture"

# Download the Tailwind CSS CLI based on the OS and architecture
tailwindcss:
	@echo "Downloading Tailwind CSS CLI..."
ifeq ($(OS),windows)
	@powershell -Command "Invoke-WebRequest -Uri https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$(TAILWINDCSS_PKG) -OutFile tailwindcss.exe"
else
	@curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$(TAILWINDCSS_PKG)
	@mv tailwindcss-$(TAILWINDCSS_PKG) tailwindcss
	@chmod +x tailwindcss
endif

# Build the production CSS using Tailwind CSS CLI
.PHONY: build-css
build-css: tailwindcss
	@echo "Building production CSS..."
	@./tailwindcss -i ./ui/styles/tailwind.css -o ./ui/static/css/style.css --minify

# Watch CSS files for changes and rebuild using Tailwind CSS CLI
.PHONY: watch-css
watch-css: tailwindcss
	@echo "Watching CSS files for changes..."
	@./tailwindcss -i ./ui/styles/tailwind.css -o ./ui/static/css/style.css --watch

# Clean the 'dist' directory
.PHONY: clean
clean:
	@echo "Cleaning the dist directory..."
ifeq ($(OS),windows)
	@powershell -Command "Remove-Item -Recurse -Force dist; New-Item -ItemType Directory -Path dist/ui"
	@powershell -Command "Copy-Item -Recurse -Force ui/static dist/static"
	@powershell -Command "Copy-Item -Recurse -Force data dist/data"
else
	@rm -rf dist
	@mkdir -p dist/ui
endif

# Build the web application and output the binary to the 'dist' directory
.PHONY: build
build: clean build-css
	@echo "Building web application..."
ifeq ($(OS),windows)
	@set GOOS=linux && set GOARCH=amd64 && go build -ldflags="-s" -o=./dist/web ./cmd/web
else
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o=./dist/web ./cmd/web
	@cp -r ui/static dist/ui
	@cp -r data dist/data
endif

# Run the web application in development mode using the Air tool for live reloading
.PHONY: dev
dev:
	@echo "Running the web application in development mode..."
	@echo "  NOTE: you should already be running 'make watch-css' in a separate terminal"
ifeq ($(OS),windows)
	@air -c .\.air-win.toml
else
	@air -c .air.toml
endif