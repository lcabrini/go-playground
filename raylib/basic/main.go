package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Width  int32  = 1024
	Height int32  = 768
	Title  string = "Basic"
)

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(Width, Height, Title)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
