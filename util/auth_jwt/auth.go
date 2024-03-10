package auth_jwt

import (
	"errors"
	"strings"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type MapClaims struct {
	AccessToken  bool
	RefreshToken bool
	Token        string
	ExpUnix      int64
	IatUnix      int64
	Sub          string
}

type RespJWT struct {
	AccessToken           string
	AccessTokenExpiredAt  time.Time
	RefreshToken          string
	RefreshTokenExpiredAt time.Time
}

type encodeDecodeTokenJWT struct {
	algorithm              *jwtv5.SigningMethodHMAC
	accessTokenExpiration  time.Duration
	refreshTokenExpiration time.Duration
	secret                 string
}

var encodeDecode *encodeDecodeTokenJWT

func GetJWT() *encodeDecodeTokenJWT {
	if encodeDecode == nil {
		encodeDecode = &encodeDecodeTokenJWT{
			algorithm:              jwtv5.SigningMethodHS256,
			accessTokenExpiration:  time.Duration(viper.GetInt("jwt_access_token_duration")) * time.Hour,
			refreshTokenExpiration: time.Duration(viper.GetInt("jwt_refresh_token_duration")) * time.Hour,
			secret:                 viper.GetString("jwt_secret"),
		}
	}
	return encodeDecode
}

func (e encodeDecodeTokenJWT) GetToken(email string) (*RespJWT, error) {
	userClaims := jwtv5.MapClaims{
		"sub": email,
	}
	accessToken, err := e.encodeToken("access_token", userClaims)
	if err != nil {
		return nil, err
	}
	refreshToken, err := e.encodeToken("refresh_token", userClaims)
	if err != nil {
		return nil, err
	}
	return &RespJWT{
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  time.Now().UTC().Add(e.accessTokenExpiration),
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: time.Now().UTC().Add(e.refreshTokenExpiration),
	}, nil
}

func (e encodeDecodeTokenJWT) GetClaims(auth string) (*MapClaims, error) {
	token, err := e.decodedToken(auth)
	if err != nil || token == nil {
		return nil, err
	}
	mapClaims := MapClaims{
		Token: token.Raw,
	}
	if claims, ok := token.Claims.(jwtv5.MapClaims); ok && token.Valid {
		if sub, ok := claims["sub"].(string); ok {
			mapClaims.Sub = sub
		}
		if exp, ok := claims["exp"].(float64); ok {
			mapClaims.ExpUnix = int64(exp)
		}
		if iat, ok := claims["iat"].(float64); ok {
			mapClaims.IatUnix = int64(iat)
		}
		if accToken, ok := claims["access_token"].(bool); ok {
			mapClaims.AccessToken = accToken
		}
		if refreshToken, ok := claims["refresh_token"].(bool); ok {
			mapClaims.RefreshToken = refreshToken
		}
	}
	return &mapClaims, nil
}

func (e encodeDecodeTokenJWT) decodedToken(auth string) (*jwtv5.Token, error) {
	if auth == "" {
		return nil, errors.New("authorization header is missing")
	}

	parts := strings.Split(auth, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid authorization header format")
	}

	tokenString := parts[1]
	token, err := jwtv5.Parse(tokenString, func(token *jwtv5.Token) (interface{}, error) {
		if token.Method != jwtv5.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(e.secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (e encodeDecodeTokenJWT) encodeToken(tokenType string, userClaims jwtv5.MapClaims) (string, error) {
	exp := float64(time.Now().UTC().Add(e.accessTokenExpiration).Unix())
	iat := float64(time.Now().UTC().Unix())
	jwtClaims := jwtv5.MapClaims{
		"exp": exp,
		"iat": iat,
	}
	if tokenType == "access_token" {
		jwtClaims["access_token"] = true
	} else if tokenType == "refresh_token" {
		jwtClaims["refresh_token"] = true
	}
	for key, value := range userClaims {
		jwtClaims[key] = value
	}
	token := jwtv5.NewWithClaims(e.algorithm, jwtClaims)
	secret := []byte(e.secret)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
