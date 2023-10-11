GOARCH   := amd64
GOOS     := linux

SRC_FILE := cmd/kadrion/main.go
BIN_DIR  := bin
INSTALL_DIR := /opt/kadrion

BINARY   := kadrion

.PHONY: all build install

all: build install clean

build: $(BIN_DIR)/$(BINARY)

install: $(INSTALL_DIR)/$(BINARY)
	@echo 'export PATH="$$PATH:$(INSTALL_DIR)"' >> ~/.bashrc
	source ~/.bashrc

clean:
	@echo "Removing kadrion..."
	@sed -i '/\/opt\/kadrion/d' ~/.bashrc
	@echo "Removing /opt/kadrion..."
	sudo rm -rf $(INSTALL_DIR)


$(BIN_DIR)/$(BINARY): $(SRC_FILE)
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $@ $(SRC_FILE)

$(INSTALL_DIR)/$(BINARY): $(BIN_DIR)/$(BINARY)
	sudo mkdir -p $(INSTALL_DIR)
	sudo cp $< $@
