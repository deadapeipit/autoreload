package model

import "testing"

func TestStatus_CheckStatus(t *testing.T) {
	t.Log("Test func Randomvaluestatus")

	for i := 0; i <= 5; i++ {
		for l := 0; l <= 6; l++ {
			a := Status{
				Water: i,
				Wind:  l,
			}
			if gotStatus := a.CheckStatus(); gotStatus != "SAFE" {
				t.Errorf("try no= %d gotstatus = %v, expect %v", i, gotStatus, "SAFE")
			}
		}
	}
	for i := 6; i <= 8; i++ {
		for l := 0; l <= 15; l++ {
			a := Status{
				Water: i,
				Wind:  l,
			}
			if gotStatus := a.CheckStatus(); gotStatus != "STANDBY" {
				t.Errorf("try no= %d gotstatus = %v, expect %v", i, gotStatus, "STANDBY")
			}
		}
	}
	for i := 0; i <= 8; i++ {
		for l := 7; l <= 15; l++ {
			a := Status{
				Water: i,
				Wind:  l,
			}
			if gotStatus := a.CheckStatus(); gotStatus != "STANDBY" {
				t.Errorf("try no= %d gotstatus = %v, expect %v", i, gotStatus, "STANDBY")
			}
		}
	}
	for i := 9; i <= 100; i++ {
		for l := 0; l <= 100; l++ {
			a := Status{
				Water: i,
				Wind:  l,
			}
			if gotStatus := a.CheckStatus(); gotStatus != "DANGER" {
				t.Errorf("try no= %d gotstatus = %v, expect %v", i, gotStatus, "DANGER")
			}
		}
	}
	for i := 0; i <= 100; i++ {
		for l := 16; l <= 100; l++ {
			a := Status{
				Water: i,
				Wind:  l,
			}
			if gotStatus := a.CheckStatus(); gotStatus != "DANGER" {
				t.Errorf("try no= %d gotstatus = %v, expect %v", i, gotStatus, "DANGER")
			}
		}
	}
}
