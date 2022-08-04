package model

//ユーザー登録
func CreateUser(u User) error {
	db := ConnectionByTCP()
	//db.Create(&product) // pass pointer of data to Create
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

//ユーザー情報変更
func UpdateUser(u User) error {
	db := ConnectionByTCP()

	result := db.Save(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

//ユーザー情報表示
func GetUser(id string) User {
	db := ConnectionByTCP()
	var u User
	_ = db.Where("id = ?", id).Find(&u)
	return u
}
