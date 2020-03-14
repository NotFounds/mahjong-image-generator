package mahjong

import (
	"image"
	"image/color"
	"regexp"

	"github.com/disintegration/imaging"
)

var (
	background = color.NRGBA{0, 0, 0, 0}
	width      = 154
	height     = 220
)

type Tile struct {
	t string
	i string
}

type Tiles struct {
	Baseimage image.Image
}

func NewTiles() (*Tiles, error) {
	f, _ := Assets.Open("/images/tiles.png")
	defer f.Close()
	img, err := imaging.Decode(f, imaging.AutoOrientation(true))
	if err != nil {
		return nil, err
	}
	return &Tiles{img}, nil
}

func (tiles Tiles) GetTile(t Tile) image.Image {
	offsetX := map[string]int{
		"1": 0,
		"2": 166,
		"3": 332,
		"4": 497,
		"5": 663,
		"6": 829,
		"7": 995,
		"8": 1160,
		"9": 1326,
		"E": 0,
		"S": 166,
		"W": 332,
		"N": 497,
		"C": 663,
		"F": 829,
		"P": 995,
	}
	offsetY := map[string]int{
		"m": 0,
		"p": 261,
		"s": 522,
		"E": 783,
		"S": 783,
		"W": 783,
		"N": 783,
		"C": 783,
		"F": 783,
		"P": 783,
	}

	x := offsetX[t.i]
	y := offsetY[t.t]

	return imaging.Crop(tiles.Baseimage, image.Rect(x, y, width+x, height+y))
}

func (tiles Tiles) Generate(pai []Tile) image.Image {
	dst := imaging.New((width)*len(pai), height, background)
	for i, p := range pai {
		dst = imaging.Paste(dst, tiles.GetTile(p), image.Pt(width*i, 0))
	}
	return dst
}

func (tiles Tiles) GenerateFromString(query string) image.Image {
	return tiles.Generate(Parse(query))
}

func Parse(query string) []Tile {
	reg := regexp.MustCompile(`([1-9]+(m|p|s)|(E|S|W|N|P|F|C))`)
	res := reg.FindAllString(query, -1)

	tiles := []Tile{}
	for _, p := range res {
		length := len(p)
		if length >= 2 {
			t := string([]rune(p[length-1:]))
			for _, n := range []rune(p[0 : length-1]) {
				tiles = append(tiles, Tile{t, string(n)})
			}
		} else {
			tiles = append(tiles, Tile{p, p})
		}
	}
	return tiles
}
