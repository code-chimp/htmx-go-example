help:
	@echo "The following commands are available:"
	@echo "  run:             Run the web application using the Go command. (any)"
	@echo "  dev:             Run the web application in development mode using the Air tool for live reloading. (linux/mac)"
	@echo "  dev/win:         Run the web application in development mode on Windows using the Air tool for live reloading. (windows)"
	@echo "  build:           Build the web application and output the binary to the 'bin' directory. (linux)"
	@echo "  build/arm:       Build the web application for Linux running on an ARM architecture and output the binary to the 'bin' directory. (linux)"
	@echo "  build/mac:       Build the web application for Apple Silicon chips and output the binary to the 'bin' directory. (mac)"
	@echo "  build/mac/intel: Build the web application for Intel-based Macs and output the binary to the 'bin' directory. (mac)"
	@echo "  build/win:       Build the web application for Windows (amd64 architecture) and output the binary to the 'bin' directory. (windows)"

# run: Run the web application using the Go command.
run:
	go run ./cmd/web

# dev: Run the web application in development mode using the Air tool for live reloading
dev:
	air -c .air.toml

# dev/win: Run the web application in development mode on Windows using the Air tool for live reloading
dev/win:
	air -c .air-win.toml

# ======================
# Build
# ======================

# build: Build the web application and output the binary to the 'bin' directory.
.PHONY: build
build:
	@echo "Building web application..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o=./bin/linux_amd64/web ./cmd/web

# build/arm: Build the web application for Linux running on an ARM architecture and output the binary to the 'bin' directory.
.PHONY: build/arm
build/arm:
	@echo "Building web application for Linux running on an ARM architecture..."
	GOOS=linux GOARCH=arm64 go build -ldflags="-s" -o=./bin/linux_arm64/web ./cmd/web

# build/mac: Build the web application for Apple Silicon chips and output the binary to the 'bin' directory.
.PHONY: build/mac
build/mac:
	@echo "Building web application for Apple Silicon..."
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s" -o=./bin/darwin_arm64/web ./cmd/web

# build/mac/intel: Build the web application for Intel-based Macs and output the binary to the 'bin' directory.
.PHONY: build/mac/intel
build/mac/intel:
	@echo "Building web application for Apple Intel..."
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s" -o=./bin/darwin_amd64/web ./cmd/web

# build/win: Build the web application for Windows (amd64 architecture) and output the binary to the 'bin' directory.
.PHONY: build/win
build/win:
	@echo "Building web application for Windows..."
	GOOS=windows GOARCH=amd64 go build -ldflags="-s" -o=./bin/windows_amd64/web.exe ./cmd/web
