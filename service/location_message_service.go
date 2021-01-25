package service

import (
	"errors"
	"fmt"
	"math"
	shape "github.com/MindorksOpenSource/gogeom"
)
import (
	"quasar-backend/constants"
	"quasar-backend/model"
)

func GetLocation(distances []float64) (x float64, y float64, err error) {
	centerCandidate := make(map[model.Point]int8)
	index := 0

	previousCircle := &model.Circle{}

	for _, position := range constants.Satellites {
		if len(distances) != 3 {
			return 0,0, errors.New("the message origin position could not be retrieved")
		}

		if index == 0 {
			previousCircle.X = position.X
			previousCircle.Y = position.Y
			previousCircle.R = distances[index]
		} else {
			newCircle := &model.Circle{}
			newCircle.X = position.X
			newCircle.Y = position.Y
			newCircle.R = distances[index]
			intersectionPositions :=extractIntersectionPoints(previousCircle, newCircle)
			for _, position := range intersectionPositions {
				centerCandidate[position]++
			}
		}
		index++
	}

	for position, count := range centerCandidate {
		if count == 2 {
			return position.X, position.Y, nil
		}
	}
	return 0,0, errors.New("The message origin position could not be retrieved")
}

func extractIntersectionPoints(a *model.Circle, b *model.Circle) (p []model.Point) {
	dx, dy := b.X - a.X, b.Y - a.Y
	centerDistance := math.Sqrt(dx * dx + dy * dy)
	radiusSum := a.R + b.R //radius and

	if centerDistance == radiusSum {
		if b.X == a.X {
			p = append(p, model.Point{X: b.X, Y: b.Y + b.R})
			return p
		}
		if b.Y == a.Y {
			p = append(p, model.Point{X: b.X + b.R, Y: b.Y})
			return p
		}
	}

	r := shape.RadiusFormOfCircle{a.X, a.Y, a.R, b.X, b.Y,b.R}

	if r.DoesCircleIntersect() {
		x0, y0, x1, y1 := r.CalculateXY()
		fmt.Println(x0, y0, x1, y1)
		p = append(p, model.Point{X: forceRoundValue(x0), Y: forceRoundValue(y0)})
		p = append(p, model.Point{X: forceRoundValue(x1), Y: forceRoundValue(y1)})

	}

	return p
}

func forceRoundValue(number float64) (result float64) {
	return math.Trunc(math.Round(number*1000))/1000
}