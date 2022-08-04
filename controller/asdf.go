package controller

import "gmo_2022_summer/model"

type TRLIst struct {
	ID           int
	Name         string
	UserTR       bool
	ConsumptingC int
}

func asdf() {
	userID := "Pi"
	// public training 全取得
	pt := model.ReadPublicTrainigs()
	ut := model.ReadUserTrainings(userID)

}

func PTtoTRL(pt []model.PublicTraining) []TRLIst {
	var trl []TRLIst
	for i := 0; i < len(trl); i++ {
		tr := TRLIst{
			ID:           trl[i].ID,
			Name:         trl[i].Name,
			UserTR:       false,
			ConsumptingC: trl[i].ConsumptingC,
		}
		trl = append(trl, tr)
	}
}
