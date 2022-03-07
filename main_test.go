package main

import (
	_ "net/http"
	"os"
	"testing"
)

func Test_getEnvOrDefault(t *testing.T) {
	tests := []struct {
		key  string
		def  string
		want string
	}{
		//Lang overrides
		{key: "LANG", def: "es_ES.UTF-8", want: "en_US.UTF-8"},
		//Default value should take effect
		{key: "HTTPS_PORT", def: "8443", want: "8443"},
		//Empty value and non-existent should be empty
		{key: "INVENTED", def: "", want: ""},
		//Gopher is always awesome
		{key: "GOPHER", def: "bad", want: "awesome"},
	}
	os.Setenv("LANG", "en_US.UTF-8")
	os.Setenv("GOPHER", "awesome")
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := getEnvOrDefault(tt.key, tt.def); got != tt.want {
				t.Errorf("getEnvOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
