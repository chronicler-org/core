package authService

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	attendantService "github.com/chronicler-org/core/src/attendant/service"
	authDTO "github.com/chronicler-org/core/src/auth/dto"
	appEnum "github.com/chronicler-org/core/src/auth/enum"
	authInterface "github.com/chronicler-org/core/src/auth/interface"
	authExceptionMessage "github.com/chronicler-org/core/src/auth/messages"
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

func (service *AuthService) Login(dto authDTO.AuthLoginDTO) (authDTO.ResponseAuthDTO, error) {
	user, userType, err := service.validateUserByEmail(dto)
	if err != nil {
		return authDTO.ResponseAuthDTO{}, err
	}

	authToken, expirationTime, err := service.generateAT(user.GetID().String(), userType)

	return authDTO.ResponseAuthDTO{
		AuthToken: authToken,
		ExpiresIn: expirationTime,
		User:      user,
	}, err
}

func (service *AuthService) validateUserByEmail(dto authDTO.AuthLoginDTO) (authInterface.IUser, string, error) {
	email := dto.Email

	hashedPassword := ""
	userType := ""
	var user authInterface.IUser

	manager, err := service.managerService.FindManagerByEmail(email)
	if err == nil {
		hashedPassword = manager.Password
		userType = appEnum.ManagerRole
		user = manager
	} else {
		attendant, err := service.attendantService.FindAttendantByEmail(email)
		if err == nil {
			hashedPassword = attendant.Password
			userType = appEnum.AttendantRole
			user = attendant
		}
	}

	matchPass, err := service.validatePassword(dto.Password, hashedPassword)
	if err != nil {
		return nil, "", err
	}
	if !matchPass {
		return nil, "", appException.UnauthorizedException(authExceptionMessage.LOGIN_FAILED)
	}

	return user, userType, nil
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
	authSecret := []byte(os.Getenv("AT_SECRET"))
	authTokenExpiresIn := os.Getenv("AT_EXPIRES_IN")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	expirationHours, err := strconv.Atoi(authTokenExpiresIn)
	if err != nil {
		return "", time.Time{}, err
	}

	expirationTime := time.Now().Add(time.Duration(expirationHours) * time.Hour)
	token.Claims.(jwt.MapClaims)["exp"] = expirationTime.Unix()

	authToken, err := token.SignedString(authSecret)
	if err != nil {
		return "", time.Time{}, err
	}

	return authToken, expirationTime, nil
}

func (service *AuthService) validatePassword(password string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}
