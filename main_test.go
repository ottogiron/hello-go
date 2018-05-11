package main

import "testing"

func Test_sum(t *testing.T) {

	type args struct {
		val1 int
		val2 int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Successful sum", args{val1: 1, val2: 2}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sum(tt.args.val1, tt.args.val2)
			if (err != nil) != tt.wantErr {
				t.Errorf("sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
