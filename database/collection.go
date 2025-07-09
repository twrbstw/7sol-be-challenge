package database

type DatabaseCollection string

const (
	COLLECTION_USERS DatabaseCollection = "users"
)

func (c DatabaseCollection) String() string {
	return string(c)
}
