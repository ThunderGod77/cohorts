package students

import (
	"cohort/db"
	"cohort/global"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Student struct {
	Username string `json:"username" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func AddStudents(c *fiber.Ctx) error {
	s := new(Student)
	//to get body of the post request
	if err := c.BodyParser(s); err != nil {
		return fiber.NewError(400, err.Error())
	}

	//validating inputs
	v := validator.New()

	//validation errors
	err := v.Struct(s)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// to check if the user already exists
	present, err := db.GetStudent(s.Email, s.Username)
	// in case of error getting data from the database
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// in case user already exists
	if present {
		return fiber.NewError(fiber.StatusOK, "User with same username or email already exists!")
	}

	// to hash the password
	hash, err := hashPassword(s.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error!")
	}
	//to add student to the database
	err = db.AddStudent(s.Email, s.Username, hash)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	//logging about the creation of new users
	global.InfoLogger.Println("User with name " + s.Username + " has been created!")
	//response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"err": false, "msg": "User created successfully!"})

}
