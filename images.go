package mahjong

import (
	"time"

	"github.com/jessevdk/go-assets"
)


// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"images"}, "/images": []string{"tiles.png"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1584177943, 1584177943062282100),
		Data:     nil,
	}, "/images": &assets.File{
		Path:     "/images",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1584167086, 1584167086415410000),
		Data:     nil,
	}, "/images/tiles.png": &assets.File{
		Path:     "/images/tiles.png",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1584166530, 1584166530941710100),
		Data:     []byte(_Assets8081c7988a2281304337e346d3d11cb790e1300e),
	}}, "")