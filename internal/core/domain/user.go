package domain

type User struct {
	ID          int
	Version     int
	FullName    string
	Phonenumber *string
}

func NewUser(id, version int, fullName string, phoneNumber *string) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		Phonenumber: phoneNumber,
	}
}
func NewUserUninitialized(fullName string, phoneNumber *string) User {
	return NewUser(UninitializedID, UninitializedVersion, fullName, phoneNumber)
}
