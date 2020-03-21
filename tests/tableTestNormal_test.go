package tests

import (
	"testing"
)

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
