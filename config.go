package fsnotify

import (
	cond "github.com/vela-ssoc/vela-cond"
	auxlib2 "github.com/vela-ssoc/vela-kit/auxlib"
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/pipe"
)

type config struct {
	name  string
	path  []string
	match *cond.Cond
	pipe  *pipe.Px
	onErr *pipe.Px
	co    *lua.LState
}

func newConfig(L *lua.LState) *config {
	tab := L.CheckTable(1)
	cfg := &config{
		co:    xEnv.Clone(L),
		pipe:  pipe.New(pipe.Env(xEnv)),
		onErr: pipe.New(pipe.Env(xEnv)),
	}

	tab.Range(func(key string, val lua.LValue) {
		switch key {
		case "name":
			cfg.name = auxlib2.CheckProcName(val, L)
		case "path":
			switch val.Type() {
			case lua.LTString:
				cfg.path = []string{val.String()}
			case lua.LTTable:
				cfg.path = auxlib2.LTab2SS(val.(*lua.LTable))
			default:
				//todo
			}
		default:
			//todo
		}
	})

	if err := cfg.valid(); err != nil {
		L.RaiseError("%v", err)
		return nil
	}
	return cfg
}

func (cfg *config) valid() error {
	if e := auxlib2.Name(cfg.name); e != nil {
		return e
	}

	return nil
}
