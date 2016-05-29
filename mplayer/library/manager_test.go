package library

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManger failed, not empty.")
	}

	m0 := &MusicEntry {
		"1", "My Heart Will Go On", "Celion Dion",
		"http://qbox.me/24501234", "MP3"}
	mm.Add(m0)

	m0 = &MusicEntry {
		"1", "I'm Not Affried", "em",
		"http://qbox.me/123", "WAV"}
	mm.Add(m0)


	if mm.Len() != 2 {
		t.Error("MusicManager.Add() failed.")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist ||
		m.Name != m0.Name || m.Source != m0.Source ||
		m.Type != m0.Type {
		t.Error("MusicManager.Find() failed, Found iteam mismatch.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManger.Get() failed.", err)
	}
	m, err = mm.Get(100)
	if m != nil || err == nil {
		t.Error("MusicManger.Get failed.", err)
	} else {
		t.Logf("expected: %v", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 1 {
		t.Error("MusicManager.Remove() failed.", err)
	}
	m = mm.RemoveByName("I'm Not Affried")
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManger.RemoveByName() failed.", err)
	}
}
