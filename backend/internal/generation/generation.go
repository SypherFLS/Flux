package generation

import (
	"math/rand"
	"FLUX/internal/generation/payback"
	"time"
	_ "fmt"
)

var numberToColor = map[int]string{
	0: "green",

	1: "red",
	2: "black",
	3: "red",
	4: "black",
	5: "red",
	6: "black",
	7: "red",
	8: "black",
	9: "red",
	10: "black",

	11: "black",
	12: "red",
	13: "black",
	14: "red",
	15: "black",
	16: "red",
	17: "black",
	18: "red",

	19: "red",
	20: "black",
	21: "red",
	22: "black",
	23: "red",
	24: "black",
	25: "red",
	26: "black",
	27: "red",
	28: "black",

	29: "black",
	30: "red",
	31: "black",
	32: "red",
	33: "black",
	34: "red",
	35: "black",
	36: "red",
}

type RouletInter interface {
	Start()
	Roll() payback.Payback
}


func StartRoller() {
	var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		defer ticker.Stop()

		for {
			Roll(rng)
			<-ticker.C
		}
	}()
}


func Roll(rng *rand.Rand) payback.Payback{
	var pb payback.Payback

	num := rng.Intn(37)
	pb.Number = num
	pb.Color = numberToColor[num]

	return pb
}

