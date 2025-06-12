package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - drawing")
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawText("Some text", 0, 0, 20, rl.LightGray)
		rl.DrawPixel(100, 100, rl.White)
		rl.DrawPixel(102, 102, rl.White)
		rl.DrawLine(100, 0, 100, 100, rl.White)

		rl.EndDrawing()
	}
}
