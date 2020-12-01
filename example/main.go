package main

import (
	"flag"
	"os"

	"github.com/zachwe/govosk"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "file to transcribe")
	flag.Parse()
	model := govosk.NewModel("vosk-model-en-us/daanzu-20200905-lgraph")
	rec := govosk.NewRecognizer(model)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}

	vosk.VoskFinalResult(rec, buffer)
}
