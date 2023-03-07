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

type TypeDoc struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type TypeDocs struct {
	TypeDocs []TypeDoc `json:"type_docs"`
}

func TypeDocResponse(typedoc TypeDoc) models.TypeDoc {
	return models.TypeDoc{Name: typedoc.Name, Desc: typedoc.Desc}
}

func SetTypeDocs(db *gorm.DB) {
	var dbTypeDoc models.TypeDoc
	if err := db.First(&dbTypeDoc).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		jsonFile, err := os.Open("./seeds/json/typedocs.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)

		var typedocs TypeDocs
		var nTypeDocs []models.TypeDoc
		json.Unmarshal(byteValue, &typedocs)
		for _, td := range typedocs.TypeDocs {
			nTypeDocs = append(nTypeDocs, TypeDocResponse(td))
		}
		if len(typedocs.TypeDocs) == len(nTypeDocs) {
			if err := db.Create(&nTypeDocs); err != nil {
				fmt.Println(err)
			}

			fmt.Println("type_docs successfully created")
		}

	}

}
