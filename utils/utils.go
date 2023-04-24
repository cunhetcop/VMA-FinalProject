// utils/utils.go
package utils

import (
	"nguyenhalinh/go/database"
	"nguyenhalinh/go/models"
)


func GetRoleIDByName(roleName string) (uint, error) {
    var role models.Role
    err := database.DB.Where("name = ?", roleName).First(&role).Error
    if err != nil {
        return 0, err
    }
    return role.ID, nil
}
