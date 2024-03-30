package authService

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	attendantService "github.com/chronicler-org/core/src/attendant/service"
	managerService "github.com/chronicler-org/core/src/manager/service"
)

type AuthService struct {
	managerService   *managerService.ManagerService
	attendantService *attendantService.AttendantService
}

func InitAuthService(
	managerService *managerService.ManagerService,
	attendantService *attendantService.AttendantService,
) *AuthService {
	return &AuthService{
		managerService:   managerService,
		attendantService: attendantService,
	}
}

func (service *AuthService) generateAT(id string, userType string) (string, time.Time, error) {
	jtiBytes := make([]byte, 12)
	if _, err := rand.Read(jtiBytes); err != nil {
		return "", time.Time{}, err
	}
	jti := hex.EncodeToString(jtiBytes)

	payload := jwt.MapClaims{
		"jti":  jti,
		"sub":  id,
		"role": userType,
	}

	authSecret := os.Getenv("AT_SECRET")
	authTokenExpiresIn := os.Getenv("AT_EXPIRES_IN")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	expirationHours, err := strconv.Atoi(authTokenExpiresIn)
	if err != nil {
		return "", time.Time{}, err
	}

	expirationTime := time.Now().Add(time.Duration(expirationHours) * time.Hour)
	token.Claims.(jwt.MapClaims)["exp"] = expirationTime.Unix()

	accessToken, err := token.SignedString(authSecret)
	if err != nil {
		return "", time.Time{}, err
	}

	return accessToken, expirationTime, nil
}
