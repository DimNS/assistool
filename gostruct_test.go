package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoStructTab_Beautify(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		srv     *GoStructTab
		input   string
		want    string
		wantErr error
	}{
		{
			name:  "simple",
			srv:   &GoStructTab{},
			input: `&postgres.shiftModel{BaseModel:schema.BaseModel{}, ID:"sid10", CreatedAt:time.Date(2019, time.January, 2, 12, 0, 0, 0, time.UTC), OrganizationID:"", Earnings:(*decimal.Decimal)(nil), OfferID:sql.NullString{String:"", Valid:false}, BikeRental:false, SuspendTime:<nil>, ActualStartTime:sql.NullTime{Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid:true}, Tags:sql.NullString{String:"", Valid:false}, CarryingCapacity:sql.NullInt32{Int32:0, Valid:false}, PartnerData:(*domain.ShiftPartnerData)(nil)}`,
			want: `&postgres.shiftModel{
  BaseModel:schema.BaseModel{},
  ID:"sid10",
  CreatedAt:time.Date(2019, time.January, 2, 12, 0, 0, 0, time.UTC),
  OrganizationID:"",
  Earnings:(*decimal.Decimal)(nil),
  OfferID:sql.NullString{
    String:"",
    Valid:false
  },
  BikeRental:false,
  SuspendTime:<nil>,
  ActualStartTime:sql.NullTime{
    Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
    Valid:true
  },
  Tags:sql.NullString{
    String:"",
    Valid:false
  },
  CarryingCapacity:sql.NullInt32{
    Int32:0,
    Valid:false
  },
  PartnerData:(*domain.ShiftPartnerData)(nil)
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.srv.Beautify(tt.input)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestGoStructTab_goStructBeautify(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		srv     *GoStructTab
		input   string
		want    string
		wantErr error
	}{
		{
			name:  "simple",
			srv:   &GoStructTab{},
			input: `&postgres.shiftModel{BaseModel:schema.BaseModel{}, ID:"sid10", CreatedAt:time.Date(2019, time.January, 2, 12, 0, 0, 0, time.UTC), OrganizationID:"", Earnings:(*decimal.Decimal)(nil), OfferID:sql.NullString{String:"", Valid:false}, BikeRental:false, SuspendTime:<nil>, ActualStartTime:sql.NullTime{Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), Valid:true}, Tags:sql.NullString{String:"", Valid:false}, CarryingCapacity:sql.NullInt32{Int32:0, Valid:false}, PartnerData:(*domain.ShiftPartnerData)(nil)}`,
			want: `&postgres.shiftModel{
  BaseModel:schema.BaseModel{},
  ID:"sid10",
  CreatedAt:time.Date(2019, time.January, 2, 12, 0, 0, 0, time.UTC),
  OrganizationID:"",
  Earnings:(*decimal.Decimal)(nil),
  OfferID:sql.NullString{
    String:"",
    Valid:false
  },
  BikeRental:false,
  SuspendTime:<nil>,
  ActualStartTime:sql.NullTime{
    Time:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
    Valid:true
  },
  Tags:sql.NullString{
    String:"",
    Valid:false
  },
  CarryingCapacity:sql.NullInt32{
    Int32:0,
    Valid:false
  },
  PartnerData:(*domain.ShiftPartnerData)(nil)
}`,
			wantErr: nil,
		},
		{
			name:    "empty string",
			srv:     &GoStructTab{},
			input:   ``,
			want:    ``,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.srv.goStructBeautify(tt.input)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
