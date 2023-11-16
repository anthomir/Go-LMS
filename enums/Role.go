package enums

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleUser    Role = "user"
	RoleTeacher Role = "teacher"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleUser, RoleTeacher:
		return true
	}

	return false
}
