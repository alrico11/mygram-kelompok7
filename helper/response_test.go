package helper

import (
	"reflect"
	"testing"
)

func TestAPIResponse(t *testing.T) {
	type args struct {
		status string
		data   interface{}
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := APIResponse(tt.args.status, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("APIResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatValidationError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatValidationError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatValidationError() = %v, want %v", got, tt.want)
			}
		})
	}
}
