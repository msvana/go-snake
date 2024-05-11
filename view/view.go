package view

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/msvana/go-snake/model"
)

type ViewConfig struct {
	TileWidth   int
	TileHeight  int
	PanelHeight int
}

type View struct {
	config      *ViewConfig
	win         *gtk.Window
	drawingArea *gtk.DrawingArea
	game        *model.Game
}

func NewView(viewConfig *ViewConfig, game *model.Game) *View {
	windowWidth := game.FieldX * viewConfig.TileWidth
	windowHeight := game.FieldY*viewConfig.TileHeight + viewConfig.PanelHeight

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	drawingArea, err := gtk.DrawingAreaNew()

	if err != nil {
		log.Fatal("Unable to create drawing area:", err)
	}

	win.SetTitle("Snake")
	win.SetDefaultSize(windowWidth, windowHeight)
	win.Add(drawingArea)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	view := View{
		config:      viewConfig,
		win:         win,
		drawingArea: drawingArea,
		game:        game,
	}

	drawingArea.Connect("draw", view.draw)

	return &view
}

func (v *View) OnKeyPress(callback func(keyEvent *gdk.EventKey)) {
	v.win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		callback(&gdk.EventKey{Event: ev})
	})
}

func (v *View) ShowAll() {
	v.win.ShowAll()
}

func (v *View) QueueDraw() {
	v.drawingArea.QueueDraw()
}

func (v *View) draw(da *gtk.DrawingArea, cr *cairo.Context) {
	if v.game.GameOver {
		v.drawGameOver(cr)
	} else {
		v.drawBackground(cr)
		v.drawFood(cr)
		v.drawSnake(cr)
		v.drawScore(cr)
	}
}

func (v *View) drawGameOver(cr *cairo.Context) {
	cr.SetSourceRGB(1.0, 0.0, 0.0)
	cr.Rectangle(0, 0, float64(v.drawingArea.GetAllocatedWidth()), float64(v.drawingArea.GetAllocatedHeight()-v.config.PanelHeight))
	cr.FillPreserve()
	cr.Stroke()
	cr.SetSourceRGB(1.0, 1.0, 1.0)
	cr.SelectFontFace("Arial", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	cr.SetFontSize(50)
	cr.MoveTo(100, 100)
	cr.ShowText("Game Over")
}

func (v *View) drawBackground(cr *cairo.Context) {
	cr.SetSourceRGB(0.0, 0.0, 1.0)
	cr.Rectangle(0, 0, float64(v.drawingArea.GetAllocatedWidth()), float64(v.drawingArea.GetAllocatedHeight()-v.config.PanelHeight))
	cr.FillPreserve()
	cr.Stroke()
}

func (v *View) drawFood(cr *cairo.Context) {
	cr.SetSourceRGB(1.0, 0.0, 0.0)
	x := v.game.Food.X * v.config.TileWidth
	y := v.game.Food.Y * v.config.TileHeight
	cr.Rectangle(float64(x), float64(y), float64(v.config.TileWidth), float64(v.config.TileHeight))
	cr.FillPreserve()
	cr.Stroke()
}

func (v *View) drawSnake(cr *cairo.Context) {
	cr.SetSourceRGB(0.0, 1.0, 0.0)
	for _, bp := range v.game.Snake.Body.Nodes {
		x := bp.Value.X * v.config.TileWidth
		y := bp.Value.Y * v.config.TileHeight
		cr.Rectangle(float64(x), float64(y), float64(v.config.TileWidth), float64(v.config.TileHeight))
		cr.FillPreserve()
		cr.Stroke()
	}
}

func (v *View) drawScore(cr *cairo.Context) {
	cr.SetSourceRGB(0.0, 0.0, 0.0)
	cr.SelectFontFace("Arial", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	cr.SetFontSize(20)
	cr.MoveTo(10, float64(v.drawingArea.GetAllocatedHeight()-v.config.PanelHeight+32))
	cr.ShowText(fmt.Sprintf("Score: %d", v.game.Score))
}
