package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Ping Pong Pun", 190, 20, 50, rl.Gray)
		rl.DrawText("Press enter to start", 190, 100, 20, rl.Gray)

		rl.EndDrawing()
	}
}
