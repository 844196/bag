package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"io/ioutil"
	"os"
)

type Options struct {
	Version bool   `short:"v" long:"version" description:"Show version"`
	File    string `short:"f" long:"file" description:"Specifies dictionary file path" value-name:"PATH"`
	Speaker string `short:"s" long:"speaker" description:"Specifies quote speaker name" value-name:"SPEAKER"`
}

func run(args []string) int {
	options := Options{}
	_, err := flags.ParseArgs(&options, args)

	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			return 0
		} else {
			return 1
		}
	}

	if options.Version {
		fmt.Println("0.1.1")
		return 0
	}

	var rawDict []byte
	if options.File == "" {
		rawDict, err = Asset("data/dict.ltsv")
	} else {
		rawDict, err = ioutil.ReadFile(options.File)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	dictionary, err := NewDictionary(rawDict)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if options.Speaker != "" {
		dictionary = dictionary.SelectBySpeaker(options.Speaker)
	}

	quote, err := dictionary.SelectRandom()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Println(quote.Content())

	return 0
}

func main() {
	os.Exit(run(os.Args))
}
