package managers

import (
	"errors"
	"fama/numbers/ports"
	ntw "moul.io/number-to-words"
)

type converter struct {
}

func newConverter() ports.Converter {
	return &converter{}
}

func (c *converter) IsSupported(language string) error {
	for _, l := range lang {
		if l == language {
			return nil
		}
	}
	return errors.New(LangNotSupported)
}

func (c *converter) To(number int, lang string) string {
	switch lang {
	case "en":
		return ntw.IntegerToEnUs(number)
	case "fr":
		return ntw.IntegerToFrFr(number)
	case "it":
		return ntw.IntegerToItIt(number)
	case "es":
		return ntw.IntegerToEsEs(number)
	case "se":
		return ntw.IntegerToSvSe(number)
	case "nl":
		return ntw.IntegerToNlNl(number)
	case "tr":
		return ntw.IntegerToTrTr(number)
	case "pt":
		return ntw.IntegerToPtPt(number)
	case "pl":
		return ntw.IntegerToPlPl(number)
	case "ru":
		return ntw.IntegerToRuRu(number)
	case "ir":
		return ntw.IntegerToIrIr(number)
	case "id":
		return ntw.IntegerToIDID(number)
	case "jp":
		return ntw.IntegerToJaJp(number)
	}
	return ""
}
