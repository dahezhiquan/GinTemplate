package jwts

import (
	"GinTemplate/common/errs"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtToken struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}

func CreateToken(val, AcSecret, rfSecret string, exp, rfExp time.Duration) (*JwtToken, error) {
	aExp := time.Now().Add(exp).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"memIdSec": val,
		"exp":      aExp,
	})
	token, err := claims.SignedString([]byte(AcSecret))
	if err != nil {
		return nil, err
	}

	rExp := time.Now().Add(rfExp).Unix()
	refreshClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"memIdSec": val,
		"exp":      rExp,
	})
	refreshToken, err := refreshClaims.SignedString([]byte(rfSecret))
	if err != nil {
		return nil, err
	}
	return &JwtToken{
		AccessToken:  token,
		RefreshToken: refreshToken,
		AccessExp:    aExp,
		RefreshExp:   rExp,
	}, nil
}

func ParseToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		memIdSec := claims["memIdSec"].(string)
		exp := int64(claims["exp"].(float64))
		if exp <= time.Now().Unix() {
			return "", errs.LoginExpirationError
		}
		return memIdSec, nil
	} else {
		return "", err
	}
}
