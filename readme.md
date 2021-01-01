# webpanimation  
Packge is binding to libwebp v1.1 providing methods to create webp animations from golang `image.Image` interface
## Installing
`go get github.com/sizeofint/webpanimation`


## Example
```
package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/sizeofint/webpanimation"
)

func main() {
	var buf bytes.Buffer

	fileReader1, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}

	fileReader2, err := os.Open("test2.jpg")
	if err != nil {
		log.Fatal(err)
	}

	jpgImg, _, err := image.Decode(fileReader1)
	if err != nil {
		log.Fatal(err)
	}

	jpgImg2, _, err := image.Decode(fileReader2)
	if err != nil {
		log.Fatal(err)
	}

	webpanim := webpanimation.NewWebpAnimation(500, 500, 2) // Create 500x500 animaton with loop count 2, pass 0 for endless loop
	defer webpanim.ReleaseMemory()                          // dont forget call this or you will have memory leaks
	webpConfig := webpanimation.NewWebpConfig()
	webpConfig.SetLossless(1)

	err = webpanim.AddFrame(jpgImg, 0, webpConfig) // first frame
	if err != nil {
		log.Fatal(err)
	}

	err = webpanim.AddFrame(jpgImg2, 3000, webpConfig) // second frame after 3 second
	if err != nil {
		log.Fatal(err)
	}

	err = webpanim.AddFrame(nil, 6000, webpConfig) // end loop on 6 second and start over
	if err != nil {
		log.Fatal(err)
	}

	err = webpanim.Encode(&buf) // encode animation and write result bytes in buffer
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("animation.webp", buf.Bytes(), 0777) // write bytes on disk
}

```
## Dependencies
The only dependency is libwebp v1.1, it source code are embeded in package so no additional installations are needed on machine.