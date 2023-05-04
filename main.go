package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

const (
	scale    = 2
	gravity  = 9.8
	maxSpeed = 5
	deAcc    = 6
)

type DogState struct {
	Sprite       rl.Texture2D
	Speed        float32
	MaxFrames    int
	SpriteWidth  int
	SpriteHeight int
	LoopFrom     int
	Reverse      bool
}

type Dog struct {
	Position    rl.Vector2
	AnimCounter int
	State       DogState
	Direction   int
	YSpeed      float32
	XSpeed      float32
}

type Timer struct {
	StartTime float64 // Start time (seconds)
	LifeTime  float64 // Lifetime (seconds)
}

func StartTimer(timer *Timer, lifetime float64) {
	timer.StartTime = rl.GetTime()
	timer.LifeTime = lifetime
}

func TimerDone(timer Timer) bool {
	return rl.GetTime()-timer.StartTime >= timer.LifeTime
}

func GetElapsed(timer Timer) float64 {
	return rl.GetTime() - timer.StartTime
}

func scaleImage(path string) rl.Texture2D {
	image := rl.LoadImage(path)
	rl.ImageResizeNN(image, image.Width*scale, image.Height*scale)
	return rl.LoadTextureFromImage(image)
}

func main() {

	rl.InitWindow(800, 450, "framework")

	houndIdle := scaleImage("assets/hellhound/hell-hound-idle.png")

	runState := DogState{
		scaleImage("assets/hellhound/hell-hound-run.png"),
		9,
		5,
		67,
		32,
		0,
		false,
	}

	walkState := DogState{
		scaleImage("assets/hellhound/hell-hound-walk.png"),
		3,
		12,
		64,
		32,
		0,
		false,
	}

	standState := DogState{
		houndIdle,
		0,
		1,
		64,
		32,
		0,
		false,
	}

	idleState := DogState{
		houndIdle,
		0,
		6,
		64,
		32,
		2,
		false,
	}

	riseState := DogState{
		houndIdle,
		0,
		6,
		64,
		32,
		0,
		true,
	}

	fallState := DogState{
		scaleImage("assets/hellhound/hell-hound-jump.png"),
		0,
		0,
		64,
		48,
		4,
		false,
	}

	d := Dog{
		rl.NewVector2(250, 0),
		0,
		walkState,
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
		if !(d.State == idleState || d.State == riseState || d.State == fallState) {
			isStanding = false
			if rl.IsKeyDown(rl.KeyRight) {
				d.Direction = 1
			} else if rl.IsKeyDown(rl.KeyLeft) {
				d.Direction = -1
			} else {
				isStanding = true
			}

			if !isStanding {
				if d.XSpeed*float32(d.Direction) > 2 {
					d.State = runState
				} else {
					d.State = walkState
				}
			} else {
				d.State = standState
			}

			if d.YSpeed != 0 {
				d.State = fallState
			}

			if rl.IsKeyDown(rl.KeySpace) && d.XSpeed == 0 {
				StartTimer(&riseTimer, 2)
				d.State = idleState
			}
		}

		if TimerDone(riseTimer) {
			if d.State == idleState {
				d.State = riseState
			} else if d.State == riseState {
				if d.AnimCounter == 5 {
					d.State = standState
				}
			}
		}
		if d.YSpeed == 0 && lastState == fallState {
			d.State = standState
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
			if lastState == fallState {
				d.Position.Y += 16
			}
		}

		// draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("currentFrame: "+strconv.Itoa(d.AnimCounter), 190, 200, 20, rl.LightGray)

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
