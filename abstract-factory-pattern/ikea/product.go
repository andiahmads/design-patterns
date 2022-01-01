package ikea

type Leifarne struct{}

//definiskan spesifikasi berdasarkan produk
func (l *Leifarne) Price() float64 {
	return 1095000
}
func (l *Leifarne) IsIoTEnabled() bool {
	return false
}
func (l *Leifarne) IsSoft() bool {
	return false
}

//buat produk coffeetable
type MejaKocok struct{}

func (m *MejaKocok) Price() float64 {
	return 29000
}
func (m *MejaKocok) Size() Dimension {}
