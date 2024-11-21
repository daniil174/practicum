package fullname

// User — пользователь в системе.
type UserT struct {
	FirstName string
	LastName  string
}

// FullName возвращает имя и фамилию пользователя.
func (u UserT) FullName() string {
	return u.FirstName + " " + u.LastName
}
