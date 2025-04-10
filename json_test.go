package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONTab_Beautify(t *testing.T) {
	t.Parallel()

	type args struct {
		json      string
		ntconvert bool
	}
	tests := []struct {
		name    string
		srv     *JSONTab
		args    args
		want    string
		wantErr error
	}{
		{
			name: "simple",
			srv:  &JSONTab{},
			args: args{
				json:      `{"level":"info","time":"2024-03-14T12:36:45.977Z"}`,
				ntconvert: false,
			},
			want: `{
    "level": "info",
    "time": "2024-03-14T12:36:45.977Z"
}`,
		},
		{
			name: "unicode",
			srv:  &JSONTab{},
			args: args{
				json:      `{"Fio":"\u0418\u0432\u0430\u043d\u043e\u0432 \u0418\u0432\u0430\u043d \u0418\u0432\u0430\u043d\u043e\u0432\u0438\u0447"}`,
				ntconvert: false,
			},
			want: `{
    "Fio": "Иванов Иван Иванович"
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.srv.Beautify(tt.args.json, tt.args.ntconvert)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestJSONTab_jsonBeautify(t *testing.T) {
	t.Parallel()

	type args struct {
		str string
		ntv bool
	}
	tests := []struct {
		name    string
		srv     *JSONTab
		args    args
		want    string
		wantErr error
	}{
		{
			name: "simple",
			srv:  &JSONTab{},
			args: args{
				str: `{"level":"info","time":"2024-03-14T12:36:45.977Z"}`,
				ntv: false,
			},
			want: `{
    "level": "info",
    "time": "2024-03-14T12:36:45.977Z"
}`,
			wantErr: nil,
		},
		{
			name: "empty string",
			srv:  &JSONTab{},
			args: args{
				str: ``,
				ntv: false,
			},
			want:    ``,
			wantErr: nil,
		},
		{
			name: "convert \\n and \\t",
			srv:  &JSONTab{},
			args: args{
				str: `{"level":"info","time":"2024-03-14T12:36:45.977Z","multiline":"CreateZone\n\tzone_service.go:86\nfunc1\n\tgrpc.pb.go:164"}`,
				ntv: true,
			},
			want: `{
    "level": "info",
    "multiline": "CreateZone
	zone_service.go:86
func1
	grpc.pb.go:164",
    "time": "2024-03-14T12:36:45.977Z"
}`,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.srv.jsonBeautify(tt.args.str, tt.args.ntv)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestJSONTab_bypassMap(t *testing.T) {
	t.Parallel()

	type args struct {
		data map[string]any
	}
	tests := []struct {
		name string
		srv  *JSONTab
		args args
		want map[string]any
	}{
		{
			name: "success",
			srv:  &JSONTab{},
			args: args{
				data: func() map[string]any {
					str := `{"level":"info","time":"2024-03-14T12:36:31.338Z","request":{"Orders":[{"WaybillNumber":"29-0241-9289","OrderNumber":"29-0241-9289"}]},"grpc.request.message":"{\"waybill_number\":\"29-0241-9289\", \"order_id\":\"01HRY96YW12SAPT3JF3YR0G52Y\", \"customer_kpp\":\"\"}"}`
					buff := bytes.NewBuffer([]byte(str))
					data := map[string]any{}

					_ = json.NewDecoder(buff).Decode(&data)

					return data
				}(),
			},
			want: map[string]any{
				"grpc.request.message": map[string]any{
					"customer_kpp":   "",
					"order_id":       "01HRY96YW12SAPT3JF3YR0G52Y",
					"waybill_number": "29-0241-9289",
				},
				"level": "info",
				"request": map[string]any{
					"Orders": []any{
						map[string]any{
							"OrderNumber":   "29-0241-9289",
							"WaybillNumber": "29-0241-9289",
						},
					},
				},
				"time": "2024-03-14T12:36:31.338Z",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.srv.bypassMap(tt.args.data)
			assert.Equal(t, tt.want, got)
		})
	}
}
