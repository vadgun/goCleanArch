package pkg

import "strings"

func CreateMissingFieldsMessage(missingFields []string) string {
	msg := "Falta el campo "
	if len(missingFields) > 1 {
		msg = "Faltan los campos "
	}
	return msg + strings.Join(missingFields, ", ")
}
