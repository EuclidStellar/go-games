package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type BulletMover struct {
	container *Element
	speed     float64
}

func CreateBulletMover(container *Element) *BulletMover {
	return &BulletMover{container: container}
}

func (mover *BulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *BulletMover) onUpdate() error {
	c := mover.container

	c.frame.p.x += BULLET_SPEED * math.Cos(c.rotation)
	c.frame.p.y += BULLET_SPEED * math.Sin(c.rotation)

	if c.frame.p.x > SCREEN_WIDTH || c.frame.p.x < 0 || c.frame.p.y > SCREEN_HEIGHT || c.frame.p.y < 0 {
		c.active = false
	}

	return nil
}

func (mover *BulletMover) onCollide(other *Element) error {
	return nil
}
