package alfamart

import (
	"github.com/andiahmads/design-patterns/abstract_factory"
	"github.com/andiahmads/design-patterns/abstract_factory/alfamart/product"
)

//buat Factory
type Ikea struct{}

func (i *Ikea) BuatKursi() abstract_factory.Kursi {
	return &product.Kursikocok{}

}
func (i *Ikea) BuatMeja() abstract_factory.Meja {
	return &product.MejaKocok{}
}
func (i *Ikea) BuatSofa() abstract_factory.Sofa {
	return &product.Sofakocok{}

}
