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

	MidX int32 = Width / 2
	MidY int32 = Height / 2
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

	var scale float32 = 1.0
	tw := float32(texture.Width)
	th := float32(texture.Height)
	if texture.Height > Height || texture.Width > Width {
		scaleX := float32(Width) / tw
		scaleY := float32(Height) / th
		scale = scaleX
		if scaleY < scaleX {
			scale = scaleY
		}

		tw *= scale
		th *= scale
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		pos := rl.Vector2{X: float32(MidX) - tw/2.0, Y: float32(MidY) - th/2.0}
		rl.DrawTextureEx(texture, pos, 0, scale, rl.White)
		rl.EndDrawing()
	}

	rl.UnloadTexture(texture)
	rl.CloseWindow()
}
