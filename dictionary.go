package main

import (
	"bytes"
	"fmt"
	goltsv "github.com/ymotongpoo/goltsv"
	"math/rand"
	"time"
)

type Dictionary []map[string]string

func (d *Dictionary) SelectBySpeaker(speaker string) *Dictionary {
	var selected Dictionary

	for i := 0; i < len(*d); i++ {
		if item := (*d)[i]; item["speaker"] == speaker {
			selected = append(selected, item)
		}
	}

	return &selected
}

func (d *Dictionary) SelectRandom() (map[string]string, error) {
	size := len(*d)

	if size == 0 {
		return nil, fmt.Errorf("filterd by speaker or dictionary all empty")
	}

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(size)

	return (*d)[idx], nil
}

func NewDictionary(data []byte) (*Dictionary, error) {
	ioreader := bytes.NewReader(data)
	ltsvreader := goltsv.NewReader(ioreader)

	lines, err := ltsvreader.ReadAll()
	if err != nil {
		return nil, err
	}

	filetype, isExistsFiletypeKey := lines[0]["filetype"]
	if false == isExistsFiletypeKey || filetype != "polyaness_dict" {
		return nil, fmt.Errorf("Invalid filetype")
	}

	dictionary := Dictionary(lines[1:])

	return &dictionary, nil
}
