//database/seeder/createDefaultRoles.go
package seeder

import (
	"nguyenhalinh/go/models"
	"nguyenhalinh/go/database"
)

func CreateDefaultRoles() (userRole models.Role, adminRole models.Role, err error) {

	err = database.DB.Where("name = ?", "user").FirstOrCreate(&userRole, models.Role{Name: "user"}).Error
	if err != nil {
		return models.Role{}, models.Role{}, err
	}

	err = database.DB.Where("name = ?", "admin").FirstOrCreate(&adminRole, models.Role{Name: "admin"}).Error
	if err != nil {
		return models.Role{}, models.Role{}, err
	}

	return userRole, adminRole, nil
}
