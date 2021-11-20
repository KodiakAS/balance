package splitter

import (
	"reflect"
	"testing"

	types "github.com/KodiakAS/balance/types"
)

func Test_getMinBucket(t *testing.T) {
	type args struct {
		items types.BucketList
	}
	tests := []struct {
		name string
		args args
		want *types.Bucket
	}{
		{
			name: "normal",
			args: args{types.BucketList{
				types.Bucket{Items: types.ItemList{{Data: "No1", Factor: 3}}, Factor: 3},
				types.Bucket{Items: types.ItemList{{Data: "No2", Factor: 31}}, Factor: 31},
				types.Bucket{Items: types.ItemList{{Data: "No3", Factor: 0}}, Factor: 0},
			}},
			want: &types.Bucket{Items: types.ItemList{{Data: "No3", Factor: 0}}, Factor: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinBucket(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMinBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	type args struct {
		items      types.ItemList
		bucket_num int
	}
	tests := []struct {
		name    string
		args    args
		want    types.BucketList
		wantErr bool
	}{
		{
			name: "normal",
			args: args{types.ItemList{
				{Data: "1", Factor: 1},
				{Data: "2", Factor: 2},
				{Data: "3", Factor: 4},
				{Data: "4", Factor: 10},
			}, 3},
			want: types.BucketList{
				types.Bucket{Items: types.ItemList{{Data: "2", Factor: 2}, {Data: "1", Factor: 1}}, Factor: 3},
				types.Bucket{Items: types.ItemList{{Data: "3", Factor: 4}}, Factor: 4},
				types.Bucket{Items: types.ItemList{{Data: "4", Factor: 10}}, Factor: 10},
			},
			wantErr: false,
		},
		{
			name: "has_zeros",
			args: args{types.ItemList{
				{Data: "1", Factor: 1},
				{Data: "2", Factor: 0},
				{Data: "3", Factor: 0},
				{Data: "4", Factor: 0},
				{Data: "5", Factor: 0},
				{Data: "6", Factor: 0},
			}, 3},
			want: types.BucketList{
				types.Bucket{Items: types.ItemList{{Data: "1", Factor: 1}, {Data: "2", Factor: 0}, {Data: "5", Factor: 0}}, Factor: 1},
				types.Bucket{Items: types.ItemList{{Data: "3", Factor: 0}, {Data: "6", Factor: 0}}, Factor: 0},
				types.Bucket{Items: types.ItemList{{Data: "4", Factor: 0}}, Factor: 0},
			},
			wantErr: false,
		},
		{
			name:    "empty_itemlist",
			args:    args{types.ItemList{}, 1},
			want:    types.BucketList{types.Bucket{}},
			wantErr: false,
		},
		{
			name:    "zero_buckets",
			args:    args{types.ItemList{}, 0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Split(tt.args.items, tt.args.bucket_num)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
