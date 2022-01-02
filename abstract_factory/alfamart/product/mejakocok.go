package product

import "github.com/andiahmads/design-patterns/abstract_factory"

//buat produk coffeetable
type MejaKocok struct{}

func (m *MejaKocok) Price() float64 {
	return 29000
}

func (m *MejaKocok) Size() abstract_factory.Dimension {
	return abstract_factory.Dimension{
		Length: 2190,
		Width:  200,
		Height: 300,
	}
}

func (m *MejaKocok) IsFoldable() bool {
	return false
}
