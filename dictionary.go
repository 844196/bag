package main

import (
	"bytes"
	"fmt"
	goltsv "github.com/ymotongpoo/goltsv"
	"math/rand"
	"time"
)

type Dictionary []*Quote

func (d *Dictionary) SelectBySpeaker(speaker string) *Dictionary {
	var selected Dictionary
	size := len(*d)

	for i := 0; i < size; i++ {
		if item := (*d)[i]; item.Speaker == speaker {
			selected = append(selected, item)
		}
	}

	return &selected
}

func (d *Dictionary) SelectRandom() (*Quote, error) {
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

	meta, body := lines[0], lines[1:]

	filetype, isExistsFiletypeKey := meta["filetype"]
	if false == isExistsFiletypeKey || filetype != "polyaness_dict" {
		return nil, fmt.Errorf("Invalid filetype")
	}

	lineSize := len(body)
	dictionary := make(Dictionary, lineSize)
	for i := 0; i < lineSize; i++ {
		dictionary[i] = &Quote{Speaker: body[i]["speaker"], Content: body[i]["quote"]}
	}

	return &dictionary, nil
}
