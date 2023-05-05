package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	paddleWidth  = 10
	paddleHeight = 100
	paddleSpeed  = 4
	ballSize     = 15
	ballSpeed    = 5
)

type Game struct {
	paddle1 *Paddle
	paddle2 *Paddle
	ball    *Ball
	score   *Score
}

type Paddle struct {
	x, y    float64
	speed   int
	upKey   ebiten.Key
	downKey ebiten.Key
}

type Ball struct {
	x, y   float64
	xSpeed float64
	ySpeed float64
}

type Score struct {
	player1 int
	player2 int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := &Game{
		paddle1: &Paddle{
			x:       30,
			y:       (screenHeight - paddleHeight) / 2,
			upKey:   ebiten.KeyW,
			downKey: ebiten.KeyS,
		},
		paddle2: &Paddle{
			x:       screenWidth - 30 - paddleWidth,
			y:       (screenHeight - paddleHeight) / 2,
			upKey:   ebiten.KeyUp,
			downKey: ebiten.KeyDown,
		},
		ball: &Ball{
			x:      (screenWidth - ballSize) / 2,
			y:      (screenHeight - ballSize) / 2,
			xSpeed: ballSpeed,
			ySpeed: ballSpeed,
		},
		score: &Score{
			player1: 0,
			player2: 0,
		},
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	g.paddle1.Update()
	g.paddle2.Update()
	g.ball.Update()
	g.ball.CollidesWith(g.paddle1, g.paddle2)
	g.CheckScore()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.paddle1.Draw(screen)
	g.paddle2.Draw(screen)
	g.ball.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (p *Paddle) Update() {
	if ebiten.IsKeyPressed(p.upKey) {
		p.y -= paddleSpeed
	}
	if ebiten.IsKeyPressed(p.downKey) {
		p.y += paddleSpeed
	}

	if p.y < 0 {
		p.y = 0
	}
	if p.y+paddleHeight > screenHeight {
		p.y = screenHeight - paddleHeight
	}
}

func (b *Ball) Update() {
	b.x += b.xSpeed
	b.y += b.ySpeed

	if b.y < 0 || b.y+ballSize > screenHeight {
		b.ySpeed = -b.ySpeed
	}
}

func (p *Paddle) CollidesWith(b *Ball) bool {
	return b.x+ballSize > p.x && b.x < p.x+paddleWidth &&
		b.y+ballSize > p.y && b.y < p.y+paddleHeight
}

func (b *Ball) CollidesWith(p1, p2 *Paddle) {
	if p1.CollidesWith(b) || p2.CollidesWith(b) {
		b.xSpeed = -b.xSpeed
	}
}

func (g *Game) CheckScore() {
	if g.ball.x < 0 {
		g.score.player2++
		g.ball.Reset()
	} else if g.ball.x+ballSize > screenWidth {
		g.score.player1++
		g.ball.Reset()
	}
}

func (b *Ball) Reset() {
	b.x = (screenWidth - ballSize) / 2
	b.y = (screenHeight - ballSize) / 2
	b.xSpeed = float64(rand.Intn(2*ballSpeed) - ballSpeed)
	b.ySpeed = float64(rand.Intn(2*ballSpeed) - ballSpeed)
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	paddle := ebiten.NewImage(paddleWidth, paddleHeight)
	paddle.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(paddle, op)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(b.x), float32(b.y), float32(ballSize)/2, color.RGBA{R: 255, G: 100, B: 100, A: 1}, true)
}
