package managers

import (
	"fama/numbers/ports"
)

type numbersManager struct {
	converter ports.Converter
}

func newNumbersManager(c ports.Converter) ports.NumbersManager {
	return &numbersManager{
		converter: c,
	}
}

func (n *numbersManager) ToWords(number int, lang string) (string, error) {
	err := n.converter.IsSupported(lang)
	if err != nil {
		return "", err
	}
	return n.converter.To(number, lang), nil
}
