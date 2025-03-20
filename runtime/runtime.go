package runtime

import (
	"github.com/dop251/goja"
	"github.com/jeremiasbots/sakura/vm"
)

func RunScript(code string) (goja.Value, *goja.Runtime, error) {
	jsVM := vm.New()
	value, err := jsVM.RunString(code)

	return value, jsVM.Runtime, err
}
