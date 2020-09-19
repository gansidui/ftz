package ftz

import (
	"testing"
)

func TestFTZ(t *testing.T) {
	simplifiedString := "我爱学习，我爱吃饭。"
	traditionalString := SimplifiedToTraditional(simplifiedString)

	if TraditionalToSimplified(traditionalString) != simplifiedString {
		t.Fatal()
	}
}
