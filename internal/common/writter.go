package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendInternalErr is function to set internal server error, with 500 header
func SendInternalErr(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte(`{'status': 500, 'message': 'internal server error'}`))
}

// SendOK is function to set success response, with 200 header
func SendOk(w http.ResponseWriter) {
	w.WriteHeader(200)
	w.Write([]byte(`{'status': 200, 'message': 'ok'}`))
}

// SendBadRequest is function to set bad request, with 400 header and message
func SendBadRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(400)
	w.Write([]byte(fmt.Sprintf(`{'status':400, 'message': '%s'}`, message)))
}

// SendOKWithData is function to set success, with 200 header, and data
func SendOkWithData(w http.ResponseWriter, any interface{}) {
	byteData, _ := json.Marshal(any)

	w.WriteHeader(200)
	w.Write(byteData)
}
