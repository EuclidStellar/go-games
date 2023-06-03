package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Point struct {
	x float64
	y float64
}

func (point Point) toSdlPoint() *sdl.Point {
	return &sdl.Point{
		X: int32(point.x),
		Y: int32(point.y)}
}

type Size struct {
	w float64
	h float64
}

type Rect struct {
	p Point
	s Size
}

func (rect Rect) setCenter(center Point) {
	rect.p.x = center.x - rect.s.w/2.0
	rect.p.y = center.y - rect.s.h/2.0
}

func (rect Rect) center() Point {
	return Point{
		x: rect.p.x - rect.s.w/2.0,
		y: rect.p.y - rect.s.h/2.0}
}

func (rect Rect) toSdlRect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(rect.p.x),
		Y: int32(rect.p.y),
		W: int32(rect.s.w),
		H: int32(rect.s.h)}
}
