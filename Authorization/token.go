package authorization

import (
	model "backend_en_go/Model"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(data *model.Login) (string, error) {
	if signKey == nil {
		return "", fmt.Errorf("signKey is nil")
	}
	claim := model.Claim{
		Email:  data.Email,
		Rol:    data.Rol,
		UserID: data.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "Harold",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// validar token

func ValidateToken(t string) (model.Claim, error) {

	token, err := jwt.ParseWithClaims(t, &model.Claim{}, VerifyFunction)
	if err != nil {
		return model.Claim{}, err
	}
	if !token.Valid {
		return model.Claim{}, errors.New("token no valido")
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("no se pudo obtener los claim")

	}
	return *claim, nil
}

func VerifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
