package homework

import (
	//"reflect"
	"testing"
)

func Test_average(t *testing.T) {
	type args struct {
		input [15]float32
	}
	tests := []struct {
		name       string
		args       args
		wantResult float32
	}{
		{
			name: "positive_case_01",
			args: args{
				input: [15]float32{3, 8, 16, 4, 5, 6, 5, 134, 2, 5, 6, 7, 12, 34, 2},
			},
			wantResult: 16.6,
		},
		{
			name: "positive_case_02",
			args: args{
				input: [15]float32{4, 12, 16, 4, 5, 6, 5, 444, 4, 5, 6, 12, 12, 34, 2},
			},
			wantResult: 38.066666,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := average(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("average() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
