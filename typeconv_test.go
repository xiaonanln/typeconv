package typeconv

import (
	"testing"
	"encoding/json"
	"log"
)

func interfaceVal(v interface{}) interface{} {
	data, err := json.Marshal(v)
	if err != nil{
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
	log.Printf("%T %v => %T %v", 1, 1, interfaceVal(1),interfaceVal(1))

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

func testFloatToInt(t *testing.T) {
	defer func() {
		recover()
	}()
	Int(interfaceVal(1.1)) // should panic
	t.Fatalf("Int 1.1 should not be 1") // should not goes here
}

