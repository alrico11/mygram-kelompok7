package middleware

import (
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt"
)

func Test_jwtService_GenerateToken(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		s       *jwtService
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GenerateToken(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtService.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("jwtService.GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtService_ValidateToken(t *testing.T) {
	type args struct {
		encodedToken string
	}
	tests := []struct {
		name    string
		s       *jwtService
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ValidateToken(tt.args.encodedToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtService.ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jwtService.ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
