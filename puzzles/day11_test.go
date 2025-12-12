package puzzles

import "testing"

func TestParseInput(t *testing.T) {
	test := `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`
	graph := parseInput(test)
	got := graph["you"]

	if len(got) != 2 {
		t.Errorf("Expected %d got %d", 2, len(got))
	}

	want := []string{"bbb", "ccc"}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Expected %q got %q", want[i], got[i])
		}
	}

}

func TestDfs(t *testing.T) {
	test := `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`
	graph := parseInput(test)
	cache := make(map[string]int)
	if got, want := dfs(graph, "svr", "out", cache, false, false), 2; got != want {
		t.Errorf("Expected %d got %d", want, got)
	}
}
