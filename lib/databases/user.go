package databases

import (
	"task-management/config"
	"task-management/middlewares"
	"task-management/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(newUser *models.User) (interface{}, error) {
	// insert value of newUser into database
	tx := config.DB.Create(&newUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return newUser, nil
}

func LoginUser(dataLogin *models.LoginUser) (interface{}, error) {
	var user models.User
	// Check if user has email of login data is exist in database
	tx := config.DB.Where("email = ?", dataLogin.Email).First(&user)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	// Check if inputed password is match to password in database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dataLogin.Password))
	if err != nil {
		return nil, err
	}

	// Generate token by using user's ID
	user.Token, _ = middlewares.CreateToken(int(user.ID))

	// Update value of user token into database
	tx2 := config.DB.Save(&user)
	if tx2.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	return user.Token, nil
}

func GetUserByID(id int) (interface{}, error) {
	var user models.User
	tx := config.DB.First(&user, id)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	var getUser models.GetUser
	getUser.ID = user.ID
	getUser.Name = user.Name
	getUser.Email = user.Email

	return getUser, nil
}

func UpdateUser(id int, newData *models.UpdateDataUser) (interface{}, error) {
	var user models.User
	// Check if user has email of update data is exist in database
	tx := config.DB.First(&user, id)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	var updateData models.User

	if len(newData.Name) == 0 {
		updateData.Name = user.Name
	} else {
		updateData.Name = newData.Name
	}

	if len(newData.Email) == 0 {
		updateData.Email = user.Email
	} else {
		updateData.Email = newData.Email
	}

	if len(newData.Password) == 0 {
		updateData.Password = user.Password
	} else {
		// Encrypt user's password before insert into database
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(newData.Password), bcrypt.DefaultCost)
		newData.Password = string(hashPassword)
		updateData.Password = newData.Password
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newData.CurrentPassword))
	if err != nil {
		return nil, err
	}

	tx2 := config.DB.Where("id = ?", id).Updates(&updateData)
	if tx2.Error != nil || tx.RowsAffected == 0 {
		return nil, tx2.Error
	}

	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user *models.User
	tx := config.DB.Delete(&user, id)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	return user, tx.Error
}

func GetUserByEmail(email string) (interface{}, int) {
	var user models.User
	tx := config.DB.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return tx.Error, 0
	}

	return user, 1
}
