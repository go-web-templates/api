.PHONY: all clean

# Go cmd projects
SRC_DIR := cmd

# Binaries out dir
BIN_DIR := bin

# Cmd dirs list
SUBDIRS := $(wildcard $(SRC_DIR)/*)

# Binaries list based on cmd dirs
BINARIES := $(patsubst $(SRC_DIR)/%,$(BIN_DIR)/%,$(SUBDIRS))

# Build all and copy the runtime required files to the $BIN_DIR
all: $(BINARIES) copy_output create_init_script

clean:
	rm -rf $(BIN_DIR)

migrations-setup: all
	./bin/migrate setup
	
migrations-up: all
	./bin/migrate up
	
migrations-down: all
	./bin/migrate down

$(BIN_DIR)/%: $(SRC_DIR)/%
	@mkdir -p $(BIN_DIR)
	go build -o $@ $</main.go

copy_output:
	@mkdir -p $(BIN_DIR)
	cp ./app-conf.yml $(BIN_DIR)/app-conf.yml
	cp ./app-conf-dev.yml $(BIN_DIR)/app-conf-dev.yml
	cp -rf ./migrations $(BIN_DIR)/migrations

# Create an init script to docker
define INIT_SCRIPT
#! /bin/sh
set -e
./migrate setup
./migrate up
./app
endef
export INIT_SCRIPT

create_init_script:
	@echo "$$INIT_SCRIPT" > $(BIN_DIR)/init.sh
	chmod +x $(BIN_DIR)/init.sh
