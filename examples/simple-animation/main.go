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

	fileFrame1, err := os.Open("frame1.png")
	if err != nil {
		log.Fatal(err)
	}
	fileFrame2, err := os.Open("frame2.png")
	if err != nil {
		log.Fatal(err)
	}
	fileFrame3, err := os.Open("frame3.png")
	if err != nil {
		log.Fatal(err)
	}
	fileFrame4, err := os.Open("frame4.png")
	if err != nil {
		log.Fatal(err)
	}
	fileFrame5, err := os.Open("frame5.png")
	if err != nil {
		log.Fatal(err)
	}

	frame1, _, err := image.Decode(fileFrame1)
	if err != nil {
		log.Fatal(err)
	}
	frame2, _, err := image.Decode(fileFrame2)
	if err != nil {
		log.Fatal(err)
	}
	frame3, _, err := image.Decode(fileFrame3)
	if err != nil {
		log.Fatal(err)
	}
	frame4, _, err := image.Decode(fileFrame4)
	if err != nil {
		log.Fatal(err)
	}
	frame5, _, err := image.Decode(fileFrame5)
	if err != nil {
		log.Fatal(err)
	}

	webpanim := webpanimation.NewWebpAnimation(1062, 938, 0) // Create 500x500 animaton with loop count 2, pass 0 for endless loop
	webpanim.WebPAnimEncoderOptions.SetKmin(9)
	webpanim.WebPAnimEncoderOptions.SetKmax(17)
	defer webpanim.ReleaseMemory() // dont forget call this or you will have memory leaks
	webpConfig := webpanimation.NewWebpConfig()
	webpConfig.SetLossless(1)

	err = webpanim.AddFrame(frame1, 0, webpConfig) // first frame
	if err != nil {
		log.Fatal(err)
	}
	err = webpanim.AddFrame(frame2, 1000, webpConfig) // first frame
	if err != nil {
		log.Fatal(err)
	}
	err = webpanim.AddFrame(frame3, 2000, webpConfig) // first frame
	if err != nil {
		log.Fatal(err)
	}
	err = webpanim.AddFrame(frame4, 3000, webpConfig) // first frame
	if err != nil {
		log.Fatal(err)
	}
	err = webpanim.AddFrame(frame5, 4000, webpConfig) // first frame
	if err != nil {
		log.Fatal(err)
	}

	err = webpanim.AddFrame(nil, 5000, webpConfig) // end loop on 6 second and start over
	if err != nil {
		log.Fatal(err)
	}

	err = webpanim.Encode(&buf) // encode animation and write result bytes in buffer
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("animation.webp", buf.Bytes(), 0777) // write bytes on disk
}
