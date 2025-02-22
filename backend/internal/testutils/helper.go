package testutils

import "testing"

func CheckErr(t *testing.T, err error, msg string) {
    if err != nil {
        t.Fatalf("%s: %v", msg, err)
    }
}