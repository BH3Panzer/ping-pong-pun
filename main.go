package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var State string = "main_menu"

// Plate class to play pong with infos like position or color
type Plate struct {
	x, y, width, height int32
	color               rl.Color
}

//create instance of plate for player one at left side of screen
var player_one = Plate{
	x:      0,
	y:      0,
	width:  20,
	height: 140,
	color:  rl.Black,
}

//create instance of plate for player two at right side of screen
var player_two = Plate{
	x:      1280 - 20,
	y:      0,
	width:  20,
	height: 140,
	color:  rl.Black,
}

// function to check all inputs
func check_inputs() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		State = "game"
	}
}

// function to move player 1 with up and down arrow keys
func move_player_one() {
	if rl.IsKeyDown(rl.KeyUp) {
		player_one.y -= 10
	}
	if rl.IsKeyDown(rl.KeyDown) {
		player_one.y += 10
	}
}

func main() {
	var width int32 = 1280
	var height int32 = 720
	var center_title int32 = int32(width / 3)
	rl.InitWindow(width, height, "Ping Pong Pun")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rl.ToggleFullscreen()

	rl.HideCursor()

	for !rl.WindowShouldClose() {
		if State == "main_menu" {
			rl.BeginDrawing()

			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("Ping Pong Pun", center_title, 20, 50, rl.Gray)
			rl.DrawText("Press enter to start", center_title, 100, 20, rl.Gray)
			fps := rl.GetFPS()
			var converted_fps = int(fps)

			rl.DrawText("fps: "+strconv.Itoa(converted_fps), 10, 10, 20, rl.Black)

			check_inputs()

			rl.EndDrawing()
		} else if State == "game" {
			rl.BeginDrawing()

			move_player_one()
			rl.ClearBackground(rl.RayWhite)
			// draw player one
			rl.DrawRectangle(player_one.x, player_one.y, player_one.width, player_one.height, player_one.color)
			// draw player two
			rl.DrawRectangle(player_two.x, player_two.y, player_two.width, player_two.height, player_two.color)

			fps := rl.GetFPS()
			var converted_fps = int(fps)

			rl.DrawText("fps: "+strconv.Itoa(converted_fps), 10, 10, 20, rl.Black)

			rl.EndDrawing()
		}
		
	}
}
