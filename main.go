package main

import (
	"fmt"
	"strings"
)

func main() {
	classes()
	fmt.Println("\n--------------------------------")
	inheritance()
	fmt.Println("\n--------------------------------")
	encapsulation()
	fmt.Println("\n--------------------------------")
	abstraction()
	fmt.Println("\n--------------------------------")
	polymorphism()
	fmt.Println("\n--------------------------------")
	HOF()
	fmt.Println("\n--------------------------------")
	mrekursif()
	fmt.Println("\n--------------------------------")
	array()
}

func classes() {
	type murid struct {
		nama  string
		nilai int
	}
	var s1 = murid{}
	s1.nama = "Adhi"
	s1.nilai = 4

	var s2 = murid{"Aksal", 3}

	var s3 = murid{nama: "Kevin"}

	fmt.Println("\nClasses")
	fmt.Println("Nama murid 1 =", s1.nama)
	fmt.Println("Nilai murid 1 =", s1.nilai)
	fmt.Println("Nama murid 2 =", s2.nama)
	fmt.Println("Nilai murid 2 =", s2.nilai)
	fmt.Println("Nama murid 3 =", s3.nama)
	fmt.Println("Nilai murid 3 =", s3.nilai)
}

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println(b.title)
	fmt.Println("Konten: ", b.content)
	fmt.Println("Author: ", b.fullName())
	fmt.Println("Bio: ", b.bio)
}

func inheritance() {
	author1 := author{
		"Adhi",
		"Ardiansyah",
		"Golang Enthusiast",
	}
	blogPost1 := blogPost{
		"Inheritance",
		"Inheritance di Golang",
		author1,
	}
	fmt.Println()
	blogPost1.details()
}

func encapsulation() {
	fmt.Println()
	fmt.Println("Encapsulation")
	slc := []string{"Adhi", "Ardiansyah", "Kapital"}

	for x := 0; x < len(slc); x++ {

		res := strings.ToUpper(slc[x])
		// res := strings.toUpper(slc[x]) // -> error

		fmt.Println(res)
	}

}

type Segitiga struct {
	alas, tinggi float32
}

type Persegi struct {
	panjang float32
}

type PersegiPanjang struct {
	panjang, lebar float32
}

type Lingkaran struct {
	jarijari float32
}

func (t Segitiga) Luas() float32 {
	return 0.5 * t.alas * t.tinggi
}

func (l Persegi) Luas() float32 {
	return l.panjang * l.panjang
}

func (r PersegiPanjang) Luas() float32 {
	return r.panjang * r.lebar
}

func (c Lingkaran) Luas() float32 {
	return 3.14 * (c.jarijari * c.jarijari)
}

type Luas interface {
	Luas() float32
}

func abstraction() {
	fmt.Println()
	fmt.Println("Abstraction")

	t := Segitiga{alas: 15, tinggi: 25}
	s := Persegi{panjang: 5}
	r := PersegiPanjang{panjang: 5, lebar: 10}
	c := Lingkaran{jarijari: 5}

	var a Luas

	a = t
	fmt.Println("Luas dari Segitiga", a.Luas())

	a = s
	fmt.Println("Luas dari Persegi", a.Luas())

	a = r
	fmt.Println("Luas dari PersegiPanjang", a.Luas())

	a = c
	fmt.Println("Luas dari Lingkaran", a.Luas())
}

type Angka interface {
	Luas() float64
}

type PersegiPanjang2 struct {
	panjang float64
	lebar   float64
}

type Lingkaran2 struct {
	sisi float64
}

func (r PersegiPanjang2) Luas() float64 {
	luas := r.panjang * r.lebar
	return luas
}

func (s Lingkaran2) Luas() float64 {

	luas := s.sisi * s.sisi
	return luas
}

func polymorphism() {
	fmt.Println()
	fmt.Println("Polymorphism")

	persegipanjang := PersegiPanjang2{

		panjang: 10.5,
		lebar:   12.25,
	}

	lingkaran := Lingkaran2{
		sisi: 15.0,
	}

	// printing the calculated result
	fmt.Printf("Luas dari persegi panjang: %.3f satuan persegi.\n", persegipanjang.Luas())
	fmt.Printf("Luas dari lingkaran: %.3f satuan persegi.\n", lingkaran.Luas())
}

func penjumlahan(x, y int) int {
	return x + y
}
func jumlahSebagian(x int) func(int) int {
	return func(y int) int {
		return penjumlahan(x, y)
	}
}
func HOF() {
	fmt.Println()
	fmt.Println("Higher Order Function")
	partial := jumlahSebagian(3)
	fmt.Println(partial(7))
}

func rekursif(n int) int {
	if n == 0 {
		return 1
	}
	return n * rekursif(n-1)
}

func mrekursif() {
	fmt.Println()
	fmt.Println("Rekursif")
	fmt.Println(rekursif(7))

	var rekursif2 func(n int) int

	rekursif2 = func(n int) int {
		if n < 2 {
			return n
		}
		return rekursif2(n-1) + rekursif2(n-2)

	}

	fmt.Println(rekursif2(7))
}

func array() {
	fmt.Println()
	fmt.Println("ARRAY")
	var names [3]string
	names[0] = "Jean"
	names[1] = "Pierre"
	names[2] = "Polnareff"

	fmt.Println(names[0], names[1], names[2])

	fmt.Println("")

	var names2 = [3]string{"Jean", "Pierre", "Polnareff"}

	fmt.Println(len(names2))

	fmt.Println("")

	var names3 = [...]string{"Jean", "Pierre", "Polnareff", "Adhi"}

	fmt.Println(names3[2])

	fmt.Println("")

	var names4 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(names4)

	fmt.Println("")

	var names5 = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(names5); i++ {
		fmt.Println("berhitung", names5[i])
	}

	fmt.Println("")

	var names6 = [5]int{1, 2, 3, 4, 5}
	for index, value := range names6 {
		fmt.Println("index", index, "value", value)
	}
}
