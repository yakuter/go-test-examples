package main

import (
	"testing"
)

func Topla(sayilar []int) int {
	toplam := 0
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}

	return toplam
}

func TestTopla(t *testing.T) {
	sonuc := Topla([]int{2, 5})
	if sonuc != 7 {
		t.Error("Beklenen sonuc 6, elde edilen ", sonuc)
	}
}

var tableCarpmaDegerleri = []struct {
	girdix        int
	girdiy        int
	beklenenSonuc int
}{
	{2, 2, 4},
	{5, 3, 15},
	{8, 4, 32},
	{12, 5, 60},
}

func Carpma(sayilar []int) int {
	carpim := 1
	for i := range sayilar {
		carpim = carpim * sayilar[i]
	}
	return carpim
}

//Table Test (Normal)
func TestCarpma(t *testing.T) {
	for _, deger := range tableCarpmaDegerleri {
		eldeEdilenSonuc := Carpma([]int{deger.girdix, deger.girdiy})
		if eldeEdilenSonuc != deger.beklenenSonuc {
			t.Errorf("Beklenen sonuc %d, elde edilen %d", deger.beklenenSonuc, eldeEdilenSonuc)
		}
	}
}

var tableCikarmaDegerleri = []struct {
	girdix        int
	girdiy        int
	beklenenSonuc int
}{
	{2, 2, 0},
	{5, 3, 2},
	{8, 4, 4},
	{12, 5, 7},
}

func Cikarma(x, y int) int {
	return x - y
}

// Table Test (Subtest Mantığıyla)
func TestCikarma(t *testing.T) {
	for _, deger := range tableCikarmaDegerleri {
		testIsmi := string(deger.girdix) + "x" + string(deger.girdiy)
		t.Run(testIsmi, func(t *testing.T) {
			eldeEdilenSonuc := Cikarma(deger.girdix, deger.girdiy)
			if eldeEdilenSonuc != deger.beklenenSonuc {
				t.Errorf("Beklenen sonuc %d, elde edilen %d", deger.beklenenSonuc, eldeEdilenSonuc)
			}
		})
	}
}

func TestSubtests(t *testing.T) {
	// <setup code>
	t.Run("A=1", func(t *testing.T) {})
	t.Run("A=2", func(t *testing.T) {})
	t.Run("B=1", func(t *testing.T) {})
	// <tear-down code>
}

func TestTLog(t *testing.T) {
	t.Parallel() // marks TLog as capable of running in parallel with other tests
	tests := []struct {
		name string
	}{
		{"test 1"},
		{"test 2"},
		{"test 3"},
		{"test 4"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Error(tt.name)
		})
	}
}
