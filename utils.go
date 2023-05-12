package main

import rl "github.com/gen2brain/raylib-go/raylib"

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

func ScaleImage(path string) rl.Texture2D {
	image := rl.LoadImage(path)
	rl.ImageResizeNN(image, image.Width*scale, image.Height*scale)
	return rl.LoadTextureFromImage(image)
}

