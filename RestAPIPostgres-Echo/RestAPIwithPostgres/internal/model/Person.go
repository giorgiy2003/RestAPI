package Model

type Person struct {
	Id        int `json:"id" db:"peson_id"`
	Email     string `json:"email" db:"peson_email"`
	Phone     string `json:"phone" db:"peson_phone"`
	FirstName string `json:"firstName" db:"peson_firstName"`
	LastName  string `json:"lastName" db:"peson_lastName"`
}
