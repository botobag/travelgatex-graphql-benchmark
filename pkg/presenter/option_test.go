package presenter

import (
	"encoding/json"
	"testing"
)

func TestGenerator_Gen(t *testing.T) {
	g := NewOptionsGen()
	n := 10
	result := g.Gen(n)
	if len(result) != n {
		t.Fatalf("wanted to generate %v options, got %v", n, len(result))
	}

	expected := g.Gen(n)
	if len(result) != len(expected) {
		t.Fatal("len of generated options doesn't match")
	}

	resultData, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("can't marshal: %v", err)
	}
	expectData, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("can't marshal: %v", err)
	}
	if string(resultData) != string(expectData) {
		t.Fatal("generated option are not equal")
	}
	t.Log(string(resultData))
}
