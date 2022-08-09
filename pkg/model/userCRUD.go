package model

// CreateUser ユーザー登録
func CreateUser(u User) error {
	db, _ := Connection()
	//db.Create(&product) // pass pointer of data to Create
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// UpdateUser ユーザー情報変更
func UpdateUser(u User) error {
	db, _ := Connection()

	result := db.Debug().Save(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// GetUser ユーザー情報表示
func GetUser(userid string) User {
	db, _ := Connection()
	var u User
	_ = db.Where("user_id = ?", userid).Find(&u)
	return u
}