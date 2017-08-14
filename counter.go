package main

import (
	"strconv"
	"errors"
	"fmt"
)

type BytesFabric map[byte]*byte

func (this *BytesFabric) GetByString(t string) (*byte, error) {
	v, err := this.parseByString(t)
	if err != nil {
		return nil, err
	}
	if _, ok := this[v]; !ok {
		this[v] = &v
	}

	return this[v], nil
}

func (this *BytesFabric) parseByString(t string) (byte, error) {
	if len(t) < 15 {
		panic("Can't make hour by string having length less than 15")
	}
	if i, err := strconv.Atoi(t[11:13]); err != nil {
		return 0, errors.New("Can't convert string to int (hours), given " + t)
	} else {
		return byte(i), nil
	}

	return 0, nil
}

type IntFabric map[int64]*int64

func (this *IntFabric) Get(v int64) (*int64, error) {
	if _, ok := (*this)[v]; !ok {
		(*this)[v] = &v
	}

	return (*this)[v], nil
}

func (this *IntFabric) Print() {
	cnt := 0
	a := int64(0)
	for i, _ := range (*this) {
		a = a * int64(i)
		cnt++
	}

	fmt.Printf("IntFabric (%d): [\n", cnt)
	for i, v := range (*this) {
		fmt.Printf("\t%d: %d\n", i, *v)
	}
	fmt.Println("]")
}

type EntryAndType struct {
	Entry int64
	EntryType int64
}
func NewEntryAndType(entry, entryType int64) *EntryAndType {
	return &EntryAndType{entry, entryType}
}

type EntryAndTypeFabric map[string]*EntryAndType
func (e EntryAndTypeFabric) Get(entry, entryType int64) {
	key := fmt.Sprintf("%d|%d", entry, entryType)
	if _, ok := e[key]; !ok {
		e[key] = NewEntryAndType(entry, entryType)
	}
	return e[key]
}

type DayOfWeek struct {
	DayOfWeek byte
	EntryAndType *EntryAndType
}
func NewDayOfWeek(dayOfWeek byte, entryAndType *EntryAndType) *DayOfWeek {
	return &DayOfWeek{dayOfWeek, entryAndType}
}



type Counter struct {

}

func NewCounter() {

}