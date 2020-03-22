package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertionTopla(sayilar []int) (int, error) {
	toplam := 0
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}

	return toplam, nil
}

func TestAssertionTopla(t *testing.T) {

	// Assert birden fazla kullanılacaksa
	// assert := assert.New(t)

	sonuc, err := AssertionTopla([]int{2, 5})

	// Eşittir
	assert.Equal(t, 7, sonuc, "sonuç eşit olmalı")

	// Eşit değildir
	assert.NotEqual(t, 6, sonuc, "Sonuçlar eşit olmamalı")

	// Sonuc nil olmalıdır
	assert.Nil(t, err)

	// Sonuç nil değilse
	err = errors.New("toplanmadi")
	if assert.NotNil(t, err) {
		// now we know that object isn't nil, we are safe to make
		// Sonucun nil olmadığını bildiğimiz için ilave testler yapabiliriz
		assert.Equal(t, "toplanmadi", err.Error())
	}
}
