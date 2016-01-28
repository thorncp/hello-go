package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)
func layout(g *gocui.Gui) error {
	theSizeX, theSizeY := 12, 6
	maxX, maxY := g.Size()
	centerX, centerY := maxX/2, maxY/2
	topLeftX, topLeftY := centerX - theSizeX, centerY - theSizeY
	bottomRightX, bottomRightY := centerX + theSizeX, centerY + theSizeY
	if v, err := g.SetView("greeting", topLeftX, topLeftY, bottomRightX, bottomRightY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, " 1 2 3 | 4 5 6 | 7 8 9")
		fmt.Fprintln(v, " 1 2 3 | 4 5 6 | 7 8 9")
		fmt.Fprintln(v, " 1 2 3 | 4 5 6 | 7 8 9")
		fmt.Fprintln(v, " ---------------------")
		fmt.Fprintln(v, " 1 2 3 | 4 5 6 | 7 8 9")
		fmt.Fprintln(v, " 1 2 3 | 4 5 6 | 7 8 9")
		fmt.Fprintln(v, " 1 2 3 | 4 5 6 | 7 8 9")
		fmt.Fprintln(v, " ---------------------")
		fmt.Fprintln(v, " 1 2 3 | 4 5 6 | 7 8 9")
		fmt.Fprintln(v, " 1 2 3 | 4 5 6 | 7 8 9")
		fmt.Fprintln(v, " 1 2 3 | 4 5 6 | 7 8 9")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	var err error
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.SetLayout(layout)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
