package model

// "gorm.io/gorm"
import (
	"log"
	"time"
)

//トップページのトレーニング登録機能
func TrainingAdd(u TrainingHistory) error {
	db := ConnectionByTCP()
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

//ユーザー登録
func UserCreate(u User) error {
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
func UserUpdate(u User) error {
	db := ConnectionByTCP()

	result := db.Save(u)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

//ユーザー情報表示
func GetUser(id string) []User {
	db := ConnectionByTCP()
	var u []User
	_ = db.Where("id = ?", id).Find(&u)
	return u
}

func GetUserWeight(id string) []User {
	db := ConnectionByTCP()
	var u []User
	_ = db.Where("id = ?", id).Find(&u)
	//エラー↓
	//return u.
	return u
}

//user_idとtraining_dateを指定してTrainingHistoryから情報を抜き出す
func PeriodData(id string, start time.Time, stop time.Time) []TrainingHistory {
	db := ConnectionByTCP()
	var th []TrainingHistory
	//"UO"のところは認証情報からとってくる
	_ = db.Debug().Where("user_id = ? AND ? <= created_at <= ?", id, start, stop).Find(&th)
	return th
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

//カスタムトレーニング追加
func AddCustomeTR(u UserTraining) error {
	db := ConnectionByTCP()
	log.Println(u)
	result := db.Create(&u)
	log.Println(result)
	log.Println(result.Error)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

//カスタムトレーニング削除
func DeleteCustomeTR(u UserTraining) error {
	db := ConnectionByTCP()
	log.Println(u)
	//"10"のところにuser_id認証データ持ってくる
	result := db.Delete(&u, 3)
	log.Println(result)
	log.Println(result.Error)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ReadPublicTrainigs() []PublicTraining {
	db := ConnectionByTCP()
	var pt []PublicTraining
	_ = db.Find(&pt)
	return pt
}

func ReadUserTrainings(id string) []UserTraining {
	db := ConnectionByTCP()
	var ut []UserTraining
	_ = db.Where("user_id = ?", id).Find(&ut)
	return ut
}
