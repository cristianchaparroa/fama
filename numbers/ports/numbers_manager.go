package ports

// NumbersManager is in charge to define the
// methods to convert the numbers to words
type NumbersManager interface {

	// ToWords generates the representation of number in words
	ToWords(number int, lang string) (string, error)
}
