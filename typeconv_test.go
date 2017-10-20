package typeconv

import (
	"encoding/json"
	"log"
	"testing"
)

func interfaceVal(v interface{}) interface{} {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	var v2 interface{}
	err = json.Unmarshal(data, &v2)
	if err != nil {
		panic(err)
	}
	return v2
}

func TestInt(t *testing.T) {
	log.Printf("%T %v => %T %v", 1, 1, interfaceVal(1), interfaceVal(1))

	if Int(interfaceVal(1)) != 1 {
		t.Fatalf("Int 1 should not be 1")
	}

	if Int(interfaceVal(int32(1))) != 1 {
		t.Fatalf("Int 1 should not be 1")
	}

	if Int(interfaceVal(1.0)) != 1 {
		t.Fatalf("Int 1.0 should not be 1")
	}

	testFloatToInt(t)
}

func TestIntTuple(t *testing.T) {
	if !isIntTupleEqual(IntTuple([]int{1, 2, 3}), []int64{1, 2, 3}) {
		t.Fatalf("fail")
	}
	if !isIntTupleEqual(IntTuple([]int64{1, 2, 3}), []int64{1, 2, 3}) {
		t.Fatalf("fail")
	}
	if isIntTupleEqual(IntTuple([]int64{1, 2, 3, 4}), []int64{1, 2, 3}) {
		t.Fatalf("fail")
	}
}

type AnotherStringType string

func TestString(t *testing.T) {
	s := "abc"
	if String(interface{}(s)) != "abc" {
		t.Fatalf("String convert failed")
	}
	s2 := AnotherStringType("xxx")
	if String(s2) != string(s2) {
		t.Fatalf("String convert failed")
	}
}

func isIntTupleEqual(L1, L2 []int64) bool {
	if len(L1) != len(L2) {
		return false
	}

	for i, v := range L1 {
		if v != L2[i] {
			return false
		}
	}
	return true
}

func testFloatToInt(t *testing.T) {
	Int(interfaceVal(1.1)) // should panic
}
