package main

import (
	"fmt"
	"io"
	"syscall/js"

	"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type bencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}

func Open(r io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func initalizeLoadingPage() {

	document := js.Global().Get("document")
	document.Get("body").Call("insertAdjacentHTML", "beforeend", `
		<input type="file" id="fileInput">
		<output id="fileOutput"></output>
	`)

	fileInput := document.Call("getElementById", "fileInput")

	fileInput.Set("oninput", js.FuncOf(func(v js.Value, x []js.Value) any {
		fileInput.Get("files").Call("item", 0).Call("arrayBuffer").Call("then", js.FuncOf(func(v js.Value, x []js.Value) any {
			data := js.Global().Get("Uint8Array").New(x[0])
			dst := make([]byte, data.Get("length").Int())
			js.CopyBytesToGo(dst, data)
			go printer(dst)
			return nil
		}))
		return nil
	}))

}
func printer(data []byte) {
	fmt.Println(data)

}

func main() {

	initalizeLoadingPage()

	// It's a requirement, go WASM execution stack always waits
	// Do not delete that
	select {}
}
