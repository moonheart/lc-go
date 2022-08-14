package binary_watch

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) (res []string) {
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return res
}

/*
1: 6
2: 5+4+3+2+1
3: 4+3+2+1
*/
