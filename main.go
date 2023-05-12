package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	scale    = 2
	gravity  = 9.8
	maxSpeed = 5
	deAcc    = 6
)

type Dog struct {
	Position    rl.Vector2
	AnimCounter int
	State       DogState
	Direction   int
	YSpeed      float32
	XSpeed      float32
}

func main() {

	rl.InitWindow(800, 450, "framework")

	dogState := GetStates()

	d := Dog{
		rl.NewVector2(250, 0),
		0,
		dogState["standState"],
		-1,
		0,
		0,
	}


	var riseTimer Timer

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	framesCounter := 0

	for !rl.WindowShouldClose() {
		// animation && controls
		framesCounter++
		if framesCounter >= 6 {
			framesCounter = 0
			d.AnimCounter++
		}

		if d.AnimCounter >= d.State.MaxFrames {
			d.AnimCounter = d.State.LoopFrom
		}

		lastState := d.State

		isStanding := true
		if !(d.State == dogState["idleState"] || d.State == dogState["riseState"] || d.State == dogState["fallState"]) {
			isStanding = false
			if rl.IsKeyDown(rl.KeyRight) {
				d.Direction = 1
			} else if rl.IsKeyDown(rl.KeyLeft) {
				d.Direction = -1
			} else {
				isStanding = true
			}

			if !isStanding {
				if d.XSpeed*float32(d.Direction) > 3 {
					d.State = dogState["runState"]
				} else {
					d.State = dogState["walkState"]
				}
			} else {
				d.State = dogState["standState"]
			}

			if d.YSpeed != 0 {
				d.State = dogState["fallState"]
			}

			if rl.IsKeyDown(rl.KeySpace) && d.XSpeed == 0 {
				StartTimer(&riseTimer, 2)
				d.State = dogState["idleState"]
			}
		}

		if TimerDone(riseTimer) {
			if d.State == dogState["idleState"] {
				d.State = dogState["riseState"]
			} else if d.State == dogState["riseState"] {
				if d.AnimCounter == 5 {
					d.State = dogState["standState"]
				}
			}
		}
		if d.YSpeed == 0 && lastState == dogState["fallState"] {
			d.State = dogState["standState"]
		}

		if lastState != d.State {
			d.AnimCounter = 0
		}

		// movement & logic

		if !isStanding {
			d.XSpeed += float32(d.Direction) * d.State.Speed * scale * 0.01
			if d.XSpeed*float32(d.Direction) > maxSpeed {
				d.XSpeed = maxSpeed * float32(d.Direction)
			}
		} else if d.XSpeed != 0 {
			direction := 1
			if d.XSpeed < 0 {
				direction = -1
			}
			d.XSpeed -= float32(direction) * deAcc * scale * 0.01
			if d.XSpeed*float32(direction) < 0.1 {
				d.XSpeed = 0
			}
		}
		d.Position.X += d.XSpeed

		if d.Position.Y+float32(d.State.SpriteHeight) < 340 {
			d.YSpeed += gravity * 0.01
			d.Position.Y += d.YSpeed
		} else {
			d.YSpeed = 0
			if lastState == dogState["fallState"] {
				d.Position.Y += 16
			}
		}

		// draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		// rl.DrawText("currentFrame: "+strconv.Itoa(d.AnimCounter), 190, 200, 20, rl.LightGray)

		currentFrame := d.AnimCounter

		if d.State.Reverse {
			currentFrame = d.State.MaxFrames - d.AnimCounter - 1
		}

		if d.Direction != 1 {
			rl.DrawTextureRec(d.State.Sprite,
				rl.NewRectangle(float32(d.State.SpriteWidth*currentFrame)*scale, 0, float32(d.State.SpriteWidth)*scale, float32(d.State.SpriteHeight)*scale),
				d.Position, rl.RayWhite)
		} else {
			rl.DrawTextureRec(d.State.Sprite,
				rl.NewRectangle(float32(d.State.SpriteWidth*(currentFrame))*scale, 0, float32(d.State.SpriteWidth)*-1*scale, float32(d.State.SpriteHeight)*scale),
				d.Position, rl.RayWhite)
		}

		rl.EndDrawing()

	}
}
