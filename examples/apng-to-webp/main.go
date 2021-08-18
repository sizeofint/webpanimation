package main

import (
	"bytes"
	"github.com/kettek/apng"
	"github.com/sizeofint/webpanimation"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var buf bytes.Buffer
	gifFile, err := os.Open("animation.png")
	if err != nil {
		log.Fatal(err)
	}
	png, err := apng.DecodeAll(gifFile)
	if err != nil {
		log.Fatal(err)
	}
	ihrdFrame := png.Frames[0].Image

	width := ihrdFrame.Bounds().Max.X
	height := ihrdFrame.Bounds().Max.Y
	webpanim := webpanimation.NewWebpAnimation(width, height, int(png.LoopCount))
	webpanim.WebPAnimEncoderOptions.SetKmin(9)
	webpanim.WebPAnimEncoderOptions.SetKmax(17)
	defer webpanim.ReleaseMemory() // dont forget call this or you will have memory leaks
	webpConfig := webpanimation.NewWebpConfig()
	webpConfig.SetLossless(1)

	ihrdRGBA := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(ihrdRGBA, ihrdRGBA.Bounds(), ihrdFrame, image.Point{}, draw.Src)

	timeline := 0
	outputBuffer := &image.RGBA{
		Stride: ihrdRGBA.Stride,
		Rect:   ihrdRGBA.Rect,
		Pix:    make([]uint8, len(ihrdRGBA.Pix)),
	}
	copy(outputBuffer.Pix, ihrdRGBA.Pix)

	for _, img := range png.Frames {
		if img.IsDefault {
			continue
		}
		var previousOutputBuffer *image.RGBA

		// backup previous output buffer, if we have to dispose current frame
		if img.DisposeOp == apng.DISPOSE_OP_PREVIOUS {
			previousOutputBuffer = &image.RGBA{
				Stride: outputBuffer.Stride,
				Rect:   outputBuffer.Rect,
				Pix:    make([]uint8, len(outputBuffer.Pix)),
			}
			copy(previousOutputBuffer.Pix, outputBuffer.Pix)
		}

		if img.BlendOp == apng.BLEND_OP_OVER {
			draw.Draw(outputBuffer, outputBuffer.Bounds(), img.Image, image.Pt(-img.XOffset, -img.YOffset), draw.Over)
		} else {
			draw.Draw(outputBuffer, outputBuffer.Bounds(), img.Image, image.Pt(-img.XOffset, -img.YOffset), draw.Src)
		}

		err = webpanim.AddFrame(outputBuffer, timeline, webpConfig)
		if err != nil {
			log.Fatal(err)
		}
		if img.DisposeOp == apng.DISPOSE_OP_BACKGROUND {
			outputBuffer = image.NewRGBA(image.Rect(0, 0, width, height))
		} else if img.DisposeOp == apng.DISPOSE_OP_PREVIOUS {
			outputBuffer = previousOutputBuffer
		}

		timeline += int(img.GetDelay() * 1000)
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
