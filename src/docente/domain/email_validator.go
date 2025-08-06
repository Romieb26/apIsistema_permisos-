package domain


type IEmailValidator interface {
	Validate(email string) (bool, error)
}