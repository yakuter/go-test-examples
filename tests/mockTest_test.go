package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/* Gerçek Fonksiyonlar ********** */
type Matematik interface {
	MockTopla([]int) (int, error)
}

type islem struct{}

func (*islem) MockTopla(sayilar []int) (int, error) {
	toplam := 0
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}
	return toplam, nil
}

/* Test işlemleri ********** */
type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) MockTopla(sayilar []int) (int, error) {
	// Called, mock objesine bir methodun çağırıldığını haver verir ve döndürmek üzere bir dizi (array) değer alır
	args := mock.Called(sayilar)
	result := args.Get(0)

	return result.(int), args.Error(1)
	// Veya
	// return args.Int(0), args.Error(1)
	// Diğer seçenekler: args.Int(0) args.Bool(1) args.String(2)
}

func TestMockTopla(t *testing.T) {

	// Test objemizin bir kopyasını oluşturalım
	mockRepo := new(MockRepository)

	// setup expectations
	mockRepo.On("MockTopla", []int{2, 3}).Return(5, nil)

	// call the code we are testing
	testMatematik := Matematik(mockRepo)

	sonuc, err := testMatematik.MockTopla([]int{2, 3})

	// Beklentilerin karşılandığını iddia et (assert)
	mockRepo.AssertExpectations(t)

	assert.Equal(t, 5, sonuc)
	assert.Nil(t, err)
}
