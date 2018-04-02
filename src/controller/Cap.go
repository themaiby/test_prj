/* just empty controller */
package controller

import "net/http"

type capController struct{}

var CapController = new(capController)

func (this *capController) Cap(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Empty controller"))
}
