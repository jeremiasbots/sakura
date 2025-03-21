package http

import (
	"io"
	goHttp "net/http"

	"github.com/dop251/goja"
)

var HTTPObject = map[string]interface{}{
	"fetch": Fetch,
}

func Fetch(call goja.FunctionCall, rt *goja.Runtime) goja.Value {
	url := call.Argument(0)
	urlString := url.String()
	response, err := goHttp.Get(urlString)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	textResponse := string(body)
	return rt.ToValue(textResponse)
}
