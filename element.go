package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Element struct {
	rotation   float64
	active     bool
	tag        string
	origin     Rect
	frame      Rect
	collisions []CollisionArea
	components []Component
}

func (el *Element) draw(renderer *sdl.Renderer) error {
	for _, comp := range el.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (el *Element) update() error {
	for _, comp := range el.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (el *Element) collide(other *Element) error {
	for _, comp := range el.components {
		err := comp.onCollide(other)
		if err != nil {
			return err
		}
	}

	return nil
}

func (el *Element) addComponent(new Component) {
	for _, existing := range el.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	el.components = append(el.components, new)
}

func (el *Element) getComponent(withType Component) Component {
	for _, comp := range el.components {
		if reflect.TypeOf(comp) == reflect.TypeOf(withType) {
			return comp
		}
	}

	panic(fmt.Sprintf(
		"no component with type %v",
		reflect.TypeOf(withType)))
}
