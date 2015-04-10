package main

import (
	"flag"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	//Flags for catch the inputs when exc the binary
	inputDir := flag.String("in", "./", "Input Directory")
	output := flag.String("out", "animation.gif", "Output")
    delay := flag.Int("delay", 100, "Delay Between images")
	flag.Parse()

	//Read Files from input directory
	files, err := ioutil.ReadDir(*inputDir)
	if err != nil {
		panic(err.Error())
	}
	images := []*image.Paletted{}
	delays := []int{}

	//Loop adding each image for a list of images
	for _, file := range files {

		if !strings.HasSuffix(file.Name(), ".png") && !strings.HasSuffix(file.Name(), ".jpeg") && !strings.HasSuffix(file.Name(), ".jpg") {
			continue
		}

		reader, err := os.Open(filepath.Join(*inputDir, file.Name()))
		if err != nil {
			panic(err.Error())
		}
		img, _, err := image.Decode(reader)
		if err != nil {
			panic(err.Error())
		}

		// Draw de image paletted and append to images list
		b := img.Bounds()
		drawer := draw.FloydSteinberg
		pm := image.NewPaletted(b, palette.Plan9[:256])
		drawer.Draw(pm, b, img, image.ZP)
		images = append(images, pm)
		delays = append(delays, *delay)
	}
	writer, err := os.Create(*output) //Where i will save
	if err != nil {
		panic(err.Error())
	}

	// Write the list images to outputfile.gif
	err = gif.EncodeAll(writer, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		panic(err.Error())
	}

}
