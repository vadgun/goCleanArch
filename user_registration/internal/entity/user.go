package entity

type User struct {
	User     string `json:"usuario"`
	Password string `json:"contraseña"`
	Email    string `json:"correo"`
	Phone    string `json:"telefono"`
}
