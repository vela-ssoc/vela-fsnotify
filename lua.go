package fsnotify

import (
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-kit/lua"
	"reflect"
)

var (
	xEnv   vela.Environment
	typeof = reflect.TypeOf((*watch)(nil)).String()
)

/*

 */

func newLuaFsnotify(L *lua.LState) int {
	cfg := newConfig(L)
	proc := L.NewVelaData(cfg.name, typeof)
	if proc.IsNil() {
		proc.Set(newWatch(cfg))
	} else {
		w := proc.Data.(*watch)
		xEnv.Free(w.cfg.co)
		w.cfg = cfg
	}

	L.Push(proc)
	return 1
}

func WithEnv(env vela.Environment) {
	xEnv = env
	env.Set("fsnotify", lua.NewFunction(newLuaFsnotify))
}
