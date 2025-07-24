package ports

type IHasher interface {
	HashPassword(pwd string) (string, error)
	ComparePassword(hashed, password string) error
}
