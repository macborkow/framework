package main

import rl "github.com/gen2brain/raylib-go/raylib"

type DogState struct {
	Sprite       rl.Texture2D
	Speed        float32
	MaxFrames    int
	SpriteWidth  int
	SpriteHeight int
	LoopFrom     int
	Reverse      bool
}

func GetStates() map[string]DogState {

	houndIdle := ScaleImage("assets/hellhound/hell-hound-idle.png")

	runState := DogState{
		ScaleImage("assets/hellhound/hell-hound-run.png"),
		9,
		5,
		67,
		32,
		0,
		false,
	}

	walkState := DogState{
		ScaleImage("assets/hellhound/hell-hound-walk.png"),
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
		ScaleImage("assets/hellhound/hell-hound-jump.png"),
		0,
		0,
		64,
		48,
		4,
		false,
	}

	dogState := make(map[string]DogState)
	dogState["fallState"] = fallState
	dogState["riseState"] = riseState
	dogState["idleState"] = idleState
	dogState["standState"] = standState
	dogState["walkState"] = walkState
	dogState["runState"] = runState

	return dogState
}
