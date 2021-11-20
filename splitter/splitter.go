package splitter

import (
	"errors"
	"sort"

	types "github.com/KodiakAS/balance/types"
)

func getMinBucket(items types.BucketList) *types.Bucket {
	sort.Sort(items)
	return &items[0]
}

func Split(items types.ItemList, bucket_num int) (types.BucketList, error) {
	if bucket_num <= 0 {
		return nil, errors.New("bucket number must be greater than 0")
	}
    
	var buckets types.BucketList
	var zerosGoto int
	for i := 0; i < bucket_num; i++ {
		buckets = append(buckets, types.Bucket{})
	}
	if len(items) == 0 {
		return buckets, nil
	}
	sort.Sort(sort.Reverse(items))
	for _, i := range items {
		if i.Factor == 0 {
			buckets[zerosGoto].AddItem(i)
			zerosGoto = (zerosGoto + 1) % bucket_num
		} else {
			minBucket := getMinBucket(buckets)
			minBucket.AddItem(i)
		}
	}
	return buckets, nil
}
