package library

import (
	"fmt"
	"testing"
)

func TestOps(t *testing.T)  {
	fmt.Println("library manager test.")

	mm := NewMusicManager()

	if mm == nil {
		t.Error("New MusicManager failed")
	}

	fmt.Println("New MusicManager size:", mm.Len())

	m0 := &Music{
		"1", "Hello World", "WXX", 
		"https://github.com/wxxhub/go_study", "MP3"}
	
	fmt.Println("add m0")
	mm.Add(m0)

	fmt.Println("New MusicManager size:", mm.Len())

	m := mm.Find(m0.Name)

	if m == nil {
		t.Error("MusicManager.Find() failed.")
	} else {
		fmt.Println("find", m0.Name)
	}

	m, err := mm.Get(0)

	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	} else {
		fmt.Println("success Get", m.Name)
	}

	m = mm.Remove(0)
	if m == nil {
		t.Error("MusicManager.Remove() failed.", err)
	} else {
		fmt.Println("success Remove", m.Name)
	}

	fmt.Println("New MusicManager size:", mm.Len())
}

func TestTest(t *testing.T)  {
	fmt.Println("Test test.")
}