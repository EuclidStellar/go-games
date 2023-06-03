package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardShooter struct {
	container *Element
	cooldown  time.Duration
	lastShot  time.Time
}

func CreateKeyboardShooter(container *Element, cooldown time.Duration) *KeyboardShooter {
	return &KeyboardShooter{
		container: container,
		cooldown:  cooldown}
}

func (shooter *KeyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (shooter *KeyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := shooter.container.frame.p

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(shooter.lastShot) >= shooter.cooldown {
			shooter.shoot(pos.x+6, pos.y-2)
			shooter.shoot(pos.x-6, pos.y-2)

			shooter.lastShot = time.Now()
		}
	}

	return nil
}

func (shooter *KeyboardShooter) onCollide(other *Element) error {
	return nil
}

func (shooter *KeyboardShooter) shoot(x, y float64) {
	if bul, ok := BulletFromPool(); ok {
		bul.active = true
		bul.frame.p.x = x
		bul.frame.p.y = y
		bul.rotation = 270 * (math.Pi / 180)
	}
}
