package middleware

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/mbaezahuenupil/go-mongodb-test/src/util"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		outBody, err := json.Marshal(r.Body)
		if err != nil {
			panic(err)
		}
		outHeader, err := json.Marshal(r.Header)
		if err != nil {
			panic(err)
		}

		request := fmt.Sprintf("Request: [%s%s] - RemoteAddr: [%s] - Method: [%s] - Header: [%v] - Body[%v]", r.Host, r.RequestURI, r.RemoteAddr, r.Method, string(outHeader), string(outBody))
		util.LOGGER.Infoln(request)
		next.ServeHTTP(w, r)
	})
}
