package lib

import (
	"bytes"
	"image"
	"image/color"
	"image/png"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
)

var Colors = map[string]color.NRGBA{
	"red":         {0xf4, 0x43, 0x36, 0xff},
	"pink":        {0xe9, 0x1e, 0x63, 0xff},
	"purple":      {0x9c, 0x27, 0xb0, 0xff},
	"indigo":      {0x3f, 0x51, 0xb5, 0xff},
	"blue":        {0x21, 0x96, 0xf3, 0xff},
	"light_blue":  {0x03, 0xa9, 0xf4, 0xff},
	"cyan":        {0x00, 0xbc, 0xd4, 0xff},
	"teal":        {0x00, 0x96, 0x88, 0xff},
	"green":       {0x4c, 0xaf, 0x50, 0xff},
	"light_green": {0x8b, 0xc3, 0x4a, 0xff},
	"lime":        {0xcd, 0xdc, 0x39, 0xff},
	"yellow":      {0xff, 0xeb, 0x3b, 0xff},
	"amber":       {0xff, 0xc1, 0x07, 0xff},
	"orange":      {0xff, 0x98, 0x00, 0xff},
	"brown":       {0x79, 0x55, 0x48, 0xff},
}

func Plot(c chart.Chart, name string) telegram.FileReader {
	img := image.NewRGBA(image.Rect(0, 0, 1280, 720))
	igr := imgg.AddTo(img, 0, 0, 1280, 720, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	c.Plot(igr)

	writer := bytes.NewBuffer(nil)
	err := png.Encode(writer, img)
	if err != nil {
		panic(err)
	}

	reader := telegram.FileReader{
		Name:   name + ".png",
		Reader: writer,
		Size:   -1,
	}
	return reader
}
