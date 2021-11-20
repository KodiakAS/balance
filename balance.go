package balance

import (
	jp "github.com/KodiakAS/balance/jsonparser"
	is "github.com/KodiakAS/balance/splitter"
	types "github.com/KodiakAS/balance/types"
)

func BalanceItemsFromJson(path string, chunks int) (types.BucketList, error) {
	items, err := jp.ParseJsonFile(path)
	if err != nil {
		return nil, err
	}
	buckets, err := is.Split(items.Items, chunks)
	if err != nil {
		return nil, err
	}
	return buckets, nil
}

func BalanceItems(items types.ItemList, chunks int) (types.BucketList, error) {
	buckets, err := is.Split(items, chunks)
	if err != nil {
		return nil, err
	}
	return buckets, nil
}
