QR Code Generator using WebAssembly
==================================

QR Code generator written in Go and compiled to WebAssembly.

## How It works

Take user's input and generate the QR code bytes and display the result as a base64-encoded PNG image.

## Running it yourself

Clone this repository, then compile the code to WASM using this command

```sh
make
```

Then go to `server` directory and run the server

```sh
go run server.go
```

Then go to http://localhost:8000/wasm_exec.html

## Useful Resources

* [QR Code library](https://github.com/skip2/go-qrcode)

* [Quick tutorial: Write Go, run WASM!](https://dev.to/cia_rana/quick-tutorial-write-go-run-wasm-2ilf)

* [syscall/js documentation](https://golang.org/pkg/syscall/js/)

## License

This code is licensed under MIT License, see LICENSE file.

`server/wasm_exec.html` and `server/wasm_exec.js` is licensed under [this license](https://github.com/golang/go/blob/master/LICENSE).

