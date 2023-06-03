package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
	onCollide(other *Element) error
}
