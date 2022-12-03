package day2

type Move struct {
	slug string
}

func (m Move) String() string {
	return m.slug
}

var (
	UnknownMove = Move{""}
	Rock        = Move{"A"}
	Paper       = Move{"B"}
	Scissors    = Move{"C"}
)

func MoveFromString(s string) Move {
	switch s {
	case Rock.slug:
		return Rock
	case Paper.slug:
		return Paper
	case Scissors.slug:
		return Scissors
	}

	return UnknownMove
}

func MeFromString(s string) Move {
	switch s {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	}

	return UnknownMove
}

func MovePoints(m Move) int {
	switch m {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}

	return -1
}
