package config

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestInitDB(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
