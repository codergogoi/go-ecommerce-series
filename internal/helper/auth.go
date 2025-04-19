package helper

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-ecommerce-app/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashedPassword(p string) (string, error) {

	if len(p) < 6 {
		return "", errors.New("password length should be at least 6 characters long")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(p), 10)

	if err != nil {
		// log actual error and report to logging tool
		return "", errors.New("password hash failed")
	}

	return string(hashP), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {

	if id == 0 || email == "" || role == "" {
		return "", errors.New("required inputs are missing to generate token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		return "", errors.New("unable to signed the token")
	}

	return tokenStr, nil
}

func (a Auth) VerifyPassword(pP string, hP string) error {

	if len(pP) < 6 {
		return errors.New("password length should be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP))

	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}

func (a Auth) VerifyToken(t string) (domain.User, error) {
	tokenArr := strings.Split(t, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, nil
	}

	tokenStr := tokenArr[1]

	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)
		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {

	authHeader := ctx.GetReqHeaders()["Authorization"]
	user, err := a.VerifyToken(authHeader[0])

	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	}

}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)

}

func (a Auth) GenerateCode() (string, error) {
	return RandomNumbers(6)
}

func (a Auth) AuthorizeSeller(ctx *fiber.Ctx) error {

	authHeader := ctx.GetReqHeaders()["Authorization"]
	user, err := a.VerifyToken(authHeader[0])

	if err != nil {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	} else if user.ID > 0 && user.UserType == domain.SELLER {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  errors.New("please join seller program to manage products"),
		})
	}

}
