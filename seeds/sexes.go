package seeds

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/riada/models"
	"io/ioutil"
	"os"

	"gorm.io/gorm"
)

type Sex struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

type Sexes struct {
	Sexes []Sex `json:"sexes"`
}

func SexResponse(sex Sex) models.Sex {
	return models.Sex{Name: sex.Name, ShortName: sex.ShortName}
}

func SetSexes(db *gorm.DB) {
	var dbSex models.Sex
	// db := connection.Db()
	if err := db.First(&dbSex).Error; errors.Is(err, gorm.ErrRecordNotFound) {

		jsonFile, err := os.Open("./seeds/json/sexes.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, err := ioutil.ReadAll(jsonFile)

		var jsexes Sexes
		var nsexes []models.Sex
		json.Unmarshal(byteValue, &jsexes)
		for _, jsex := range jsexes.Sexes {
			nsexes = append(nsexes, SexResponse(jsex))
			// fmt.Println(jsex)
		}
		// fmt.Println(nsexes)
		if err := db.Create(&nsexes); err != nil {
			fmt.Println(err)
		}
		fmt.Println("created sexes successfully")

	}

	fmt.Println(dbSex)
}
