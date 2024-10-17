package utils

import (
    "os"
    "time"

    "github.com/golang-jwt/jwt/v4"
    "github.com/joho/godotenv"
)

type JWTClaims struct {
    ID             string `json:"id"`
    GithubUsername string `json:"github_username"`
    jwt.RegisteredClaims
}

func GenerateJWT(githubUsername, id string) (string, error) {
    _ = godotenv.Load()
    secret := os.Getenv("JWT_SECRET")

    claims := JWTClaims{
        ID:             id,
        GithubUsername: githubUsername,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    signedToken, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return signedToken, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, *JWTClaims, error) {
    _ = godotenv.Load()
    secret := os.Getenv("JWT_SECRET")

    claims := &JWTClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    return token, claims, err
}