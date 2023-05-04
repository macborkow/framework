package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

const (
	spriteHeight = 32
)

type DogState struct {
	Sprite      rl.Texture2D
	Speed       float32
	MaxFrames   int
	SpriteWidth int
}

type Dog struct {
	Position    rl.Vector2
	AnimCounter int
	State       DogState
	Direction   int
}

func main() {

	rl.InitWindow(800, 450, "framework")

	runState := DogState{
		rl.LoadTexture("assets/hellhound/hell-hound-run.png"),
		0.9,
		5,
		67,
	}

	walkState := DogState{
		rl.LoadTexture("assets/hellhound/hell-hound-walk.png"),
		0.3,
		12,
		64,
	}
	d := Dog{
		rl.NewVector2(250, 250),
		0,
		walkState,
		-1,
	}

	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	framesCounter := 0

	for !rl.WindowShouldClose() {
		framesCounter++
		d.Position.X += float32(d.Direction) * d.State.Speed
		if framesCounter >= 6 {
			framesCounter = 0
			d.AnimCounter++
		}
		if d.AnimCounter >= d.State.MaxFrames {
			d.AnimCounter = 0
		}

		if rl.IsKeyDown(rl.KeyRight) {
			d.Direction = 1
		} else if rl.IsKeyDown(rl.KeyLeft) {
			d.Direction = -1
		}
		if rl.IsKeyDown(rl.KeyLeftShift) {
			d.State = runState
		} else {
			d.State = walkState
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("currentFrame: "+strconv.Itoa(d.AnimCounter), 190, 200, 20, rl.LightGray)
		if d.Direction != 1 {
			rl.DrawTextureRec(d.State.Sprite,
				rl.NewRectangle(float32(d.State.SpriteWidth*d.AnimCounter), 0, float32(d.State.SpriteWidth), spriteHeight),
				d.Position, rl.RayWhite)
		} else {
			rl.DrawTextureRec(d.State.Sprite,
				rl.NewRectangle(float32(d.State.SpriteWidth*(d.AnimCounter+1)), 0, float32(d.State.SpriteWidth)*-1, spriteHeight),
				d.Position, rl.RayWhite)
		}

		rl.EndDrawing()

	}
}
