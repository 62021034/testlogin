package controllers

import (
	"database/sql"
	"log"
	"regexp"
	"testlogin/database"
	"testlogin/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt" //เข้ารหัส
)

func GetUsers(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT * FROM users")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Userses{}

	for rows.Next() {
		users := models.Users{}
		if err := rows.Scan(&users.UsersId, &users.LoginEmail, &users.LoginPassword, &users.UsersName, &users.UserSurname); err != nil {
			return err
		}
		result.Userses = append(result.Userses, users)
	}
	return c.JSON(result)
}

func Register(c *fiber.Ctx) error {
	var datausers map[string]string

	if err := c.BodyParser(&datausers); err != nil {
		return err
	}
	//เข้ารหัส
	// enCryPassword, _ := bcrypt.GenerateFromPassword([]byte(datausers["password"]), 14)
	hash, _ := HashPassword(datausers["password"])

	users := &models.Users{
		LoginEmail:    datausers["email"],
		LoginPassword: hash,
		UsersName:     datausers["name"],
		UserSurname:   datausers["lastname"],
	}

	// ตรวจสอบอีเมล
	r, err := regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]*@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		log.Fatal(err)
	}

	if !r.MatchString(datausers["email"]) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid email address")
	}

	// ตรวจสอบชื่อ
	if datausers["name"] == "" || datausers["lastname"] == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Name and surname must not be empty")
	}

	sqlstmt := `
	INSERT INTO users (login_email,login_password,user_name,user_surname)
	VALUES ($1, $2, $3, $4)`

	_, err = database.DB.Exec(sqlstmt, users.LoginEmail, users.LoginPassword, users.UsersName, users.UserSurname)
	if err != nil {
		log.Panic(err)
	}

	return c.JSON(users)
}

func Login(c *fiber.Ctx) error {
	var datauserslogin map[string]string
	if err := c.BodyParser(&datauserslogin); err != nil {
		return err
	}

	var loginUsers models.Users

	rows := database.DB.QueryRow("SELECT login_email, login_password FROM users WHERE login_email = $1", datauserslogin["email"])

	if err := rows.Scan(&loginUsers.LoginEmail, &loginUsers.LoginPassword); err != nil {
		if err == sql.ErrNoRows {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "user not found",
			})
		}
		return err
	}

	matchs := CheckPasswordHash(loginUsers.LoginPassword, datauserslogin["password"])

	// if err := bcrypt.CompareHashAndPassword([]byte(loginUsers.LoginPassword), []byte(datauserslogin["password"])); err == nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON([]byte(datauserslogin["password"]))
	// }

	// fiber.Map{
	// 	"message": "success",
	// }
	return c.JSON(matchs)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
