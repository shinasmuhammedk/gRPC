package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("mysecretkey")


//Generate Token
func GenerateToken(userId int32) (string, error) {
	claims := jwt.MapClaims{
        "user_id":userId,
        "exp":time.Now().Add(time.Hour * 24).Unix(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    return token.SignedString(secretKey)
}

//Validate Token
func ValidateToken(tokenStr string)(*jwt.Token, error){
    return jwt.Parse(tokenStr, func (token *jwt.Token)(interface{},error){
        return secretKey,nil
    })
}