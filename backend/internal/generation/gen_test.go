package generation 

import (
	"testing"
	"time"
)

func Test_gen(t *testing.T) { 
	temp := NewRouletteGenerator()

	gen := Generator(temp)

	for i :=0;  i < 4;  i++ {
		res, _ := gen.Generate(t.Context())
		t.Log(res) 
		time.Sleep(time.Millisecond*3)
	}
	
}