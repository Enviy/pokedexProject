package main

import (
	"os"
	"fmt"
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	_ "embed"
	"log"

	"Enviy/pokedexProject/util"
	"Enviy/pokedexProject/pokemon"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var img *ebiten.Image
var factsFont font.Face
var artFont font.Face

//go:embed images/pokedex_open.png
var pokedexImage []byte
//go:embed fonts/SFNSMono.ttf
var loadedFonts []byte

func init() {
	var err error
	imgDecoded, _, err := image.Decode(bytes.NewReader(pokedexImage))
	if err != nil {
		log.Fatal(err)
	}
	img = ebiten.NewImageFromImage(imgDecoded)

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	factsFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 12,
		DPI: dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	art_tt, err := opentype.Parse(loadedFonts)
	if err != nil {
		log.Fatal(err)
	}
	artFont, err = opentype.NewFace(art_tt, &opentype.FaceOptions{
		Size: 8,
		DPI: dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	newPokemon bool
	voice      string
	facts      string
	art        string
	text       string
	title      string
	runes      []rune
	counter    int
}

func (g *Game) Update() error {
	// Collect user input
	g.runes = ebiten.AppendInputChars(g.runes[:0])
	g.text += string(g.runes)
	if util.RepeatKey(ebiten.KeyEnter) || util.RepeatKey(ebiten.KeyNumpadEnter) {
		if g.text == "" {
			return nil
		}
		if g.text == "yes" {
			g.text = ""
			g.facts = ""
			g.art = ""
			g.newPokemon = false
		} else if g.text == "no" {
			os.Exit(1)
		} else {
			// Attempt to find pokemon info.
			content, found, err := pokemon.GetBase(g.text)
			if err != nil {
				log.Fatal(err)
				return err
			}
			if !found {
				g.facts = "Check spelling or internet.\nRetry\n"
			} else if found {
				art, err := pokemon.GetArt(content)
				if err != nil {
					log.Fatal(err)
					g.art = fmt.Sprintf("%v", err)
				}
				g.art = art
				g.facts = pokemon.GetInfo(content)
			}
			g.text = ""
			g.newPokemon = true
			g.voice = "initial"
		}
	}
	if g.voice == "drawn" {
		g.voice = "off"
		util.OakLine(g.facts)
	}
	// If the backspace key is pressed, remove one character.
	if util.RepeatKey(ebiten.KeyBackspace) {
		if len(g.text) >= 1 {
			g.text = g.text[:len(g.text)-1]
		}
	}
	g.counter ++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// x, y are the text box coordinates.
	const (
		x_facts = 634
		y_facts = 280
		x_art = 120
		y_art = 270
	)
	screen.DrawImage(img, nil)
	if !g.newPokemon {
		prompt := "Type a Pokemon's name:\n" + g.text
		text.Draw(screen,
			prompt,
			factsFont,
			x_facts,
			y_facts,
			color.White)
		text.Draw(screen,
			g.title,
			artFont,
			x_art + 30,
			y_art,
			color.White)
	}
	if g.newPokemon {
		// draw ascii art
		text.Draw(screen,
			g.art,
			artFont,
			x_art,
			y_art,
			color.White)
		// draw collected facts
		wrappedFacts := util.WordWrap(g.facts, 50)
		text.Draw(screen,
			wrappedFacts + g.text,
			factsFont,
			x_facts,
			y_facts,
			color.White)
		if g.voice == "initial" {
			g.voice = "drawn"
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1040, 880)
	ebiten.SetWindowTitle("Pokédex")
	ebiten.SetWindowResizable(true)
	game := &Game{
		title: util.Banner(),
		counter: 0,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
