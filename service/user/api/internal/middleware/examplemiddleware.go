package middleware

import (
	"fmt"
	"net/http"
)

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("中间件")
		//_, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	fmt.Println(err)
		//}
		// Passthrough to next handler if need
		next(w, r)
	}
}
