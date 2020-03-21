package tests

import (
	"testing"
)

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
