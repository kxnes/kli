package kli

import (
	"fmt"
	"os"
	"strconv"
)

const (
	flagError  = "flag: %v, value: %v can't be set: %v"
	parseError = "error: \"%v\" while parsing %v\n"
)

type Args struct {
	Flags []Flag
}

type Flag interface {
	GetName() string
	SetValue([]string) error
}

type Float64SliceFlag struct {
	Name  string
	Value []float64
}

func (f *Float64SliceFlag) GetName() string {
	return f.Name
}

func (f *Float64SliceFlag) SetValue(values []string) (err error) {
	if len(values) < 1 {
		return fmt.Errorf(flagError, f.Name, values, "len must be >= 1")
	}

	f.Value = make([]float64, len(values))

	for i, v := range values {
		f.Value[i], err = strconv.ParseFloat(v, 64)
		if err != nil {
			f.Value = []float64{}
			return fmt.Errorf(flagError, f.Name, values, err)
		}
	}

	return
}

type RuneFlag struct {
	Name  string
	Value rune
}

func (f *RuneFlag) GetName() string {
	return f.Name
}

func (f *RuneFlag) SetValue(values []string) error {
	f.Value = 0

	if len(values) < 1 {
		return fmt.Errorf(flagError, f.Name, values, "len must be >= 1")
	}

	if len(values) != 1 || len(values[0]) != 1 {
		return fmt.Errorf(flagError, f.Name, values, "to long argument")
	}

	f.Value = rune(values[0][0])
	return nil
}

func (args *Args) Parse() {
	os.Args = os.Args[1:]

	s := make(map[string][]string)
	for _, f := range args.Flags {
		s[f.GetName()] = make([]string, 0, cap(os.Args))
	}

	for i, v := range os.Args {
		// check that is a flag key
		if _, ok := s[v]; ok {
			for _, a := range os.Args[i+1:] {
				// check that is a not flag key
				if _, ok := s[a]; !ok {
					s[v] = append(s[v], a)
				} else {
					break
				}
			}
		}
	}

	for _, f := range args.Flags {
		err := f.SetValue(s[f.GetName()])
		if err != nil {
			fmt.Printf(parseError, err, os.Args)
			os.Exit(1)
		}
	}
}
