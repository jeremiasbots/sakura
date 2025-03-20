package console

import (
	"fmt"
	"log"

	"github.com/dop251/goja"
)

var ConsoleObject = map[string]interface{}{
	"log":    Log,
	"assert": Assert,
}

func Log(call goja.FunctionCall) goja.Value {
	for _, arg := range call.Arguments {
		fmt.Print(arg.String(), " ")
	}
	fmt.Println()
	return goja.Undefined()
}

func Assert(call goja.FunctionCall) goja.Value {
	booleanArg := call.Argument(0)
	goBoolean, not := booleanArg.Export().(bool)
	if !not {
		panic("The value must be a boolean")
	}
	if !goBoolean {
		log.Fatalf("Assertion failed")
	}
	return goja.Undefined()
}
