package users

type User struct {
	ID        int64  "json:user_id"
	FirstName string "json:first_name"
	LastName  string "json:last_name"
	Email     string "json:email"
	Status    string "json:status"
}
