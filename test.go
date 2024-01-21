package httptest

import (
	"reflect"
	"testing"

	"github.com/qiniu/dyn/flag"
	"github.com/qiniu/httptest"
	"github.com/qiniu/httptest/exec"
	"github.com/qiniu/x/log"
)

type Cmd struct {
	Code []string
	Data [][]string
	exec.Context
}

func (p *Cmd) Run(stage string) {
	log.SetFlags(log.Llevel | log.LstdFlags)
	tests := []testing.InternalTest{
		{"main", p.allTests},
	}

	testing.Main(p.allMatch, tests, nil, nil)
	// testing.RunTests(p.allMatch, tests)
}

func (p *Cmd) allMatch(pat, str string) (bool, error) {
	return true, nil
}

func (p *Cmd) allTests(t *testing.T) {
	ctx := httptest.New(t)
	v := reflect.ValueOf(p)
	for i := range p.Code {
		if p.Code[i] == "case" {
			ctx.Log("==========", p.Data[i], "===========")
		} else if p.Code[i] == "tearDown" {
			// TODO: how to impl tearDown
			if i < len(p.Code)-2 {
				ctx.Fatal("TearDown Failure")
			}
		} else {
			method := v.MethodByName("Cmd_" + p.Code[i])
			flag.ExecMethod(ctx.Context, method, reflect.ValueOf(ctx), p.Data[i])
		}
	}
}

type assertArgs struct {
	BoolArg string `arg:"assert bool"`
}

func (p *Cmd) Cmd_assert(ctx *httptest.Context, args *assertArgs) {
	if args.BoolArg != "1" {
		ctx.Fatal("Assert Result is False")
	}
}
