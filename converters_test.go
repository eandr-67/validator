package validator

import (
	"testing"
)

func TestCovnert(t *testing.T) {
	res, err := Convert[string](nil)
	if res != nil {
		t.Errorf("Unexpected result: %v", res)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	res, err = Convert[string](125)
	if res == nil {
		t.Errorf("res should not be nil")
	} else if *res != "" {
		t.Errorf("res must be '', got '%v'", *res)
	}
	if err == nil {
		t.Errorf("err should not be nil")
	} else if *err != ErrMsg[CodeTypeIncorrect] {
		t.Errorf("err must be ErrMsg[CodeTypeIncorrect], got %v", err)
	}

	res, err = Convert[string]("abc")
	if res == nil {
		t.Errorf("res should not be nil")
	} else if *res != "abc" {
		t.Errorf("res must be 'abc', got '%v'", *res)
	}
	if err != nil {
		t.Errorf("err should be nil, got %v", *err)
	}
}

func TestConvertInt(t *testing.T) {
	res, err := convertInt(nil)
	if res != nil {
		t.Errorf("Unexpected result: %v", res)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	res, err = convertInt("aaa")
	if res == nil {
		t.Errorf("res should not be nil")
	} else if *res != 0 {
		t.Errorf("res must be 0, got '%v'", *res)
	}
	if err == nil {
		t.Errorf("err should not be nil")
	} else if *err != ErrMsg[CodeTypeIncorrect] {
		t.Errorf("err must be ErrMsg[CodeTypeIncorrect], got %v", err)
	}

	res, err = convertInt(int64(23))
	if res == nil {
		t.Errorf("res should not be nil")
	} else if *res != 23 {
		t.Errorf("res must be 23, got '%v'", *res)
	}
	if err != nil {
		t.Errorf("err should be nil, got %v", *err)
	}

	res, err = convertInt(float64(13.13))
	if res == nil {
		t.Errorf("res should not be nil")
	} else if *res != 13 {
		t.Errorf("res must be 13, got '%v'", *res)
	}
	if err != nil {
		t.Errorf("err should be nil, got %v", *err)
	}
}
