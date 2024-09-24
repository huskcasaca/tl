package tl_test

import "testing"

type TestCase interface {
	Name() string
	Run(t *testing.T)
}
