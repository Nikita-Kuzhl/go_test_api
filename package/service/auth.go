package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	todo "github.com/Nikita-Kuzhl/go-rest-api"
	"github.com/Nikita-Kuzhl/go-rest-api/package/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "asldk3m24mkllgm13mormo2fmeomdowmav23mlfwe"
	signingKey= "kasdjkl31kjd3k1lcm^&£"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}


func (s *AuthService) CreateUser(user todo.User) (int,error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func  generatePasswordHash(password string) string {
	hash:= sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
func (s*AuthService) GenerateToken(login,password string) (string,error){
	user,err:= s.repo.GetUser(login,generatePasswordHash(password))
	if err!=nil{
		return "",err
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodES256,&tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		} ,user.Id,
	})
	return token.SignedString([]byte(signingKey))
}