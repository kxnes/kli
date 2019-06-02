package kli

import (
	"os"
	"reflect"
	"testing"
)

func TestFloat64SliceFlagGetName(t *testing.T) {
	type fields struct {
		Name  string
		Value []float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"getter", fields{"-a", []float64{1, 2, 3}}, "-a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Float64SliceFlag{
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if got := f.GetName(); got != tt.want {
				t.Errorf("Float64SliceFlag.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64SliceFlagSetValue(t *testing.T) {
	type fields struct {
		Name  string
		Value []float64
	}
	type args struct {
		values []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		errMsg string
	}{
		{
			"one value",
			fields{"-a", []float64{1.1}},
			args{[]string{"1.1"}},
			"",
		},
		{
			"array values",
			fields{"-a", []float64{1.1, 2.2}},
			args{[]string{"1.1", "2.2"}},
			"",
		},
		{
			"array values with negative",
			fields{"-a", []float64{1.1, -2.2, 3.3}},
			args{[]string{"1.1", "-2.2", "3.3"}},
			"",
		},
		{
			"empty value",
			fields{"-a", []float64{}},
			args{[]string{}},
			"flag: -a, value: [] can't be set: len must be >= 1",
		},
		{"error while parsing",
			fields{"-a", []float64{}},
			args{[]string{"a", "b"}},
			"flag: -a, value: [a b] can't be set: strconv.ParseFloat: parsing \"a\": invalid syntax",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Float64SliceFlag{
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if err := f.SetValue(tt.args.values); (err != nil) && err.Error() != tt.errMsg {
				t.Errorf("Float64SliceFlag.SetValue() error = %v, want %v", err, tt.errMsg)
			}
			if !reflect.DeepEqual(tt.fields.Value, f.Value) {
				t.Errorf("Float64SliceFlag.Value = %v, want %v", f.Value, tt.fields.Value)
			}
		})
	}
}

func TestRuneFlagGetName(t *testing.T) {
	type fields struct {
		Name  string
		Value rune
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"getter", fields{"-o", '+'}, "-o"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &RuneFlag{
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if got := f.GetName(); got != tt.want {
				t.Errorf("RuneFlag.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneFlagSetValue(t *testing.T) {
	type fields struct {
		Name  string
		Value rune
	}
	type args struct {
		values []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		errMsg string
	}{
		{
			"usual",
			fields{"-o", '+'},
			args{[]string{"+"}},
			"",
		},
		{
			"empty value",
			fields{"-o", 0},
			args{[]string{}},
			"flag: -o, value: [] can't be set: len must be >= 1",
		},
		{
			"not rune value",
			fields{"-o", 0},
			args{[]string{"ab"}},
			"flag: -o, value: [ab] can't be set: to long argument",
		},
		{
			"to much values",
			fields{"-o", 0},
			args{[]string{"a", "b"}},
			"flag: -o, value: [a b] can't be set: to long argument",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &RuneFlag{
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if err := f.SetValue(tt.args.values); (err != nil) && err.Error() != tt.errMsg {
				t.Errorf("Float64SliceFlag.SetValue() error = %v, want %v", err, tt.errMsg)
			}
			if !reflect.DeepEqual(tt.fields.Value, f.Value) {
				t.Errorf("Float64SliceFlag.Value = %v, want %v", f.Value, tt.fields.Value)
			}
		})
	}
}

func TestArgsParse(t *testing.T) {
	type fields struct {
		Flags []Flag
	}
	tests := []struct {
		name   string
		fields fields
		osargs []string
	}{
		{
			"usual", fields{[]Flag{
			&RuneFlag{Name: "-o"},
			&Float64SliceFlag{Name: "-a"},
		}}, []string{"", "-a", "1", "2", "3", "-o", "+"}},
		{
			"usual (diff order)", fields{[]Flag{
			&RuneFlag{Name: "-o"},
			&Float64SliceFlag{Name: "-a"},
		}}, []string{"", "-o", "+", "-a", "1", "2", "3"}},
	}
	for _, tt := range tests {
		os.Args = tt.osargs

		t.Run(tt.name, func(t *testing.T) {
			args := &Args{
				Flags: tt.fields.Flags,
			}
			args.Parse()

			v1, v2 := args.Flags[0].(*RuneFlag), args.Flags[1].(*Float64SliceFlag)
			if v1.Value != '+' && !reflect.DeepEqual(v2.Value, []float64{1, 2, 3}) {
				t.Errorf("args.Flag[0] = %v, want %v", v1.Value, '+')
				t.Errorf("args.Flag[1] = %v, want %v", v2.Value, []float64{1, 2, 3})
			}
		})
	}
}
