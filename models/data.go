package models

type Data struct {
	EntrepriseID int         `json:"entrepriseId" bson:"entrepriseId"`
	ID           interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Entreprise   string      `json:"entreprise,omitempty" bson:"entreprise,omitempty"`
	Dirigeant    string      `json:"dirigeant,omitempty" bson:"dirigeant,omitempty"`
	Division_A   string      `json:"division_A,omitempty" bson:"Division_A,omitempty"`
	CA_2017      float32     `json:"CA_2017,omitempty" bson:"CA_2017,omitempty"`
	CA_2018      float32     `json:"CA_2018,omitempty" bson:"CA_2018,omitempty"`
	CA_2019      float32     `json:"CA_2019,omitempty" bson:"CA_2019,omitempty"`
	CA_2020      float32     `json:"CA_2020,omitempty" bson:"CA_2020,omitempty"`
	SFEI_2017    float32     `json:"SFEI_2017,omitempty" bson:"SFEI_2017,omitempty"`
	SFEI_2018    float32     `json:"SFEI_2018,omitempty" bson:"SFEI_2018,omitempty"`
	SFEI_2019    float32     `json:"SFEI_2019,omitempty" bson:"SFEI_2019,omitempty"`
	SFEI_2020    float32     `json:"SFEI_2020,omitempty" bson:"SFEI_2020,omitempty"`
	SFIP_2017    float32     `json:"SFIP_2017 ,omitempty" bson:"SFIP_2017,omitempty"`
	SFIP_2018    float32     `json:"SFIP_2018,omitempty" bson:"SFIP_2018,omitempty"`
	SFIP_2019    float32     `json:"SFIP_2019,omitempty" bson:"SFIP_2019,omitempty"`
	SFIP_2020    float32     `json:"SFIP_2020,omitempty" bson:"SFIP_2020,omitempty"`
	Div_2017_20  string      `json:"Div_2017_20,omitempty" bson:"Div_2017_20,omitempty"`
	PER_2017     float32     `json:"PER_2017,omitempty" bson:"PER_2017,omitempty"`
	PER_2018     float32     `json:"PER_2018,omitempty" bson:"PER_2018,omitempty"`
	PER_2019     float32     `json:"PER_2019,omitempty" bson:"PER_2019,omitempty"`
	PER_2020     float32     `json:"PER_2020,omitempty" bson:"PER_2020,omitempty"`
	PER_M        float32     `json:"PER_M,omitempty" bson:"PER_M,omitempty"`
	PBR_2017     float32     `json:"PBR_2017,omitempty" bson:"PBR_2017,omitempty"`
	PBR_2018     float32     `json:"PBR_2018,omitempty" bson:"PBR_2018,omitempty"`
	PBR_2019     float32     `json:"PBR_2019,omitempty" bson:"PBR_2019,omitempty"`
	PBR_2020     float32     `json:"PBR_2020,omitempty" bson:"PBR_2020,omitempty"`
	PBR_M        float32     `json:"PBR_M,omitempty" bson:"PBR_M,omitempty"`
	II           float32     `json:"II,omitempty" bson:"II,omitempty"`
}

type DataUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Data
}

type DataDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
