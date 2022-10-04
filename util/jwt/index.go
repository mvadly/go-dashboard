package jwt

import (
	"context"
	"errors"
	"fmt"
	"go-dashboard/config"
	"go-dashboard/util"
	"go-dashboard/v1/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	jwk "github.com/lestrrat-go/jwx/jwk"
)

type UserSession struct {
	ID       uuid.UUID `json:"_id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
}

type SessionClaims struct {
	UserSession
	jwt.StandardClaims
}

func SetToken(user models.Users) (string, error) {
	var session = UserSession{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
	}

	claims := &SessionClaims{
		session,
		jwt.StandardClaims{
			ExpiresAt: util.TimeNow().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(config.LoadEnv("JWT_KEY").(string)))

}

func getKey(token *jwt.Token) (interface{}, error) {

	fmt.Println("token disini:", token)

	// For a demonstration purpose, Google Sign-in is used.
	// https://developers.google.com/identity/sign-in/web/backend-auth
	//
	// This user-defined KeyFunc verifies tokens issued by Google Sign-In.
	//
	// Note: In this example, it downloads the keyset every time the restricted route is accessed.
	keySet, err := jwk.Fetch(context.Background(), "https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		return nil, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
	}

	key, found := keySet.LookupKeyID(keyID)

	if !found {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}

	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}

	return pubkey, nil
}

func JwtConfig() middleware.JWTConfig {
	signingKey := config.LoadEnv("JWT_KEY").(string)
	return middleware.JWTConfig{
		Claims:     &SessionClaims{},
		SigningKey: []byte(signingKey),

		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			code := http.StatusInternalServerError
			he, ok := err.(*echo.HTTPError)
			if !ok {
				return util.JSON(c, code, util.ResJSON{
					Code:    "01",
					Message: err.Error(),
					Data:    nil,
				})
			}

			return util.JSON(c, he.Code, util.ResJSON{
				Code:    "01",
				Message: he.Message.(string),
				Data:    nil,
			})
		},
		// KeyFunc: getKey,
	}
}
