package main

import (
	"2048project/game"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	tileSize   = 100
	tilePad    = 10
	boardX     = 20
	boardY     = 80
	screenW    = boardX*2 + tileSize*game.Size + tilePad*(game.Size+1)
	screenH    = boardY + 20 + tileSize*game.Size + tilePad*(game.Size+1)
)

var tileColors = map[int]color.RGBA{
	0:    {205, 193, 180, 255},
	2:    {238, 228, 218, 255},
	4:    {237, 224, 200, 255},
	8:    {242, 177, 121, 255},
	16:   {245, 149, 99, 255},
	32:   {246, 124, 95, 255},
	64:   {246, 94, 59, 255},
	128:  {237, 207, 114, 255},
	256:  {237, 204, 97, 255},
	512:  {237, 200, 80, 255},
	1024: {237, 197, 63, 255},
	2048: {237, 194, 46, 255},
}

type App struct {
	g *game.Game
}

func (a *App) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		a.g.Reset()
		return nil
	}

	var dir game.Direction
	moved := false
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW):
		dir = game.Up
		moved = true
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS):
		dir = game.Down
		moved = true
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA):
		dir = game.Left
		moved = true
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD):
		dir = game.Right
		moved = true
	}

	if moved {
		a.g.Move(dir)
	}
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{250, 248, 239, 255})

	// 标题和分数
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("2048   Score: %d", a.g.Score), boardX, 20)

	if a.g.Won {
		ebitenutil.DebugPrintAt(screen, "You Win! Press R to restart", boardX, 45)
	} else if a.g.Over {
		ebitenutil.DebugPrintAt(screen, "Game Over! Press R to restart", boardX, 45)
	} else {
		ebitenutil.DebugPrintAt(screen, "Arrow keys / WASD to move, R to restart", boardX, 45)
	}

	// 棋盘背景
	bgImg := ebiten.NewImage(
		tileSize*game.Size+tilePad*(game.Size+1),
		tileSize*game.Size+tilePad*(game.Size+1),
	)
	bgImg.Fill(color.RGBA{187, 173, 160, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(boardX), float64(boardY))
	screen.DrawImage(bgImg, op)

	// 画每个格子
	for r := 0; r < game.Size; r++ {
		for c := 0; c < game.Size; c++ {
			val := a.g.Board[r][c]
			x := boardX + tilePad + c*(tileSize+tilePad)
			y := boardY + tilePad + r*(tileSize+tilePad)

			tileImg := ebiten.NewImage(tileSize, tileSize)
			clr, ok := tileColors[val]
			if !ok {
				clr = color.RGBA{60, 58, 50, 255} // 超过2048的颜色
			}
			tileImg.Fill(clr)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(tileImg, op)

			if val != 0 {
				// 数字居中显示
				txt := fmt.Sprintf("%d", val)
				tx := x + tileSize/2 - len(txt)*3
				ty := y + tileSize/2 - 4
				ebitenutil.DebugPrintAt(screen, txt, tx, ty)
			}
		}
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

func main() {
	app := &App{g: game.NewGame()}
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("2048")
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
