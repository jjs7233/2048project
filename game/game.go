package game

import (
	"math/rand"
)

const Size = 4

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Game struct {
	Board [Size][Size]int
	Score int
	Over  bool
	Won   bool
}

func NewGame() *Game {
	g := &Game{}
	g.addTile()
	g.addTile()
	return g
}

func (g *Game) Move(dir Direction) bool {
	if g.Over {
		return false
	}

	old := g.Board
	switch dir {
	case Left:
		for r := 0; r < Size; r++ {
			g.Board[r] = mergeLine(g.Board[r], &g.Score)
		}
	case Right:
		for r := 0; r < Size; r++ {
			g.Board[r] = mergeLineReverse(g.Board[r], &g.Score)
		}
	case Up:
		for c := 0; c < Size; c++ {
			col := [Size]int{g.Board[0][c], g.Board[1][c], g.Board[2][c], g.Board[3][c]}
			col = mergeLine(col, &g.Score)
			for r := 0; r < Size; r++ {
				g.Board[r][c] = col[r]
			}
		}
	case Down:
		for c := 0; c < Size; c++ {
			col := [Size]int{g.Board[0][c], g.Board[1][c], g.Board[2][c], g.Board[3][c]}
			col = mergeLineReverse(col, &g.Score)
			for r := 0; r < Size; r++ {
				g.Board[r][c] = col[r]
			}
		}
	}

	if old != g.Board {
		g.addTile()
		g.checkWin()
		if !g.canMove() {
			g.Over = true
		}
		return true
	}
	return false
}

// mergeLine 向左合并一行
func mergeLine(line [Size]int, score *int) [Size]int {
	// 1. 去零压缩
	var compact [Size]int
	idx := 0
	for _, v := range line {
		if v != 0 {
			compact[idx] = v
			idx++
		}
	}
	// 2. 合并相邻相同
	for i := 0; i < Size-1; i++ {
		if compact[i] != 0 && compact[i] == compact[i+1] {
			compact[i] *= 2
			*score += compact[i]
			compact[i+1] = 0
		}
	}
	// 3. 再次压缩
	var result [Size]int
	idx = 0
	for _, v := range compact {
		if v != 0 {
			result[idx] = v
			idx++
		}
	}
	return result
}

// mergeLineReverse 向右合并（反转后向左合并再反转）
func mergeLineReverse(line [Size]int, score *int) [Size]int {
	reversed := [Size]int{line[3], line[2], line[1], line[0]}
	merged := mergeLine(reversed, score)
	return [Size]int{merged[3], merged[2], merged[1], merged[0]}
}

func (g *Game) checkWin() {
	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if g.Board[r][c] == 2048 {
				g.Won = true
			}
		}
	}
}

func (g *Game) canMove() bool {
	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if g.Board[r][c] == 0 {
				return true
			}
			if c < Size-1 && g.Board[r][c] == g.Board[r][c+1] {
				return true
			}
			if r < Size-1 && g.Board[r][c] == g.Board[r+1][c] {
				return true
			}
		}
	}
	return false
}

func (g *Game) addTile() {
	var empty [][2]int
	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if g.Board[r][c] == 0 {
				empty = append(empty, [2]int{r, c})
			}
		}
	}
	if len(empty) == 0 {
		return
	}
	pos := empty[rand.Intn(len(empty))]
	if rand.Intn(10) < 9 {
		g.Board[pos[0]][pos[1]] = 2
	} else {
		g.Board[pos[0]][pos[1]] = 4
	}
}

func (g *Game) Reset() {
	g.Board = [Size][Size]int{}
	g.Score = 0
	g.Over = false
	g.Won = false
	g.addTile()
	g.addTile()
}
