package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	Width  int32  = 1024
	Height int32  = 768
	Title  string = "Bouncing Balls"
)

type Ball struct {
	Pos      rl.Vector2
	Velocity rl.Vector2
	Color    rl.Color
	Radius   float32
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(Width, Height, Title)
	rl.SetTargetFPS(60)

	var balls []Ball

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
	balls = append(balls, ball)

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyB) {
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
			balls = append(balls, ball)
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
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for _, ball := range balls {
			rl.DrawCircleV(ball.Pos, ball.Radius, rl.White)
			rl.DrawCircleV(ball.Pos, ball.Radius-2, ball.Color)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
