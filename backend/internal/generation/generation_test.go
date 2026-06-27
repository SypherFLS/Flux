package generation

import (
	"testing"
	"time"
)

func Test_Roll(t *testing.T) {
	for i := 0; i < 5; i++ {
		a := Roll() 
		t.Logf("\n %v \n", a)
		time.Sleep(10 * time.Millisecond)
	}
}

func Test_Time(t *testing.T) {
    for i := 0; i < 5; i++ {
        t.Log(time.Now().UnixNano())
		time.Sleep(10 * time.Millisecond)
    }
}