package main

import (
    "io/ioutil"
    "fmt"
    "bytes"
    "os"

    bag "github.com/844196/bag/struct"
    goltsv "github.com/ymotongpoo/goltsv"
    flags "github.com/jessevdk/go-flags"
)

func builtinDict() ([]byte, error) {
    return Asset("data/dict.ltsv")
}

func readDict(path string) ([]byte, error) {
    return ioutil.ReadFile(path)
}

func convertDict(rawDict []byte) (bag.Dictionary, error) {
    ioreader := bytes.NewReader(rawDict)
    reader := goltsv.NewReader(ioreader)
    lines, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    value, isExistsFiletypeKey := lines[0]["filetype"]
    if false == isExistsFiletypeKey || value != "polyaness_dict" {
        return nil, fmt.Errorf("Invalid dictionary file")
    }

    size := len(lines)
    result := make([]bag.Quote, size)
    for i := 0; i < size; i++ {
        result[i] = bag.Quote{
            Speaker: lines[i]["speaker"],
            Content: lines[i]["quote"],
        }
    }

    return result[1:], nil
}

var opts struct {
    Version bool `short:"v" long:"version" description:"Show version"`
    File string `short:"f" long:"file" description:"Specifies dictionary file path" value-name:"PATH"`
    Speaker string `short:"s" long:"speaker" description:"Specifies quote speaker name" value-name:"SPEAKER"`
}

var parser = flags.NewParser(&opts, flags.Default)

func main() {
    _, err := parser.Parse()
    if err != nil {
        flagsErr, ok := err.(*flags.Error)
        if ok && flagsErr.Type == flags.ErrHelp {
            os.Exit(0)
        } else {
            os.Exit(1)
        }
    }
    if opts.Version {
        fmt.Println("0.1.0")
        os.Exit(0)
    }

    var rawDict []byte
    if opts.File == "" {
        rawDict, err = builtinDict()
    } else {
        rawDict, err = readDict(opts.File)
    }
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    dict, err := convertDict(rawDict)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    if opts.Speaker != "" {
        dict = *dict.SelectBySpeaker(opts.Speaker)
    }

    quote, err := dict.SelectRandom()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(quote.Content)
}
