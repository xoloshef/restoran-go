package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := "Hello world!"

	fmt.Fprint(rw, text)
}
