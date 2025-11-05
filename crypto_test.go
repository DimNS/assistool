package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCryptTab_Encrypt(t *testing.T) {
	tests := []struct {
		name    string
		srv     *CryptTab
		message string
		gcm     bool
	}{
		{
			name:    "With GCM: Case 1",
			srv:     &CryptTab{},
			message: "Hello, World!",
			gcm:     true,
		},
		{
			name:    "Without GCM: Case 1",
			srv:     &CryptTab{},
			message: "Hello, World!",
			gcm:     false,
		},
		}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.Encrypt("123456789012345678901234", tt.message, tt.gcm)
			assert.NoError(t, err)
			assert.NotEmpty(t, got)
		})
	}
}

func TestCryptTab_Decrypt(t *testing.T) {
	tests := []struct {
		name    string
		srv     *CryptTab
		encoded string
		gcm     bool
		want    string
	}{
		{
			name:    "With GCM: Case 1",
			srv:     &CryptTab{},
			encoded: "ueAj/CBN3J6uWsBPY9bFaHDqD23ZzxtI4gLo8G2yPo8ODKwUKaifBLM",
			gcm:     true,
			want:    "Hello, World!",
		},
		{
			name:    "With GCM: Case 2",
			srv:     &CryptTab{},
			encoded: "RglsQsTMrGnr75Rrch6BnwGVxp4v69hVQa+NaQ9BiQ092FaV+AQ7JYs",
			gcm:     true,
			want:    "Hello, World!",
		},
		{
			name:    "With GCM: Case 3",
			srv:     &CryptTab{},
			encoded: "vVNtZyUHwS2nSExmjKKwmmQl0YxBu/Hew5zULQDYfcZmLmPTCUQAags",
			gcm:     true,
			want:    "Hello, World!",
		},
		{
			name:    "Without GCM: Case 1",
			srv:     &CryptTab{},
			encoded: "4+42tmsAlxXJNoSPB+s3d/VPh8fCiTa3WiID5Vc",
			gcm:     false,
			want:    "Hello, World!",
		},
		{
			name:    "Without GCM: Case 2",
			srv:     &CryptTab{},
			encoded: "Vtl+inw48e3sUevJvJYoDLD3rsq0q783dENWc9c",
			gcm:     false,
			want:    "Hello, World!",
		},
		{
			name:    "Without GCM: Case 3",
			srv:     &CryptTab{},
			encoded: "1c0mbdK9lHWl7beeUPhxailWSShkF7fm2qC4K8g",
			gcm:     false,
			want:    "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.srv.Decrypt("123456789012345678901234", tt.encoded, tt.gcm)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
