package students

import (
	"cohort/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const jwtSecret = "supercomputersSecretKey"

type LoginDetails struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// to match password and hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *fiber.Ctx) error {
	l := new(LoginDetails)
	//to parse request body
	if err := c.BodyParser(&l); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error parsing request!")
	}
	//to validate login fields
	v := validator.New()
	if err := v.Struct(l); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//to get student with this username or email
	sD, ok, err := db.GetStudent(l.Username, l.Username)
	if err != nil {
		//in case of an err
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if !ok {
		// in case user does not exist
		return fiber.NewError(fiber.StatusOK, "Wrong username or email!")
	}
	//to verify the password
	res := checkPasswordHash(l.Password, (*sD).Passw)
	if res != true {
		return fiber.NewError(fiber.StatusUnauthorized, "Incorrect Password!")
	}
	//creating a jwt token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = (*sD).ID
	claims["auth"] = (*sD).Auth
	//setting expiry time as 30 days
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30)

	//signing the jwt token
	s, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error signing token!")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user": struct {
			Id       string `json:"id"`
			Username string `json:"username"`
			Auth     string `json:"auth"`
		}{
			Id:       (*sD).ID,
			Username: l.Username,
			Auth:     (*sD).Auth,
		},
	})

}
