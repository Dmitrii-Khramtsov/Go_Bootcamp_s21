package main

import (
	"testing"
)

func TestDescribePlant_UnknownPlant(t *testing.T) {
	plant := UnknownPlant{
		FlowerType: "Achilléa",
		LeafType:   "millefólium",
		Color:      119,
	}

	got, err := describePlant(plant)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "FlowerType:Achilléa, LeafType:millefólium, Color(color_scheme=rgb):119"
	if got != want {
		t.Errorf("unexpected result:\n got: %s\nwant: %s", got, want)
	}
}

func TestDescribePlant_AnotherUnknownPlant(t *testing.T) {
	plant := AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType:    "lanceolate",
		Height:      15,
	}

	got, err := describePlant(plant)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "FlowerColor:10, LeafType:lanceolate, Height(unit=inches):15"
	if got != want {
		t.Errorf("unexpected result:\n got: %s\nwant: %s", got, want)
	}
}

func TestDescribePlant_NotAStruct(t *testing.T) {
	input := 42 // not a struct

	_, err := describePlant(input)
	if err == nil {
		t.Fatal("expected error for non-struct input, got nil")
	}
}

func TestDescribePlant_EmptyStruct(t *testing.T) {
	type Empty struct{}

	plant := Empty{}

	got, err := describePlant(plant)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "" {
		t.Errorf("expected empty string, got: %q", got)
	}
}