package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const SHOT_COOLDOWN = time.Millisecond * 250

func CreatePlayer(renderer *sdl.Renderer) *Element {
	player := &Element{}

	player.frame = Rect{
		p: Point{x: SCREEN_WIDTH / 2.0, y: SCREEN_HEIGHT / 2.0},
		s: Size{w: 16.0, h: 16.0}}

	sr := CreateSpriteRenderer(
		player,
		renderer,
		"sprites/ship.png",
		Rect{p: Point{x: 32.0, y: 0.0}, s: Size{w: 16.0, h: 16.0}})
	player.addComponent(sr)

	mover := CreateKeyboardMover(player, 0.05)
	player.addComponent(mover)

	shooter := CreateKeyboardShooter(player, SHOT_COOLDOWN)
	player.addComponent(shooter)

	player.active = true

	return player
}
