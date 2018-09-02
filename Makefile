DEST_DIR := server
OUTPUT_NAME := qrcode.wasm

all:
	GOOS=js GOARCH=wasm go build -o $(DEST_DIR)/$(OUTPUT_NAME)
