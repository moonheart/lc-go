package main

import "math"

type State int
type CharType int

const (
	StateStart State = iota
	StateSigned
	StateInNumber
	StateEnd
)

const (
	CharSpace CharType = iota
	CharSign
	CharNumber
	CharOther
)

func getCharType(char byte) CharType {
	switch char {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case ' ':
		return CharSpace
	case '+', '-':
		return CharSign
	default:
		return CharOther
	}
}

func strToInt(str string) int {
	transfer := map[State]map[CharType]State{
		StateStart: {
			CharSpace:  StateStart,
			CharSign:   StateSigned,
			CharNumber: StateInNumber,
			CharOther:  StateEnd,
		},
		StateSigned: {
			CharNumber: StateInNumber,
			CharSign:   StateEnd,
			CharSpace:  StateEnd,
			CharOther:  StateEnd,
		},
		StateInNumber: {
			CharNumber: StateInNumber,
			CharSign:   StateEnd,
			CharSpace:  StateEnd,
			CharOther:  StateEnd,
		},
		StateEnd: {
			CharNumber: StateEnd,
			CharSign:   StateEnd,
			CharSpace:  StateEnd,
			CharOther:  StateEnd,
		},
	}
	state := StateStart
	sign := 1
	num := 0
	for i := 0; i < len(str); i++ {
		typ := getCharType(str[i])
		if s, ok := transfer[state][typ]; !ok {
			break
		} else {

			state = s
			if state == StateEnd {
				break
			} else {
				if typ == CharSign && str[i] == '-' {
					sign = -1
				} else if typ == CharNumber {
					if num > math.MaxInt32/10 || (num == math.MaxInt32/10 && str[i] > '7') {
						if sign == 1 {
							return math.MaxInt32
						} else {
							return math.MinInt32
						}
					}
					num = num*10 + int(str[i]) - '0'
				}
			}
		}
	}
	return sign * num
}
