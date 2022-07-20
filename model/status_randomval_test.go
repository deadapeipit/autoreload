package model

import (
	"testing"
)

func TestRandomValueStatus(t *testing.T) {
	t.Log("Test func Randomvaluestatus")

	for i := 0; i < 10000; i++ {
		if gotStatus := RandomValueStatus(); gotStatus.Status.Water > 100 || gotStatus.Status.Wind > 100 {
			t.Errorf("waterstatus = %v, windstatus %v", gotStatus.Status.Water, gotStatus.Status.Wind)
		}
	}
}
