package model

func ReadPublicTrains(id int) PublicTraining {
	db := ConnectionByTCP()
	var pt PublicTraining
	_ = db.Find(&pt)
	return pt
}
