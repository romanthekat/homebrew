package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io/fs"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
)

var cards []*widget.Card
var iconsPaths []string

var iconsDirPath = "./icons"

const sizeDimension = 128
const maxTitleLengthLetters = 18

// as naive implementation as possible
func main() {
	a := app.New()
	w := a.NewWindow("Random Icons")

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		iconsDirPath = argsWithoutProg[0]
	}

	cards = createCards()

	iconsPaths = readIcons()
	if len(iconsPaths) > 0 {
		rerollAllCards(iconsPaths)
	} else {
		w.SetTitle("No icons found in ./icons and no path provided as argument value - the app won't work")
	}

	grid := container.NewGridWithRows(1,
		widget.NewButton("Roll", func() {
			rerollAllCards(iconsPaths)
		}),
	)
	for _, card := range cards {
		grid.Add(card)
	}

	w.SetContent(grid)

	w.ShowAndRun()
}

func createCards() []*widget.Card {
	card1 := newCardWidget()
	card2 := newCardWidget()
	card3 := newCardWidget()
	card4 := newCardWidget()
	card5 := newCardWidget()
	card6 := newCardWidget()

	return []*widget.Card{card1, card2, card3, card4, card5, card6}
}

func rerollAllCards(icons []string) {
	for _, card := range cards {
		updateImage(card, icons)
	}
}

func newCardWidget() *widget.Card {
	card := widget.NewCard("", "", nil)
	card.Resize(fyne.NewSize(sizeDimension, sizeDimension))

	return card
}

func updateImage(card *widget.Card, icons []string) {
	imagePath := getRandom(icons)

	card.Image = loadImage(imagePath)
	card.SetSubTitle(getNameByImageFilename(imagePath))
}

func getRandom(list []string) string {
	return list[rand.IntN(len(list))]
}

func readIcons() []string {
	var result []string

	err := filepath.WalkDir(iconsDirPath,
		func(path string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() && filepath.Ext(path) == ".png" {
				fmt.Println(path)
				result = append(result, path)
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	return result
}

func loadImage(path string) *canvas.Image {
	image := canvas.NewImageFromFile(path)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(sizeDimension, sizeDimension))

	return image
}

func getNameByImageFilename(path string) string {
	name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

	if len(name) > maxTitleLengthLetters {
		name = name[0:maxTitleLengthLetters]
	}

	return name
}
