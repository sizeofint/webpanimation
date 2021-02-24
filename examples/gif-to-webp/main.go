package main

import (
	"bytes"
	"github.com/sizeofint/webpanimation"
	"image/gif"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var buf bytes.Buffer
	gifFile, err := os.Open("animation.gif")
	if err != nil {
		log.Fatal(err)
	}
	gif, err := gif.DecodeAll(gifFile)
	if err != nil {
		log.Fatal(err)
	}

	webpanim := webpanimation.NewWebpAnimation(gif.Config.Width, gif.Config.Height, gif.LoopCount) // Create 500x500 animaton with loop count 2, pass 0 for endless loop
	webpanim.WebPAnimEncoderOptions.SetKmin(9)
	webpanim.WebPAnimEncoderOptions.SetKmax(17)
	defer webpanim.ReleaseMemory() // dont forget call this or you will have memory leaks
	webpConfig := webpanimation.NewWebpConfig()
	webpConfig.SetLossless(1)

	timeline := 0

	for i, img := range gif.Image {

		err = webpanim.AddFrame(img, timeline, webpConfig)
		if err != nil {
			log.Fatal(err)
		}
		timeline += gif.Delay[i] * 10
	}
	err = webpanim.AddFrame(nil, timeline, webpConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = webpanim.Encode(&buf) // encode animation and write result bytes in buffer
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("animation.webp", buf.Bytes(), 0777)
}
