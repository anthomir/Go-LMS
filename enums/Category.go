package enums

type Category string

const (
	Programming Category = "Programming"
	Design      Category = "Design"
	Language    Category = "Language"
)

func (r Category) IsValid() bool {
	switch r {
	case Programming, Design, Language:
		return true
	}

	return false
}
