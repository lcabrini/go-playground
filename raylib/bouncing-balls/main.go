package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Width  int32  = 1024
	Height int32  = 768
	Title  string = "Bouncing Balls"
)

type Ball struct {
	Pos       rl.Vector2
	Velocity  rl.Vector2
	Color     rl.Color
	Radius    float32
	MouseOver bool
}

func main() {
	var err error

	count := 1
	if len(os.Args[1:]) == 1 {
		count, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid value for initial ball count: %s\n", os.Args[1])
			os.Exit(-1)
		}
	} else if len(os.Args[1:]) > 1 {
		fmt.Fprintf(os.Stderr, "usage: %s [INITIAL_BALL_COUNT]\n", os.Args[0])
		os.Exit(-1)
	}

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(Width, Height, Title)
	rl.SetTargetFPS(60)

	bgTexture := rl.LoadTexture("../../resources/gray-bg-texture.jpg")
	var balls []Ball

	for i := 0; i < count; i += 1 {
		appendBall(&balls)
	}

	for !rl.WindowShouldClose() {
		mp := rl.GetMousePosition()

		if rl.IsKeyPressed(rl.KeyB) {
			appendBall(&balls)
		}

		if rl.IsKeyPressed(rl.KeyD) {
			if len(balls) > 0 {
				balls = balls[:len(balls)-1]
			}
		}

		if rl.IsKeyPressed(rl.KeyC) {
			balls = make([]Ball, 0)
			appendBall(&balls)
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			for i := 0; i < len(balls); i += 1 {
				if balls[i].MouseOver {
					vx := float32(rl.GetRandomValue(-5, 5))
					vy := float32(rl.GetRandomValue(-5, 5))
					balls[i].Velocity = rl.Vector2{X: vx, Y: vy}
				}
			}
		}

		for i := 0; i < len(balls); i += 1 {
			balls[i].Pos.X += balls[i].Velocity.X
			balls[i].Pos.Y += balls[i].Velocity.Y

			if balls[i].Pos.X > float32(Width-int32(balls[i].Radius)) {
				balls[i].Pos.X = float32(Width - int32(balls[i].Radius))
				balls[i].Velocity.X *= -1
			}

			if balls[i].Pos.X < balls[i].Radius {
				balls[i].Pos.X = balls[i].Radius
				balls[i].Velocity.X *= -1
			}

			if balls[i].Pos.Y > float32(Height-int32(balls[i].Radius)) {
				balls[i].Pos.Y = float32(Height - int32(balls[i].Radius))
				balls[i].Velocity.Y *= -1
			}

			if balls[i].Pos.Y < balls[i].Radius {
				balls[i].Pos.Y = balls[i].Radius
				balls[i].Velocity.Y *= -1
			}

			d := math.Sqrt(math.Pow(float64(mp.X-balls[i].Pos.X), 2) + math.Pow(float64(mp.Y-balls[i].Pos.Y), 2))
			balls[i].MouseOver = false
			if d < float64(balls[i].Radius) {
				balls[i].MouseOver = true
			}
		}

		rl.BeginDrawing()
		//rl.ClearBackground(rl.Black)
		rl.DrawTexture(bgTexture, 0, 0, rl.White)

		for _, ball := range balls {
			c := rl.White
			if ball.MouseOver {
				c = rl.Yellow
			}

			rl.DrawCircleV(ball.Pos, ball.Radius, c)
			rl.DrawCircleV(ball.Pos, ball.Radius-2, ball.Color)
		}

		rl.EndDrawing()
	}

	rl.UnloadTexture(bgTexture)
	rl.CloseWindow()
}

func appendBall(balls *[]Ball) {
	ball := Ball{}
	ball.Pos.X = float32(rl.GetRandomValue(0, Width))
	ball.Pos.Y = float32(rl.GetRandomValue(0, Height))
	ball.Velocity.X = float32(rl.GetRandomValue(-5, 5))
	ball.Velocity.Y = float32(rl.GetRandomValue(-5, 5))
	r := uint8(rl.GetRandomValue(0, 255))
	g := uint8(rl.GetRandomValue(0, 255))
	b := uint8(rl.GetRandomValue(0, 255))
	ball.Color = rl.Color{R: r, G: g, B: b, A: 255}
	ball.Radius = float32(rl.GetRandomValue(10, 70))
	ball.MouseOver = false
	*balls = append(*balls, ball)
}
