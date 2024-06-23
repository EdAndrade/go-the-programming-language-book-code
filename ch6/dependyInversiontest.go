package main

import "fmt"

func main() {

}

type User struct {
	id     int
	name   string
	emails int
}

func (user User) sendEmail() {
	println("Email sent...")
	user.emails++
}

type RetrieveUserByIdRepository interface {
	perform(userId int) User
}

type RetrieveUserByIdUseCase interface {
	perform(userId int) User
}

func RetrieveUserByIdUseCaseImpl(id int, retrieveUserByIdRepository RetrieveUserByIdRepository) User {
	user := retrieveUserByIdRepository.perform(id)
	user.sendEmail()
	return user
}

func RetrieveUserByIdController(retrieveUserByIdUseCase RetrieveUserByIdUseCase) {
	id := 1
	user := retrieveUserByIdUseCase.perform(id)
	fmt.Println(user)
}
