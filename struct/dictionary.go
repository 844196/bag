package bag

import (
    "math/rand"
    "time"
    "fmt"
)

type Dictionary []Quote

func (d *Dictionary) SelectBySpeaker(speaker string) *Dictionary {
    var result Dictionary
    size := len(*d)

    for i := 0; i < size; i++ {
        if item := (*d)[i]; item.Speaker == speaker {
            result = append(result, item)
        }
    }

    return &result
}

func (d *Dictionary) SelectRandom() (*Quote, error) {
    if len(*d) == 0 {
        return nil, fmt.Errorf("filterd by speaker or dictionary all empty")
    }

    rand.Seed(time.Now().UnixNano())
    idx := rand.Intn(len(*d))

    return &(*d)[idx], nil
}
