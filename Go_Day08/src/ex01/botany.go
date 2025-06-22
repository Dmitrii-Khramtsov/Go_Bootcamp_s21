package main

import (
	"fmt"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType string
	Color int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType string
	Height int `unit:"inches"`
}

func describePlant(plant any) (string, error) {
	var s string
	v := reflect.ValueOf(plant)
	t := reflect.TypeOf(plant)
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected struct but got %s", v.Kind())
	}

	for i := range t.NumField() {
		value := v.Field(i).Interface()
		field := t.Field(i)
		color := field.Tag.Get("color_scheme");
		unit := field.Tag.Get("unit");

		switch {
		case  color != "":
			s += fmt.Sprintf("%v(color_scheme=%v):%v", field.Name, color, value)
		case unit != "":
			s += fmt.Sprintf("%v(unit=%v):%v", field.Name, unit, value)
		default:
			s += fmt.Sprintf("%v:%v", field.Name, value)
		}

		if i < t.NumField()-1 {
			s += ", "
		}
	}

	return s, nil
}

func main() {
	plant := UnknownPlant{
		FlowerType: "Achilléa",
		LeafType: "millefólium",
		Color: 119,
	}
	plant2 := AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType: "lanceolate",
		Height: 15,
	}

	res, err := describePlant(plant)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	res2, err := describePlant(plant2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res2)
}
