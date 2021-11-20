package jsonparser

import (
	"encoding/json"
	"io/ioutil"
	"os"

	types "github.com/KodiakAS/balance/types"
)

func ParseJsonFile(path string) (*types.Items, error) {
	var allItems types.Items
	jsonFile, err := os.Open(path)
	if err != nil {
		return &types.Items{}, err
	}
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &types.Items{}, err
	}
	err = json.Unmarshal(bytes, &allItems)
	if err != nil {
		return &types.Items{}, err
	}
	return &allItems, nil
}
