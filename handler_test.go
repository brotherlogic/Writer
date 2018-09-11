package main

import (
	"fmt"
	"testing"
)

type testWriter struct {
	fail bool
}

func (t *testWriter) writeLocation(name string, lat, lon float32, ip, port string) error {
	if t.fail {
		return fmt.Errorf("Built to fail")
	}
	return nil
}

func TestHandle(t *testing.T) {
	s := &Server{"", "", &testWriter{fail: false}}
	res := s.handle("test", "12.5", "12.5", "madeup", "23")
	if res != "Location Written" {
		t.Errorf("Failed: %v", res)
	}
}

func TestHandleFailLat(t *testing.T) {
	s := &Server{"", "", &testWriter{fail: false}}
	res := s.handle("test", "12.5gg", "12.5", "madeup", "23")
	if res == "Location Written" {
		t.Errorf("Failed: %v", res)
	}
}

func TestHandleFailLon(t *testing.T) {
	s := &Server{"", "", &testWriter{fail: false}}
	res := s.handle("test", "12.5", "12.5gg", "madeup", "23")
	if res == "Location Written" {
		t.Errorf("Failed: %v", res)
	}
}

func TestHandleFailWrite(t *testing.T) {
	s := &Server{"", "", &testWriter{fail: true}}
	res := s.handle("test", "12.5", "12.5", "madeup", "23")
	if res == "Location Written" {
		t.Errorf("Failed: %v", res)
	}
}
