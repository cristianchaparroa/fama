package managers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
)

func TestNumbersManager_ToWords_English_OK(t *testing.T) {
	var test = []struct {
		Number   int
		Lang     string
		Expected string
	}{
		{-20, "en", "minus twenty"},
		{15, "en", "fifteen"},
		{35000, "en", "thirty-five thousand"},
		{9567878, "en", "nine million five hundred sixty-seven thousand eight hundred seventy-eight"},
		{maxInt, "en", "nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven"},
		{-1 * maxInt, "en", "minus nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven"},
	}

	converter := newConverter()
	manager := newNumbersManager(converter)

	for _, tc := range test {
		result, err := manager.ToWords(tc.Number, tc.Lang)
		assert.Nil(t, err)
		assert.Equal(t, tc.Expected, result)
	}
}

func TestNumbersManager_ToWords_Multilingual(t *testing.T) {
	var test = []struct {
		Number   int
		Lang     string
		Expected string
	}{
		{20, "fr", "vingt"},
		{20, "it", "venti"},
		{20, "es", "veinte"},
		{20, "se", "tjugo"},
		{20, "nl", "twintig"},
		{20, "tr", "yirmi"},
		{20, "pt", "vinte"},
		{20, "pl", "dwadzieścia"},
		{20, "ru", "двадцать"},
		{20, "ir", "بیست"},
		{20, "id", "dua puluh"},
		{20, "jp", "二十"},
	}

	converter := newConverter()
	manager := newNumbersManager(converter)

	for _, tc := range test {
		result, err := manager.ToWords(tc.Number, tc.Lang)
		assert.Nil(t, err)
		assert.Equal(t, tc.Expected, result)
	}
}

func TestNumbersManager_ToWords_NotSupported(t *testing.T) {

	converter := newConverter()
	manager := newNumbersManager(converter)

	_, err := manager.ToWords(666, "de")
	assert.NotNil(t, err)
	assert.Equal(t, LangNotSupported, err.Error())
}
