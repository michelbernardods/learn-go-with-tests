package main

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{name: "Sum two numbers", args: args{50, 7}, want: 57},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumTest(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Expected = %v, want %v", got, tt.want)
			}
		})
	}
}
