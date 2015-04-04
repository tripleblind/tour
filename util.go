package gdp

import (
	"crypto/rand"
	"math/big"
)

func bdice(up, exclude int) int {

	for {

		if r := dice(up); r != exclude {
			return r
		}

	}

}

func dice(up int) int {

	if n, err := rand.Int(rand.Reader, big.NewInt(int64(up))); err != nil {
		panic(err)
	} else {
		return int(n.Int64())
	}

}