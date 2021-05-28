package tests

import (
	"go-sa/database"
)

var client database.DataEntreprise

func init() {
	// 	conf := config.MongoConfiguration{
	// 		Server:     "mongodb://localhost:27017",
	// 		Database:   "Mgo",
	// 		Collection: "DatasTest",
	// 	}
	// 	ctx := context.TODO()

	// 	db := database.ConnectDB(ctx, conf)
	// 	collection := db.Collection(conf.Collection)

	// 	client = &database.DataEntreprise{
	// 		Col: collection,
	// 		Ctx: ctx,
	// 	}
	// }

	// func TestInsertData(t *testing.T) {
	// 	tests := map[string]struct {
	// 		payload      string
	// 		expectedCode int
	// 		expected     string
	// 	}{
	// 		"should return 200": {
	// 			payload: `{		"entrepriseId": 2,
	// 			"entreprise": "Apple",
	// 			"dirigeant": "Tim Cook",
	// 			"division_A": "Oui",
	// 			"CA_2017": 17,
	// 			"CA_2018": 18,
	// 			"CA_2019": 19,
	// 			"CA_2020": 20,
	// 			"SFEI_2017": 1,
	// 			"SFEI_2018": 2,
	// 			"SFEI_2019": 3,
	// 			"SFEI_2020": 4,
	// 			"SFIP_2017": 3,
	// 			"SFIP_2018": 2,
	// 			"SFIP_2019": 1,
	// 			"SFIP_2020": 5,
	// 			"Div_2017-20": "Non",
	// 			"PER_2017": 10,
	// 			"PER_2018": 10,
	// 			"PER_2019": 15,
	// 			"PER_2020": 15,
	// 			"PER_M": 12.5,
	// 			"PBR_2017": 1,
	// 			"PBR_2018": 2,
	// 			"PBR_2019": 3,
	// 			"PBR_2020": 4,
	// 			"PBR_M": 2,
	// 			"II": 0.9}`,
	// 			expectedCode: 200,
	// 			expected:     "Apple",
	// 		},
	// 		"should return 400": {
	// 			payload:      "invalid json string",
	// 			expectedCode: 400,
	// 		},
	// 	}

	// 	for name, test := range tests {
	// 		t.Run(name, func(t *testing.T) {
	// 			req, _ := http.NewRequest("POST", "/datas", strings.NewReader(test.payload))
	// 			rec := httptest.NewRecorder()
	// 			h := http.HandlerFunc(handlers.InsertData(client))
	// 			h.ServeHTTP(rec, req)
	// 			fmt.Println(rec.Body.String())
	// 			if test.expectedCode == 200 {
	// 				data := models.Data{}
	// 				_ = json.Unmarshal([]byte(rec.Body.String()), &data)
	// 				assert.Equal(t, test.expected, data.Entreprise)
	// 				assert.NotNil(t, data.ID)
	// 				//cleanup
	// 				_, _ = client.Delete(data.ID.(string))
	// 			}
	// 		})
	// 	}
}
