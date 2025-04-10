package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPercentTab_Calculate(t *testing.T) {
	t.Parallel()

	type args struct {
		input1 string
		input2 string
	}
	tests := []struct {
		name    string
		srv     *PercentTab
		args    args
		want    string
		wantErr error
	}{
		{
			name: "success",
			srv:  &PercentTab{},
			args: args{
				input1: "10",
				input2: "20",
			},
			want:    "Result: 50.00%",
			wantErr: nil,
		},
		{
			name: "error",
			srv:  &PercentTab{},
			args: args{
				input1: "",
				input2: "20",
			},
			want:    "",
			wantErr: fmt.Errorf("calculate: first value is empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.srv.Calculate(tt.args.input1, tt.args.input2)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestPercentTab_calculate(t *testing.T) {
	t.Parallel()

	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name    string
		srv     *PercentTab
		args    args
		want    float64
		wantErr error
	}{
		{
			name: "success",
			srv:  &PercentTab{},
			args: args{
				s1: "10",
				s2: "20",
			},
			want:    50.00,
			wantErr: nil,
		},
		{
			name: "error: first value is empty",
			srv:  &PercentTab{},
			args: args{
				s1: "",
				s2: "20",
			},
			want:    0,
			wantErr: fmt.Errorf("first value is empty"),
		},
		{
			name: "error: second value is empty",
			srv:  &PercentTab{},
			args: args{
				s1: "10",
				s2: "",
			},
			want:    0,
			wantErr: fmt.Errorf("second value is empty"),
		},
		{
			name: "error: first value is not a number",
			srv:  &PercentTab{},
			args: args{
				s1: "10a",
				s2: "20",
			},
			want:    0,
			wantErr: fmt.Errorf("first value is not a number: strconv.ParseFloat: parsing \"10a\": invalid syntax"),
		},
		{
			name: "error: second value is not a number",
			srv:  &PercentTab{},
			args: args{
				s1: "10",
				s2: "20a",
			},
			want:    0,
			wantErr: fmt.Errorf("second value is not a number: strconv.ParseFloat: parsing \"20a\": invalid syntax"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.srv.calculate(tt.args.s1, tt.args.s2)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
