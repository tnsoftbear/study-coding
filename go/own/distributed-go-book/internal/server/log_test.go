package server

import (
	"reflect"
	"testing"
)

func TestLog_Append(t *testing.T) {
	type fields struct {
		records []Record
	}
	type args struct {
		record Record
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
	}{
		{
			name: "Add a record to the empty log",
			fields: fields{
				records: []Record{},
			},
			args: args{
				record: Record{Value: []byte{60, 61, 62}, Offset: 0},
			},
			want: 0,
		},
		{
			"Add a record to the log with one record",
			fields{
				records: []Record{
					{Value: []byte{60, 61, 62}, Offset: 0},
				},
			},
			args{
				record: Record{Value: []byte{60, 61, 62}, Offset: 0},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Log{
				records: tt.fields.records,
			}
			got, _ := c.Append(tt.args.record)
			if got != tt.want {
				t.Errorf("Log.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_Read(t *testing.T) {
	type fields struct {
		records []Record
	}
	type args struct {
		offset uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Record
		err    error
	}{
		{
			name: "Read a record from the empty log",
			fields: fields{
				records: []Record{},
			},
			args: args{
				offset: 0,
			},
			want: Record{},
			err:  ErrOffsetNotFound,
		},
		{
			name: "Read a record from the log with one record",
			fields: fields{
				records: []Record{
					{Value: []byte{60, 61, 62}, Offset: 0},
				},
			},
			args: args{
				offset: 0,
			},
			want: Record{Value: []byte{60, 61, 62}, Offset: 0},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Log{
				records: tt.fields.records,
			}
			got, err := c.Read(tt.args.offset)
			if err != tt.err {
				t.Errorf("Log.Read() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Log.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
