package product

import "github.com/andiahmads/design-patterns/abstract_factory"

type Sofakocok struct{}

func (s *Sofakocok) Price() float64 {
	return 90000
}

func (s *Sofakocok) Style() abstract_factory.SofaStyle {
	return abstract_factory.EuropeanStyle
}
