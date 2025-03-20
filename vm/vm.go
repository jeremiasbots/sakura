package vm

import (
	"github.com/dop251/goja"
	"github.com/jeremiasbots/sakura/console"
)

type VM struct {
	Runtime *goja.Runtime
}

func New() *VM {
	runtime := goja.New()

	runtime.Set("console", console.ConsoleObject)

	return &VM{
		Runtime: runtime,
	}
}

func (vm *VM) RunString(code string) (goja.Value, error) {
	return vm.Runtime.RunString(code)
}
func (vm *VM) Set(name string, value interface{}) {
	vm.Runtime.Set(name, value)
}

func (vm *VM) Get(name string) goja.Value {
	return vm.Runtime.Get(name)
}
