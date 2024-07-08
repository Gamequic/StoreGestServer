package middlewares

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidatorHandler(schemaType reflect.Type) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read request and dump it to read it again later
			bodyBytes, _ := ioutil.ReadAll(r.Body)

			// Safe a copy to use it later
			bodyCopy := make([]byte, len(bodyBytes))
			copy(bodyCopy, bodyBytes)

			// Set the body as it was before
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			// New instance of data
			data := reflect.New(schemaType).Interface()

			if err := json.NewDecoder(bytes.NewReader(bodyCopy)).Decode(&data); err != nil {
				panic(GormError{Code: 400, Message: err.Error(), IsGorm: true})
			}

			validate := validator.New()

			errorValidate := validate.Struct(data)
			if errorValidate != nil {
				for _, err := range errorValidate.(validator.ValidationErrors) {
					panic(GormError{Code: 400, Message: err.Error(), IsGorm: true})
				}
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
