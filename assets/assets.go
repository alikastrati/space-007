package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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
var ScoreFont = MustLoadFont("font.ttf")

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

func MustLoadFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})

	if err != nil {
		panic(err)
	}

	return face
}
