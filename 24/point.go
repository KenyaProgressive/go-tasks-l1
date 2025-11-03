package main

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point { // Конструктор структуры
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 { // Метод для расчёта раасстояния между точками по формуле
	return math.Sqrt(math.Pow((other.x-p.x), 2) + math.Pow((other.y-p.y), 2))
}
