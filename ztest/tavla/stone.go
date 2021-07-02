package calc

const (
	EmWHITE = 1 //白子
	EmBLACK = 2 //黑子
)

var tPos = [15]stPos{
	{1, EmWHITE},
	{1, EmWHITE},
	{12, EmWHITE},
	{12, EmWHITE},
	{12, EmWHITE},
	{12, EmWHITE},
	{12, EmWHITE},
	{17, EmWHITE},
	{17, EmWHITE},
	{17, EmWHITE},
	{19, EmWHITE},
	{19, EmWHITE},
	{19, EmWHITE},
	{19, EmWHITE},
	{19, EmWHITE},
}

type stPos struct {
	nRow   int32
	emType int32
}

type Stone struct {
	nId    int32
	nRow   int32
	emType int32
}

func (s *Stone) init(id int32) {
	s.nId = id
	if id < _maxCampStone {
		s.nRow = tPos[id].nRow
		s.emType = tPos[id].emType
	} else {
		s.nRow = _maxRow - tPos[id-_maxCampStone].nRow
		s.emType = EmBLACK + EmWHITE - tPos[id-_maxCampStone].emType
	}
}

func (s *Stone) GetId() int32 {
	return s.nId
}

func (s *Stone) GetRow() int32 {
	return s.nRow
}

func (s *Stone) GetEmType() int32 {
	return s.emType
}

func (s *Stone) inBag() bool {
	if s.emType == EmWHITE {
		return s.nRow >= _maxRow
	}
	return s.nRow <= 0
}

func (s *Stone) inHome() bool {
	if s.emType == EmWHITE {
		return s.nRow >= 19 && s.nRow <= 24
	}
	return s.nRow >= 1 && s.nRow <= 6
}

func (s *Stone) move(x int32) (from, to int32) {
	from = s.nRow
	_, to = s.calcMoveTo(x)
	s.nRow = to
	return
}

func (s *Stone) eat() (from, to int32) {
	from = s.nRow
	if s.emType == EmWHITE {
		s.nRow = 0
	} else if s.emType == EmBLACK {
		s.nRow = _maxRow
	}
	to = s.nRow
	return
}

func (s *Stone) calcMoveTo(x int32) (inBag bool, to int32) {
	if s.emType == EmWHITE {
		if to = s.nRow + x; to >= _maxRow {
			inBag = true
			to = _maxRow
		}
	}

	if s.emType == EmBLACK {
		if to = s.nRow - x; to <= 0 {
			inBag = true
			to = 0
		}
	}
	return
}
