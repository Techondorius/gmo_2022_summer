package model

func ReadPublicTrainigs() []PublicTraining {
	db := ConnectionByTCP()
	var pt []PublicTraining
	_ = db.Find(&pt)
	return pt
}
