package two_crystal_ball_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/cherryramatisdev/algorithms-apocalipse/internal/algorithms/two_crystal_ball"
)

func TestHappy(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	data := make([]bool, 10000)
	idx := rand.Intn(10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	result1 := two_crystal_ball.TwoCrystalBalls(data)
	if result1 != idx {
		t.Errorf("Expected %d, got %d", idx, result1)
	}
}
