package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/*
	Adapted from https://happycoding.io/tutorials/processing/arraylists/random-walkers
*/

const (
	Width  int32  = 1024
	Height int32  = 768
	Title  string = "Random Walkers"
)

type Walker struct {
	Pos   rl.Vector2
	Color rl.Color
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(Width, Height, Title)
	rl.SetTargetFPS(120)

	image := rl.LoadImageFromScreen()
	walkers := []Walker{}

	for !rl.WindowShouldClose() {
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			mpos := rl.GetMousePosition()
			r := uint8(rl.GetRandomValue(0, 255))
			g := uint8(rl.GetRandomValue(0, 255))
			b := uint8(rl.GetRandomValue(0, 255))
			walker := Walker{Pos: mpos, Color: rl.Color{R: r, G: g, B: b, A: 255}}
			walkers = append(walkers, walker)
			fmt.Println(len(walkers))
		}

		for i, walker := range walkers {
			walkers[i].Pos.X += float32(rl.GetRandomValue(-1, 1))
			walkers[i].Pos.Y += float32(rl.GetRandomValue(-1, 1))
			rl.ImageDrawPixelV(image, walker.Pos, walker.Color)
		}

		rl.BeginDrawing()
		texture := rl.LoadTextureFromImage(image)
		rl.DrawTexture(texture, 0, 0, rl.White)
		rl.EndDrawing()
		rl.UnloadTexture(texture)
	}

	rl.UnloadImage(image)
	rl.CloseWindow()
}
