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
	if v, err := g.SetView("yo", topLeftX, topLeftY, bottomRightX, bottomRightY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Wrap = true
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
	if err := g.SetCurrentView("yo"); err != nil {
		return err
	}
	return nil
}

func moveRight(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx+1, cy); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox+1, oy); err != nil {
				return err
			}
		}
	}
	return nil
}

func moveDown(g *gocui.Gui, v *gocui.View) error {
	fmt.Println("yo")
	if v != nil {
		fmt.Println("hey")
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
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
	g.Cursor = true;
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
