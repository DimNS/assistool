package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func TestIDTab_GenerateUUID(t *testing.T) { //nolint:dupl // it's ok
	t.Parallel()

	type args struct {
		upperCase bool
	}
	tests := []struct {
		name string
		srv  *IDTab
		args args
		want string
	}{
		{
			name: "lower",
			srv: &IDTab{
				uuidFunc: func() string {
					return "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"
				},
			},
			args: args{
				upperCase: false,
			},
			want: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
		},
		{
			name: "upper",
			srv: &IDTab{
				uuidFunc: func() string {
					return "A0EEBC99-9C0B-4EF8-BB6D-6BB9BD380A11"
				},
			},
			args: args{
				upperCase: true,
			},
			want: "A0EEBC99-9C0B-4EF8-BB6D-6BB9BD380A11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.srv.GenerateUUID(tt.args.upperCase)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIDTab_GenerateULID(t *testing.T) { //nolint:dupl // it's ok
	t.Parallel()

	type args struct {
		upperCase bool
	}
	tests := []struct {
		name string
		srv  *IDTab
		args args
		want string
	}{
		{
			name: "lower",
			srv: &IDTab{
				ulidFunc: func() string {
					return "01hvwtnnraqwq8dqfgq2944yfq"
				},
			},
			args: args{
				upperCase: false,
			},
			want: "01hvwtnnraqwq8dqfgq2944yfq",
		},
		{
			name: "upper",
			srv: &IDTab{
				ulidFunc: func() string {
					return "01HVWTNNRAQWQ8DQFGQ2944YFQ"
				},
			},
			args: args{
				upperCase: true,
			},
			want: "01HVWTNNRAQWQ8DQFGQ2944YFQ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.srv.GenerateULID(tt.args.upperCase)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_newUUID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := newUUID()
			id, err := uuid.Parse(got)
			assert.NotNil(t, id)
			assert.Nil(t, err)
		})
	}
}

func Test_newULID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := newULID()
			id, err := ulid.Parse(got)
			assert.NotNil(t, id)
			assert.Nil(t, err)
		})
	}
}
