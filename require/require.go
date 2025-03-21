package require

import (
	"os"

	"github.com/dop251/goja"
	"github.com/jeremiasbots/sakura/modules/ev3"
	"github.com/jeremiasbots/sakura/modules/http"
	"github.com/jeremiasbots/sakura/vm"
)

// Require implementa la función require() de JavaScript para cargar módulos
func Require(call goja.FunctionCall, rt *goja.Runtime) goja.Value {
	// Obtener el argumento (ruta del archivo)
	arg := call.Argument(0)
	value, ok := arg.Export().(string)
	if !ok {
		panic("The value must be a string: require()")
	}

	if value == "sakura:ev3" {
		return rt.ToValue(ev3.EV3Object)
	}

	if value == "sakura:http" {
		return rt.ToValue(http.HTTPObject)
	}

	file, err := os.ReadFile(value)
	if err != nil {
		panic(err)
	}
	jsVM := vm.New()

	result, err := jsVM.RunString(string(file))
	if err != nil {
		panic(err)
	}

	return result
}
