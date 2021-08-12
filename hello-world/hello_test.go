package main

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name     string
		language string
	}
	tests := []struct {
		description string
		args        args
		want        string
	}{
		{description: "say hello to people", args: args{"Michel", "new"}, want: "Hello Michel"},
		{description: "says 'hello world' when an empty string is passed", args: args{"", ""}, want: "Hello world"},
		{description: "in english", args: args{"Michel", "english"}, want: "Hello Michel"},
		{description: "in espanhol", args: args{"Michel", "espanhol"}, want: "Hola Michel"},
		{description: "in frances", args: args{"Michel", "frances"}, want: "Bonjour Michel"},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := Hello(tt.args.name, tt.args.language); got != tt.want {
				t.Errorf("Expected = %v, want %v", got, tt.want)
			}
		})
	}
}
