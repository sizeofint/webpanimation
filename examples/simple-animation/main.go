package main

import (
	"bytes"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/sizeofint/webpanimation"
)

func main() {
	var buf bytes.Buffer
	var err error
	webpanim := webpanimation.NewWebpAnimation(1062, 938, 0)
	webpanim.WebPAnimEncoderOptions.SetKmin(9)
	webpanim.WebPAnimEncoderOptions.SetKmax(17)
	defer webpanim.ReleaseMemory() // don't forget call this or you will have memory leaks
	webpConfig := webpanimation.NewWebpConfig()
	webpConfig.SetLossless(1)

	pngFrames := []string{"frame1.png", "frame2.png", "frame3.png", "frame4.png", "frame5.png"}
	timeline := 0

	for _, f := range pngFrames {
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		frame, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

		err = webpanim.AddFrame(frame, timeline, webpConfig)
		if err != nil {
			log.Fatal(err)
		}
		timeline += 1000

	}

	err = webpanim.AddFrame(nil, timeline, webpConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = webpanim.Encode(&buf) // encode animation and write result bytes in buffer
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("animation.webp", buf.Bytes(), 0777) // write bytes on disk
}
