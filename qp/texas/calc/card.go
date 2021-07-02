package calc

const (
	_MaxCard   = 15
	_CardGroup = 13
)

type Card struct {
	Seq    int32
	color  int
	number int
	level  int
}

func (c *Card) calc() {
	if _, find := descDeck[c.Seq]; find {
		c.color = cardColor(c.Seq)
		c.number = cardNumber(c.Seq)
		c.level = cardLevel(c.Seq)
	}
}

func cardColor(seq int32) int {
	return int((seq - 1) / _CardGroup)
}

func cardNumber(seq int32) int {
	return int(seq % _CardGroup)
}

func cardLevel(seq int32) int {
	return number2Level(cardNumber(seq))
}

func number2Level(n int) int {
	if n >= 2 && n <= 13 {
		return n
	} else if n == 1 {
		return n + 13
	}
	return -1
}
