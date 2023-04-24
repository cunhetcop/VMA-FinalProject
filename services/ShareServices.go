// services/Share.go
package services

import (
	"fmt"
	"nguyenhalinh/go/database"
	"nguyenhalinh/go/models"
	"nguyenhalinh/go/utils"
	"strings"
	"time"
)

func RevokeToken(token string, expiration time.Duration) error {
	err := utils.RedisClient.Set(utils.RedisClient.Context(), token, "true", expiration).Err()
	return err
}
type ForgotPasswordInput struct {
	Email string `json:"email" validate:"required,email"`
}

func ResetPassword(input *ForgotPasswordInput) (string, error) {
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return "", fmt.Errorf("email not found")
	}

	newPassword := utils.GenerateRandomPassword(10)

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}

	user.Password = hashedPassword
	if err := database.DB.Save(&user).Error; err != nil {
		return "", fmt.Errorf("error saving user: %v", err)
	}

	emailSubject := "Your password has been changed"

	emailTemplate, err := utils.GetEmailTemplate()
	if err != nil {
		return "", fmt.Errorf("error reading email template: %v", err)
	}

	htmlBody := strings.Replace(emailTemplate, "{{.Password}}", newPassword, 1)

	go func() {
		err := utils.SendEmail(user.Email, emailSubject, htmlBody)
		if err != nil {
			fmt.Printf("error sending email: %v\n", err)
		}
	}()

	return newPassword, nil
}
