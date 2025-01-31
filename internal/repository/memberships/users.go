package memberships

import (
	"github.com/Aziz8860/song-catalog/internal/models/memberships"
)

func (r *repository) CreateUser(model memberships.User) error {
	// karena pake Gorm, ga perlu mendefinisikan query
	// debug: fmt.Println(r.db)
	return r.db.Create(&model).Error
}

func (r *repository) GetUser(email, username string, id uint) (*memberships.User, error) {
	user := memberships.User{}
	// template kedua itu menggunakan AND
	res := r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
