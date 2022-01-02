package product

type Kursikocok struct{}

//definiskan spesifikasi berdasarkan produk
func (k *Kursikocok) Price() float64 {
	return 1095000
}
func (k *Kursikocok) IsIoTEnabled() bool {
	return false
}
func (k *Kursikocok) IsSoft() bool {
	return false
}
