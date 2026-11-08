package utils

import (
	"strings"
	"testing"

	// "github.com/stretchr/testify/mock"
)

func Test_generateJWT(t *testing.T) {
	type args struct{
		username string
		
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		args	args	
		wantErr  bool
	}{
		{
			name: "got an empty string",
			args: args{
				username: "",

			},
			wantErr: false,
		},
		{
			name: "sukses generate token",
			args: args{
				username: "ajji",
			},
			wantErr: false,
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := generateJWT(tt.args.username)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("generateJWT() failed: %v", gotErr)
				}
				return
			}
			if(gotErr != nil) != tt.wantErr{
				t.Errorf("generate jwt mengalami error %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if got == ""{
				t.Error("expected token string, got empt string")
			}

			if !strings.Contains(got, "."){
				t.Errorf("expected JWT token structure with dots got: %s", got)
			}
		})
	}
}


func Test_validateJwt(t *testing.T) {
	type args struct{
		username string
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		args	args
		mock	func(args args) string
		wantUser        string
		wantErr     bool
	}{
		{
			name: "jwt sukses tervalidasi",
			args: args{
				username: "garpit",

			},
			mock: func(args args) string {
				tokenString, _ := generateJWT(args.username)
				return tokenString
			},
			wantUser: "garpit",
			wantErr: false,
		},
		{
			name: "invalid token",
		},
		{
			name: "expired token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenString := tt.mock(tt.args)
			got, gotErr := validateJwt(tokenString)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("validateJwt() failed: %v", gotErr, tt.wantErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("validateJwt() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("validateJwt() = %v, want %v", got.Username, tt.wantUser)
			}
		})
	}
}
