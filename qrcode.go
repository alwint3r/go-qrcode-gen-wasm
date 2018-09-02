package main

import (
	"encoding/base64"
	"log"
	"syscall/js"

	"github.com/skip2/go-qrcode"
)

func main() {
	c := make(chan struct{}, 0)

	var generateQRCode js.Callback
	document := js.Global().Get("document")

	generateQRCode = js.NewCallback(func(args []js.Value) {
		data := document.Call("getElementById", "input-text").Get("value").String()
		if len(data) > 0 {
			log.Println("Building QR Code for", data)

			png, err := qrcode.Encode(data, qrcode.Medium, 256)
			if err != nil {
				log.Println("Failed generating QR Code. Error=", err.Error())
			}

			b64Encoded := base64.StdEncoding.EncodeToString(png)
			image := document.Call("getElementById", "qr-result")
			image.Call("setAttribute", "src", "data:image/png;base64, "+b64Encoded)
		}
	})

	document.Call("getElementById", "generate-qr").Call("addEventListener", "click", generateQRCode)
	<-c
}
