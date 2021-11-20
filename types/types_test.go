package types

import (
	"reflect"
	"testing"
)

func TestBucket_AddItem(t *testing.T) {
	type fields struct {
		Items  ItemList
		Factor int
	}
	type args struct {
		i Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "normal",
			fields: fields{},
			args:   args{Item{Data: "no1", Factor: 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bucket{
				Items:  tt.fields.Items,
				Factor: tt.fields.Factor,
			}
			b.AddItem(tt.args.i)
			if !reflect.DeepEqual(b.Factor, 10) {
				t.Errorf("Failed, %v", b.Factor)
			}
		})
	}
}
