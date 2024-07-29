package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

//go:embed *
var assets embed.FS

var PlayerSprite = MustLoadImage("playerShip1_blue.png")
var meteorImageFiles = []string{
	"meteors/meteorBrown_big1.png",

	"meteors/meteorGrey_med1.png",
	"meteors/meteorGrey_small1.png",
}
var ShootingSprite = MustLoadImage("laserGreen04.png")

var MeteorSprites []*ebiten.Image

// Loads the images one by one (probably not a good idea)
// TO-DO : Look for a better solution to importing assets
func init() {
	for _, path := range meteorImageFiles {
		MeteorSprites = append(MeteorSprites, MustLoadImage(path))
	}
}

func MustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
