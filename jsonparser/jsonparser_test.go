package jsonparser

import (
	"reflect"
	"testing"

	types "github.com/KodiakAS/balance/types"
)

func TestParseJsonFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.Items
		wantErr bool
	}{
		{
			name:    "normal",
			args:    args{"../tests/demo.json"},
			want:    &types.Items{Items: []types.Item{{Data: "Hello", Factor: 1}, {Data: "World", Factor: 2}}},
			wantErr: false,
		},
		{
			name:    "invalid_factor_type",
			args:    args{"../tests/invalid_type.json"},
			want:    &types.Items{},
			wantErr: true,
		},
		{
			name:    "file_not_found",
			args:    args{"../tests/jsckass.json"},
			want:    &types.Items{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJsonFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJsonFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJsonFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
