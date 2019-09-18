package pb

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Add int // +
	Sub int // -
	Mul int // *
}
