package product

import "github.com/andiahmads/design-patterns/abstract_factory"

type SofaIndomart struct{}

func (s *SofaIndomart) Price() float64 {
	return 9000
}

func (s *SofaIndomart) Style() abstract_factory.SofaStyle {
	return abstract_factory.AmericanStyle
}
