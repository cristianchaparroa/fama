package ports

type Converter interface {
	IsSupported(lang string) error

	To(number int, lang string) string
}
