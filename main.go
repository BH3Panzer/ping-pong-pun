package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	const width int32 = 1280
	const height int32 = 720
	const center_title int32 = width/3
	rl.InitWindow(width, height, "Ping Pong Pun")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rl.ToggleFullscreen()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Ping Pong Pun", center_title, 20, 50, rl.Gray)
		rl.DrawText("Press enter to start", center_title, 100, 20, rl.Gray)

		rl.EndDrawing()
	}
}
