package jsonparser

import (
	"encoding/json"
	"os"

	types "github.com/KodiakAS/balance/types"
)

func ParseJsonFile(path string) (*types.Items, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return &types.Items{}, err
	}

	var allItems types.Items
	if err = json.Unmarshal(bytes, &allItems); err != nil {
		return &types.Items{}, err
	}
	return &allItems, nil
}
