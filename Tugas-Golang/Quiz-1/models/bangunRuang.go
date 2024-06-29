package models

import "math"

type Bangun interface {
	Hitung() float64
}

// Bangun Datar
type Persegi struct {
	Sisi float64
}

func (p *Persegi) HitungLuas() float64 {
	return p.Sisi * p.Sisi
}

func (p *Persegi) HitungKeliling() float64 {
	return 4 * p.Sisi
}

type PersegiPanjang struct {
	Panjang, Lebar float64
}

func (pp *PersegiPanjang) HitungLuas() float64 {
	return pp.Panjang * pp.Lebar
}

func (pp *PersegiPanjang) HitungKeliling() float64 {
	return 2 * (pp.Panjang + pp.Lebar)
}

type Lingkaran struct {
	JariJari float64
}

func (l *Lingkaran) HitungLuas() float64 {
	return math.Pi * l.JariJari * l.JariJari
}

func (l *Lingkaran) HitungKeliling() float64 {
	return 2 * math.Pi * l.JariJari
}

// Bangun Ruang
type Kubus struct {
	Sisi float64
}

func (k *Kubus) HitungVolume() float64 {
	return math.Pow(k.Sisi, 3)
}

func (k *Kubus) HitungLuasPermukaan() float64 {
	return 6 * (k.Sisi * k.Sisi)
}

type Balok struct {
	Panjang, Lebar, Tinggi float64
}

func (b *Balok) HitungVolume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}

func (b *Balok) HitungLuasPermukaan() float64 {
	return 2 * (b.Panjang*b.Lebar + b.Lebar*b.Tinggi + b.Panjang*b.Tinggi)
}

type Tabung struct {
	JariJari, Tinggi float64
}

func (t *Tabung) HitungVolume() float64 {
	return math.Pi * t.JariJari * t.JariJari * t.Tinggi
}

func (t *Tabung) HitungLuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}
