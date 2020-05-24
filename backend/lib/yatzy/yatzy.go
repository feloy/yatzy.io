package yatzy

const (
	ACES = iota
	TOWS
	THREES
	FOURS
	FIVES
	SIXES
	THREE_IN_A_ROW
	FOUR_IN_A_ROW
	FULL_HOUSE
	SMALL_STRAIGHT
	LARGE_STRAIGHT
	YATZY
	CHANCE
)

func countByValue(die []int) []int {
	ret := []int{0, 0, 0, 0, 0, 0}
	for _, dice := range die {
		ret[dice-1]++
	}
	return ret
}

func ComputePoints(formula int, die []int) int {
	if formula <= SIXES {
		var sum = 0
		for _, dice := range die {
			if dice == int(formula+1) {
				sum += dice
			}
		}
		return sum
	} else if formula == THREE_IN_A_ROW {
		// Three in a row: sum of 3 or more identical die
		counts := countByValue(die)
		for i, count := range counts {
			if count > 2 {
				return (i + 1) * count
			}
		}
		return 0
	} else if formula == FOUR_IN_A_ROW {
		// Four in a row: sum of 4 or more identical die
		counts := countByValue(die)
		for i, count := range counts {
			if count > 3 {
				return (i + 1) * count
			}
		}
		return 0
	} else if formula == FULL_HOUSE {
		counts := countByValue(die)
		threes := 0
		twos := 0
		yatzys := 0
		for _, count := range counts {
			switch count {
			case 5:
				yatzys++
			case 3:
				threes++
			case 2:
				twos++
			}
		}
		if yatzys == 1 || (threes == 1 && twos == 1) {
			return 25
		}
		return 0
	} else if formula == SMALL_STRAIGHT {
		counts := countByValue(die)
		length := 0
		for _, count := range counts {
			if count > 0 {
				length++
				if length > 3 {
					return 30
				}
			} else {
				length = 0
			}
		}
		return 0
	} else if formula == LARGE_STRAIGHT {
		counts := countByValue(die)
		length := 0
		for _, count := range counts {
			if count > 0 {
				length++
				if length > 4 {
					return 40
				}
			} else {
				length = 0
			}
		}
		return 0
	} else if formula == YATZY {
		for _, dice := range die[1:] {
			if dice != die[0] {
				return 0
			}
		}
		return 50
	} else if formula == CHANCE {
		var sum = 0
		for _, dice := range die {
			sum += dice
		}
		return sum
	} else {
		return 42
	}
}
