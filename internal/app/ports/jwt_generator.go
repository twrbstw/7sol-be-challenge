package ports

type IJwtGenerator interface {
	GenerateJwt(name, email string) (string, error)
}
