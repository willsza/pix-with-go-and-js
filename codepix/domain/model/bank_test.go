package model

import "testing"

func TestBank_isValid(t *testing.T) {
	type fields struct {
		Base Base
		Code string
		Name string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := &Bank{
				Base: tt.fields.Base,
				Code: tt.fields.Code,
				Name: tt.fields.Name,
			}
			if err := bank.isValid(); (err != nil) != tt.wantErr {
				t.Errorf("Bank.isValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
