package indomart

import (
	"github.com/andiahmads/design-patterns/abstract_factory"
	"github.com/andiahmads/design-patterns/abstract_factory/indomart/product"
)

type Indomart struct{}

func (i *Indomart) BuatKursi() abstract_factory.Kursi {
	return &product.KursiIndomart{}
}

func (i *Indomart) BuatMeja() abstract_factory.Meja {
	return &product.MejaIndomart{}
}
