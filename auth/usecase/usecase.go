package authUsecase

import (
	"chatapp/auth"
	"chatapp/models"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log"
	"time"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthUseCase struct {
	userRepos  auth.UserRepository
	hashSalt   string
	signingKey string
	expireTime time.Duration
}

func NewAuthUseCase(
	user_repos auth.UserRepository,
	hash_salt string,
	signing_key string,
	tokenTLSSeconds time.Duration) *AuthUseCase {

	return &AuthUseCase{
		userRepos:  user_repos,
		hashSalt:   hash_salt,
		signingKey: signing_key,
		expireTime: tokenTLSSeconds,
	}
}

func (a *AuthUseCase) hashPasword(p string) string {
	pwd := sha256.New()
	pwd.Write([]byte(p))
	pwd.Write([]byte(a.hashSalt))
	hashPassword := fmt.Sprintf("%x", pwd.Sum(nil))
	return hashPassword
}

func (a *AuthUseCase) SignUp(username, email, password string) error {

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	user := &models.User{
		Id:       id.String(),
		Email:    email,
		Username: username,
		Password: a.hashPasword(password),
	}

	return a.userRepos.CreateUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	User_id string `json:"user_id"`
}

func (a *AuthUseCase) GenerateAuthToken(email string, password string) (string, error) {
	password = a.hashPasword(password)
	user, err := a.userRepos.GetUser(email, password)
	if err != nil {
		log.Println(err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.expireTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	key, err := base64.URLEncoding.DecodeString(a.signingKey)
	if err != nil {
		log.Println(err)
	}
	signed_str, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
		return "", err
	}


	return signed_str, err
}

func (a *AuthUseCase) ParseToken(accessToken string) (string, error) {


	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		key, err := base64.URLEncoding.DecodeString(a.signingKey)
		if err != nil {
			log.Println(err)
		}

		return key, nil
	})

	if err != nil {
		log.Println("Parse token error: ", err)
		return "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.User_id, nil
	} else {
		return "", errors.New("Error invalid access token")
	}

}
