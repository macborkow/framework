package main

import (
  rl "github.com/gen2brain/raylib-go/raylib"
  "strconv"
)

const (
	spriteWidth = 64
	spriteHeight = 32
)

func main() {
	rl.InitWindow(800, 450, "framework")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	position := rl.NewVector2(250, 250)
	frame := rl.NewRectangle(0, 0, spriteWidth, spriteHeight)
  sprite := rl.LoadTexture("assets/hellhound/hell-hound-walk.png")
  framesCounter := 0
  animCounter := 0

	for !rl.WindowShouldClose() {
    framesCounter++
    frame.X = float32(spriteWidth*animCounter)
    if framesCounter >= 6 {
      framesCounter = 0
      animCounter++
    }
    if animCounter >= 12 {
      animCounter = 0
    }

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("currentFrame: " + strconv.Itoa(animCounter), 190, 200, 20, rl.LightGray)
		rl.DrawTextureRec(sprite, frame, position, rl.RayWhite)

		rl.EndDrawing()

	}
}
