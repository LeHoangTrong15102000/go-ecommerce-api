package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// user repo thì chúng ta sẽ viết tắt là ur
func (ur *UserRepo) GetInfoUser() string {

	return "Golang Course"

}