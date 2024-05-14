package main

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/msvana/go-snake/model"
	"github.com/msvana/go-snake/view"
)

func main() {
	gtk.Init(nil)

	gameConfig := model.GameConfig{
		FieldX: 25,
		FieldY: 21,
		StartX: 12,
		StartY: 10,
	}

	viewConfig := view.ViewConfig{
		TileWidth:   25,
		TileHeight:  25,
		PanelHeight: 50,
	}

	game := model.NewGame(gameConfig)
	view := view.NewView(&viewConfig, game)

	view.OnKeyPress(func(keyEvent *gdk.EventKey) {
		switch keyEvent.KeyVal() {
		case gdk.KEY_Up:
			game.Snake.Dir = model.Up
		case gdk.KEY_Down:
			game.Snake.Dir = model.Down
		case gdk.KEY_Left:
			game.Snake.Dir = model.Left
		case gdk.KEY_Right:
			game.Snake.Dir = model.Right
		case gdk.KEY_space:
			game.Paused = !game.Paused
		}
	})

	view.ShowAll()

	glib.TimeoutAdd(500, func() bool {
		game.Tick()
		view.QueueDraw()
		return true
	})

	gtk.Main()
}
