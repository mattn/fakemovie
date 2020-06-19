package main

import (
	"flag"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/mattn/fakemovie"
)

func main() {
	var out string
	var r int
	flag.StringVar(&out, "o", "output.png", "output filename")
	flag.IntVar(&r, "r", -1, "radius of button")
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	img = fakemovie.Fake(img, r)
	f, err = os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}
