package simple_factory

import "github.com/design-pattern/creator/simple_factory/GeometricGraphics"

const (
	Circle	= iota
	Triangle
	Square
)

type GeometricFigure interface {
	Draw()
	Erase()
}


type UnknownFigure struct {}

func (u *UnknownFigure) Draw() {
	panic("UnSupportedShapeException!")
}

func (u *UnknownFigure) Erase() {
	panic("UnSupportedShapeException!")
}


type GeometricFactory struct {}

func (gf *GeometricFactory) CreateFigure(figureType int) GeometricFigure{
	switch figureType {
	case Circle:
		return new(GeometricGraphics.Circle)
	case Triangle:
		return new(GeometricGraphics.Triangle)
	case Square:
		return new(GeometricGraphics.Square)
	default:
		return new(UnknownFigure)
	}
}

func NewGeometricFactory() *GeometricFactory {
	return new(GeometricFactory)
}