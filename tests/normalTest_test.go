package tests

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
