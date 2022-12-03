package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
First column of input is what the opponent is going to play
A = Rock
B = Paper
C = Scissors

Second column is what you should play in response:
X = Rock
Y = Paper
Z = Scissors

Total score is the sum of your scores for each round
Score for a round is the score for the shape selected
Rock = 1
Paper = 2
Scissors = 3

Plus the score of the outcome of the round
Loss = 0
Draw = 3
Win = 6

---

Tests
Given
```

	A Y
	B X
	C Z

```
- Round 1: Opponent chooses Rock, I choose Paper
- Score of 8 points for the win (6) and choosing paper (2)

- Round 2: Opponent chooses Paper, I choose Rock
- Score of 1 for losing(0) and choosing rock (1)
*/
func Step1() {
	rounds := readDay2Input()

	roundScores := []int{}
	for _, round := range rounds {
		roundScores = append(roundScores, calculateRound(round[0], round[1]))
	}

	scoreSum := sumScores(roundScores)

	fmt.Println("The total score was: ", scoreSum)
}

/*
	Step 2:
	The second column actually means what result you should go for based on the
	move your opponent makes.

	X = lose
	Y = Draw
	Z = Win

	Calculate the score based on this instead
*/

func Step2() {
	rounds := readDay2Input()

	roundScores := []int{}
	for _, round := range rounds {
		roundScores = append(roundScores, calculateRound2(round[0], round[1]))
	}

	scoreSum := sumScores(roundScores)

	fmt.Println("The total score was: ", scoreSum)
}

func readDay2Input() [][]string {
	f, err := os.Open("./res/day2-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	rounds := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		round := strings.Fields(line)
		rounds = append(rounds, round)
	}

	return rounds
}

func calculateRound(opponent string, me string) int {
	meMove := MeFromString(me)
	opponentPlay := MoveFromString(opponent)

	playScore := MovePoints(meMove)
	resultScore := calculateResultPoints(opponentPlay, meMove)

	return playScore + resultScore
}

func calculateRound2(opponent string, result string) int {
	opponentPlay := MoveFromString(opponent)
	roundResult := ResultFromString(result)
	myMove := calculateMyMove(roundResult, opponentPlay)

	playScore := MovePoints(myMove)
	return playScore + RoundResultPoints(roundResult)
}

func calculateResultPoints(opponentPlay Move, mePlay Move) int {
	switch mePlay {
	case Rock:
		switch opponentPlay {
		case Rock:
			return RoundResultPoints(Draw)
		case Paper:
			return RoundResultPoints(Loss)
		case Scissors:
			return RoundResultPoints(Win)
		}
	case Paper:
		switch opponentPlay {
		case Rock:
			return RoundResultPoints(Win)
		case Paper:
			return RoundResultPoints(Draw)
		case Scissors:
			return RoundResultPoints(Loss)
		}
	case Scissors:
		switch opponentPlay {
		case Rock:
			return RoundResultPoints(Loss)
		case Paper:
			return RoundResultPoints(Win)
		case Scissors:
			return RoundResultPoints(Draw)
		}
	}

	return -1
}

func calculateMyMove(result RoundResult, opponent Move) Move {
	switch result {
	case Win:
		switch opponent {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	case Draw:
		switch opponent {
		case Rock:
			return Rock
		case Paper:
			return Paper
		case Scissors:
			return Scissors
		}
	case Loss:
		switch opponent {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}

	return UnknownMove
}

func sumScores(scores []int) int {
	scoreSum := 0

	for _, score := range scores {
		scoreSum += score
	}

	return scoreSum
}
