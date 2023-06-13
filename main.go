package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_WIDTH  = 600
	SCREEN_HEIGHT = 800
)



var elements []*Element
var bulletPool []*Element

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("failed in initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Galaxiga",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH, SCREEN_HEIGHT,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("failed in initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("failed in initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	elements = append(elements, CreatePlayer(renderer))

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			point := Point{
				x: float64(i/5)*SCREEN_WIDTH + 8.0,
				y: float64(j)*16.0 + 8.0,
			}

			elements = append(elements, CreateFoe(renderer, point))
		}
	}

	InitBulletPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, el := range elements {
			if el.active {
				err = el.update()
				if err != nil {
					fmt.Println("failed in updating element:", err)
					return
				}
				err = el.draw(renderer)
				if err != nil {
					fmt.Println("failed in drawing element:", err)
				}
			}
		}

		renderer.Present()
	}
}

func InitBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := CreateBullet(renderer)
		bulletPool = append(bulletPool, bul)
		elements = append(elements, bul)
	}
}

func BulletFromPool() (*Element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}
