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

//id, is_costom, userIdからトレーニング名、消費カロリーを算出
func GetNameConsumptingC(id int, is_custome bool) (baseCalorie int) {
	// publicTrainingsから抽出
	db := ConnectionByTCP()
	var ut UserTraining
	var pt PublicTraining
	if is_custome {
		_ = db.Debug().Where("id = ?", id).Find(&ut)
		return ut.Calorie
	}
	_ = db.Debug().Where("id = ?", id).Find(&pt)
	return pt.Mets
	//log.Println(&pt)
	//log.Println(pt.Mets)

	//db.Where("name = ? AND age = ?", "jinzhu", "22").Find(&)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
}
