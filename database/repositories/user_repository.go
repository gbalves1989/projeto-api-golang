package repositories

import (
	"strings"

	"github.com/gbalves1989/projeto-api-golang/database"
	"github.com/gbalves1989/projeto-api-golang/database/models"
	"github.com/gbalves1989/projeto-api-golang/http/requests"
)

func UserShowByEmailRepository(email string) models.User {
	var user models.User
	database.DB.Where("email = ?", strings.ToLower(email)).First(&user)
	return user
}

func UserStoreRepository(input requests.AuthRegisterRequest, hash string) models.User {
	user := models.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hash,
	}

	database.DB.Create(&user)
	return user
}

func UserShowByIDRepository(userID uint) models.User {
	var user models.User
	database.DB.First(&user, userID)
	return user
}

func UserUpdateRepository(input requests.UserUpdateRequest, user models.User) models.User {
	user.Name = input.Name
	database.DB.Save(&user)
	return user
}

func UserUpdatePasswordRepository(hash string, user models.User) models.User {
	user.PasswordHash = hash
	database.DB.Save(&user)
	return user
}

func UserUploadRepository(filePath string, user models.User) models.User {
	user.AvatarURL = filePath
	database.DB.Save(&user)
	return user
}

func UserIndexRepository(pagination requests.PaginationRequest) []models.User {
	var users []models.User
	database.DB.Limit(pagination.Limit).Offset(pagination.Offset).Find(&users)
	return users
}

func CountUsersRepository() int64 {
	var total int64
	database.DB.Model(&models.User{}).Count(&total)
	return total
}

func UserDestroyRepository(user models.User) {
	database.DB.Delete(&user)
}
