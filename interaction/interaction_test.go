package main

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		characters     string
		numberInteger  int
		numberFloating float64
	}
	tests := []struct {
		description string
		args        args
		wantString  string
		wantInt     int
		wantFloat   float64
	}{
		{description: "Test repeat in String", args: args{characters: "a"}, wantString: "aaaaa"},
		{description: "Test repeat in String", args: args{characters: "cu"}, wantString: "cucucucucu"},
		{description: "Test repeat in Integer", args: args{numberInteger: 1}, wantInt: 5},
		{description: "Test repeat in Integer", args: args{numberInteger: 100}, wantInt: 500},
		{description: "Test repeat in Float", args: args{numberFloating: 2.5}, wantFloat: 12.5},
		{description: "Test repeat in Float", args: args{numberFloating: 74.9}, wantFloat: 374.5},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if got := RepeatString(tt.args.characters); got != tt.wantString {
				t.Errorf("Expected = %v, want %v", got, tt.wantString)
			}
			if got := RepeatInteger(tt.args.numberInteger); got != tt.wantInt {
				t.Errorf("Expected = %v, want %v", got, tt.wantInt)
			}
			if got := RepeatFloat(tt.args.numberFloating); got != tt.wantFloat {
				t.Errorf("Expected = %v, want %v", got, tt.wantFloat)
			}

		})
	}
}
