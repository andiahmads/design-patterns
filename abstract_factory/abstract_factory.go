package abstract_factory

type Harga interface {
	Price() float64
}

type Kursi interface {
	Harga
	IsIoTEnabled() bool
	IsSoft() bool
}

type Dimension struct {
	Length, Width, Height int
}

type Meja interface {
	Harga
	Size() Dimension
	IsFoldable() bool
}

type SofaStyle string

const (
	EuropeanStyle SofaStyle = "european"
	AmericanStyle SofaStyle = "american"
)

type Sofa interface {
	Harga
	Style() SofaStyle
}

type FurnitureFactory interface {
	BuatKursi() Kursi
	BuatMeja() Meja
	BuatSofa() Sofa
}
