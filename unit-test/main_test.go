package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambah(t *testing.T) {
	actual := Tambah(1, 1)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestKurang(t *testing.T) {
	assert.Equal(t, 1, Kurang(2, 1))
}

func TestKali(t *testing.T) {
	assert.Equal(t, 4, Kali(2, 2))
}

func TestBagi(t *testing.T) {
	assert.Equal(t, 2, Bagi(4, 2))
}
