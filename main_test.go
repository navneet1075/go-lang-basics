package main

import "testing"

func Test_person_speak(t *testing.T) {
	type fields struct {
		fname string
		lname string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			per := person{
				fname: tt.fields.fname,
				lname: tt.fields.lname,
			}
			per.speak()
		})
	}
}

func Test_secretAgent_speak(t *testing.T) {
	type fields struct {
		person        person
		licenseToKill bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			per := secretAgent{
				person:        tt.fields.person,
				licenseToKill: tt.fields.licenseToKill,
			}
			per.speak()
		})
	}
}
