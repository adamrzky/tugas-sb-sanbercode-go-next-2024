package geometri

import (
	"fmt"
	"math"
	"strings"
)

// Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

type Tabung struct {
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2 * (b.Panjang*b.Lebar + b.Lebar*b.Tinggi + b.Tinggi*b.Panjang))
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

// Soal 2
type Phone struct {
	Name   string
	Brand  string
	Year   int
	Colors []string
}

type Displayable interface {
	DisplayInfo() string
}

func (p Phone) DisplayInfo() string {
	return fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.Name, p.Brand, p.Year, strings.Join(p.Colors, ", "))
}

// Soal 3
func LuasPersegi(sisi int, cetakDetail bool) interface{} {
	if sisi == 0 {
		if cetakDetail {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi
	if cetakDetail {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm²", sisi, luas)
	}
	return luas
}

// Soal 5
type Person struct {
	Name, Job, Gender string
	Age               int
}

type Sentence interface {
	Introduction() string
}

func (p Person) Introduction() string {
	prefix := "Pak"
	if p.Gender == "female" {
		prefix = "Bu"
	}
	return fmt.Sprintf("%s %s adalah %s yang berusia %d tahun", prefix, p.Name, p.Job, p.Age)
}
