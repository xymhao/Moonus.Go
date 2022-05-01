package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isErr(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"wrap demo true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, wrapErr())
		})
	}
}

func TestOpaqueErr(t *testing.T) {
	tests := []struct {
		name string
		want errorMo
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, OpaqueErr())
		})
	}
}

func Test_errorMo_IsMoonusCall(t *testing.T) {
	type fields struct {
		s string
		e error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// test IsTemporary
		{"1", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errorMo{
				s: tt.fields.s,
				e: tt.fields.e,
			}
			if got := err.IsMoonusCall(); got != tt.want {
				t.Errorf("IsMoonusCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
