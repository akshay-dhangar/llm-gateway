package core

import "testing"

func TestChoose(t *testing.T) {
	ps := []Provider{{"cheap", []string{"m"}, true, 1, 1, 1}, {"ready", []string{"m"}, true, 0, 2, 5}, {"down", []string{"m"}, false, 0, 2, 0}}
	p, e := Choose(ps, "m")
	if e != nil || p.Name != "ready" {
		t.Fatalf("%+v %v", p, e)
	}
	if _, e = Choose(ps, "unknown"); e == nil {
		t.Fatal("unsupported model accepted")
	}
}
