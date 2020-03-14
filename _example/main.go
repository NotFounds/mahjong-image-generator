package main

import (
	"log"

	"github.com/disintegration/imaging"
	"github.com/notfounds/mahjong-image-generator"
)

func main() {
	tiles, err := mahjong.NewTiles()
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	img := tiles.GenerateFromString("19m19p19sESWNPFC1m")
	err = imaging.Save(img, "tiles.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
