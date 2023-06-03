package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func CreateFoe(renderer *sdl.Renderer, point Point) *Element {
	foe := &Element{}

	foe.frame = Rect{p: point, s: Size{w: 16.0, h: 16.0}}

	sr := CreateSpriteRenderer(
		foe,
		renderer,
		"sprites/enemy-small.png",
		Rect{p: Point{x: 0.0, y: 0.0}, s: Size{w: 16.0, h: 16.0}})
	foe.addComponent(sr)

	foe.active = true

	return foe
}
