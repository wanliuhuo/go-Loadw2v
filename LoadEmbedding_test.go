package liuhuo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadText(t *testing.T) {
	filepath := "./testdata/text.model"
	dimension := 10
	want := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m, err := LoadText(filepath, dimension)
	if err != nil {
		t.Fatal(err)
	}
	v := m.Find("魅力")
	got := v.vec
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want: %v , but get: %v\n", got, want)
	}
}
func TestModelFind(t *testing.T) {
	filepath := "./testdata/text.model"
	dimension := 10
	m, err := LoadText(filepath, dimension)
	if err != nil {
		t.Fatal(err)
	}
	v := m.Find("魅力")
	if v == nil {
		t.Fatalf("must be nil: %v", v)
	}
	fmt.Println(v)
}

func TestCosSim(t *testing.T) {
	filepath := "./testdata/text.model"
	dimension := 10
	m, err := LoadText(filepath, dimension)
	if err != nil {
		t.Fatal(err)
	}
	v1 := m.Find("魅力")
	v2 := m.Find("美丽")
	result := CosineSim(v1, v2)
	if result < float64(-998) {
		t.Fatalf("not the same dimension %f\n", result)
	}
	fmt.Println(result)
}
