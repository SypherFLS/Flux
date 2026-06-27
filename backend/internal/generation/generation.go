package generation

import (
	"math/rand"
	"FLUX/internal/generation/payback"
	"time"
	"fmt"
)


type Roulet interface {
	Start()
	Roll() payback.Payback
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func Roll() payback.Payback{
	// red := []int{1, 3, 5, 7, 9, 12, 14, 16, 18, 21, 23, 25, 27, 28, 30, 32, 34, 36}
	// black := []int{2, 4, 6, 8, 10, 11, 13, 15, 17, 19, 20, 22, 24, 26, 29, 31, 33, 35}
	// green := []int{0}

	colors := []string{"red", "black", "green"}

	var pb payback.Payback 

	seed := time.Now().UnixNano()
	fmt.Println(seed)

	pb.Color = colors[rng.Intn(len(colors))]

	return pb
}



