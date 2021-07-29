package entity

type Vaccine struct {
	ID           string `json:"id" bson:"_id"`
	VaccineName  string `json:"vaccine_name" bson:"vaccine_name"`
	VaccineLotNo string `json:"vaccine_lot_no" bson:"vaccine_lot_no"`
}
