package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// Decoder is used to unpack csv information.
type Decoder interface {
	Fields() int
	Decode([][]string) error
}

// Encode is used to pack csv information.
type Encoder interface {
	Header() []string
	Encode() [][]string
}

// Read reads csv input and decodes the data into a given decoder interface.
func Read(rd io.Reader, dec Decoder) error {

	reader := csv.NewReader(rd)

	reader.Comment = '#'
	reader.FieldsPerRecord = dec.Fields()

	data, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var items [][]string
	for _, d := range data {
		var row []string
		for _, r := range d {
			row = append(row, r)
		}
		items = append(items, row)
	}

	if err := dec.Decode(items); err != nil {
		return err
	}

	return nil
}

// ReadFile reads csv input from a byte listand decodes elements into a given decoder interface.
func ReadBytes(data []byte, dec Decoder) error {
	return Read(bytes.NewBuffer(data), dec)
}

// ReadFile reads csv input from a file and decodes elements into a given decoder interface.
func ReadFile(path string, dec Decoder) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := Read(file, dec); err != nil {
		return err
	}

	return nil
}

// Write encodes and writes a given decoder interface.
func Write(writer io.Writer, enc Encoder) error {
	var rows [][]string

	res := enc.Encode()
	if res == nil {
		return nil
	}

	var header []string
	for n, h := range enc.Header() {
		switch n {
		case 0:
			header = append(header, "#"+h)
		default:
			header = append(header, h)
		}
	}

	w := csv.NewWriter(writer)
	rows = append(rows, header)
	rows = append(rows, res...)
	w.WriteAll(rows)

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

// WriteBytes encodes a given encoder interface into a byte slice.
func WriteBytes(enc Encoder) ([]byte, error) {
	var buf bytes.Buffer

	if err := Write(&buf, enc); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// WriteFile encodes a given encoder interface into a csv file.
func WriteFile(path string, enc Encoder) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := Write(file, enc); err != nil {
		return err
	}

	return nil
}
