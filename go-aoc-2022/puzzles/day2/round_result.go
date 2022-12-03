package day2

type RoundResult struct {
	slug string
}

func (r RoundResult) String() string {
	return r.slug
}

var (
	UnknownResult = RoundResult{""}
	Loss          = RoundResult{"X"}
	Draw          = RoundResult{"Y"}
	Win           = RoundResult{"Z"}
)

func ResultFromString(s string) RoundResult {
	switch s {
	case Loss.slug:
		return Loss
	case Draw.slug:
		return Draw
	case Win.slug:
		return Win
	}

	return UnknownResult
}

func RoundResultPoints(r RoundResult) int {
	switch r {
	case Loss:
		return 0
	case Draw:
		return 3
	case Win:
		return 6

	}

	return -1
}
