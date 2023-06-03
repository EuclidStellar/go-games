package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteRenderer struct {
	container *Element
	texture   *sdl.Texture
	frame     Rect
}

func CreateSpriteRenderer(container *Element, renderer *sdl.Renderer, filename string, frame Rect) *SpriteRenderer {
	sr := &SpriteRenderer{}

	texture, err := img.LoadTexture(renderer, filename)
	if err != nil {
		panic(fmt.Errorf("failed to load %v: %v", filename, err))
	}

	sr.texture = texture
	sr.frame = frame
	sr.container = container

	return sr
}

func (sr *SpriteRenderer) onUpdate() error {
	return nil
}

func (sr *SpriteRenderer) onDraw(renderer *sdl.Renderer) error {
	renderer.CopyEx(
		sr.texture,
		sr.frame.toSdlRect(),
		sr.container.frame.toSdlRect(),
		sr.container.rotation,
		sr.frame.center().toSdlPoint(),
		sdl.FLIP_NONE)

	return nil
}

func (sr *SpriteRenderer) onCollide(other *Element) error {
	return nil
}
