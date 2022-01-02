package product

import "github.com/andiahmads/design-patterns/abstract_factory"

type MejaIndomart struct{}

func (m *MejaIndomart) Price() float64 {
	return 8000
}

func (m *MejaIndomart) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{
		Length: 4000,
		Width:  900,
		Height: 800,
	}
}

func (m *MejaIndomart) IsFoldable() bool {
	return true
}
