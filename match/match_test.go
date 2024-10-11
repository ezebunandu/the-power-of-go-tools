package match_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/ezebunandu/match"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMatch_PrintsMatchingLinesToGivenWriter(t *testing.T){
	t.Parallel()
    args, err := os.Open("testdata/three_lines.txt") 
    if err != nil {
        fmt.Println(err)
        return
    }
	buf := new(bytes.Buffer)
	m, err := match.NewMatcher(match.WithInput(args), match.WithOutput(buf), match.WithSearchStringFromArgs([]string{"DUST"}))
    if err != nil {
        fmt.Println(err)
        return
    }
	m.PrintMatchingLines()
	want := "are sprinkled with magic DUST\na special kind of DUST\nDUST like no other"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func Test(t *testing.T){
    t.Parallel()
    testscript.Run(t, testscript.Params{
        Dir: "testdata/script",
    })
}

func TestMain(m *testing.M){
    os.Exit(testscript.RunMain(m, map[string]func() int{
        "match": match.Main,
    }))
}
