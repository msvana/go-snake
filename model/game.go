package model

type GameConfig struct {
	FieldX int
	FieldY int
	StartX int
	StartY int
}

type Game struct {
	FieldX   int
	FieldY   int
	Score    int
	GameOver bool
	Snake    *Snake
	Food     Point
	Paused   bool
}

func NewGame(config GameConfig) *Game {
	game := Game{
		FieldX:   config.FieldX,
		FieldY:   config.FieldY,
		Score:    0,
		GameOver: false,
		Snake:    NewSnake(config.StartX, config.StartY),
	}

	game.Food = GenerateFood(game.Snake, game.FieldX, game.FieldY)
	return &game
}

func (g *Game) Tick() {
	if g.Paused {
		return
	}

	g.Snake.Move()

	if g.Snake.CollidesWithFood(g.Food) {
		g.Snake.Grow()
		g.Food = GenerateFood(g.Snake, g.FieldX, g.FieldY)
		g.Score++
	}

	if g.Snake.CollidesWithSelf() || g.Snake.CollidesWithWall(g.FieldX, g.FieldY) {
		g.GameOver = true
	}

}
