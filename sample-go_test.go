package main

import (
    "testing"
)


func TestSample(t *testing.T) {
    calced, err := forTesting(10, 2)
    if err != nil {
        t.Errorf("Expecting no error, received: %v", err)
    }

	if calced != 20 {
		t.Errorf("Expected return value of 20, got %d instead", calced)
	}

    calced, err = forTesting(10, 6)
    if err == nil {
        t.Errorf("Expecting error, did not get one!")
    }

	if calced != 0 {
		t.Errorf("Expected return value of 0 due to empty return int, got %d instead", calced)
	}
}
