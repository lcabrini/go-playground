package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Width  int32  = 1024
	Height int32  = 768
	Title  string = "Image"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Printf("No image specified")
		os.Exit(1)
	}

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(Width, Height, Title)
	rl.SetTargetFPS(60)

	texture := rl.LoadTexture(os.Args[1])

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTexture(texture, 0, 0, rl.White)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
