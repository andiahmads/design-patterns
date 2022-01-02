package product

type KursiIndomart struct {
	price float64
}

func (k *KursiIndomart) Price() float64 {
	return k.price
}

func (k *KursiIndomart) IsIoTEnabled() bool {
	return true
}

func (k *KursiIndomart) IsSoft() bool {
	return false
}
