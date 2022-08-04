package model

func ReadPublicTrains(id int) PublicTraining {
	db := Connection()
	var pt PublicTraining
	_ = db.Find(&pt)
	return pt
}
