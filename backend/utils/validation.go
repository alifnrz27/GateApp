package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(validate *validator.Validate, req interface{}, validationMessages map[string]string) map[string]string {
	err := validate.Struct(req)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	// Ambil refleksi tipe yang tepat (harus struct, bukan pointer)
	var ref reflect.Type
	t := reflect.TypeOf(req)
	if t.Kind() == reflect.Ptr {
		ref = t.Elem()
	} else {
		ref = t
	}

	for _, e := range err.(validator.ValidationErrors) {
		field, ok := ref.FieldByName(e.StructField())
		var jsonKey string
		if ok {
			jsonTag := field.Tag.Get("json")
			jsonKey = strings.Split(jsonTag, ",")[0]
		} else {
			jsonKey = e.Field()
		}

		key := jsonKey + "." + e.Tag()
		msg, found := validationMessages[key]
		if !found {
			msg = e.Error()
		}
		errors[jsonKey] = msg
	}

	return errors
}
