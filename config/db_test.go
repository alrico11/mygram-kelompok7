package config

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestInitDB(t *testing.T) {
	type args struct {
		username string
		password string
		host     string
		port     string
		dbName   string
	}
	tests := []struct {
		name string
		args args
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitDB(tt.args.username, tt.args.password, tt.args.host, tt.args.port, tt.args.dbName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
