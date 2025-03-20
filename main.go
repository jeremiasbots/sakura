package main

import (
	"os"

	"github.com/jeremiasbots/sakura/require"
	"github.com/jeremiasbots/sakura/vm"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: sakura <script>")
	}
	scriptPath := os.Args[1]
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		panic(err)
	}

	jsVM := vm.New()
	jsVM.Set("require", require.Require)

	_, err = jsVM.RunString(string(script))
	if err != nil {
		panic(err)
	}
}
