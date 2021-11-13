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

	if TraditionalToSimplified("吃飯") != "吃饭" {
		t.Fatal()
	}
	if SimplifiedToTraditional("吃饭") != "吃飯" {
		t.Fatal()
	}

	if !ContainsTraditional("吃飯") {
		t.Fatal()
	}
	if ContainsTraditional("吃饭") {
		t.Fatal()
	}
}
