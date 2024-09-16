package main

import (
	"reflect"
	"testing"
)

func TestLoadFromString(t *testing.T) {
	i := Iniparser{}
	i.LoadFromString(`[owner]
name = John Doe
organization = Acme Widgets Inc.`)
	want := "John Doe"
	got := i.sections["owner"]["name"]

	if want != got {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestLoadFromFile(t *testing.T) {

	t.Run("Load from existing file", func(t *testing.T) {
		i := Iniparser{}
		err := i.LoadFromFile("test.ini")
		if err != nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}

	})
	t.Run("Load from non existing file", func(t *testing.T) {
		i := Iniparser{}
		err := i.LoadFromFile("nottest.ini")
		if err == nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}

	})

}

func TestGetSectionNames(t *testing.T) {
	i := Iniparser{}
	i.LoadFromFile("test.ini")
	got := i.GetSectionNames()
	want := []string{"owner"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestGetSections(t *testing.T) {
	i := Iniparser{}
	i.LoadFromFile("test.ini")
	got := i.GetSections()
	want := map[string]map[string]string{"owner": {"name": "John Doe", "organization": "Acme Widgets Inc."}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestGet(t *testing.T) {
	i := Iniparser{}
	i.LoadFromFile("test.ini")
	t.Run("Get existing key", func(t *testing.T) {
		got, err := i.Get("owner", "name")
		want := "John Doe"
		if want != got {
			t.Errorf("expected: %s, got: %s", want, got)
		}
		if err != nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}

	})
	t.Run("Get non existing section", func(t *testing.T) {
		_, err := i.Get("database", "name")
		want := ErrSectionNotExists
		if err == nil {
			t.Fatal("expected an error and recieved nil")
		}
		if err != want {
			t.Errorf("expected: %s, got: %s", want, err)
		}
	})
	t.Run("Get non existing key", func(t *testing.T) {
		_, err := i.Get("owner", "age")
		want := ErrKeyNotExists
		if err == nil {
			t.Fatal("expected an error and recieved nil")
		}
		if err != want {
			t.Errorf("expected: %s, got: %s", want, err)
		}
	})
}

func TestSet(t *testing.T) {
	i := Iniparser{}
	i.LoadFromFile("test.ini")
	t.Run("Set a key to existing section", func(t *testing.T) {
		err := i.Set("owner", "name", "Alex Doe")
		if err != nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}

	})
	t.Run("Set a key to non existing section", func(t *testing.T) {
		err := i.Set("database", "name", "Alex Doe")
		want := ErrSectionNotExists
		if err == nil {
			t.Fatal("expected an error and recieved nil")
		}
		if err != want {
			t.Errorf("expected: %s, got: %s", want, err)
		}
	})
}

func TestToString(t *testing.T) {
	i := Iniparser{}
	i.LoadFromFile("test.ini")
	got := i.ToString()
	want := `[owner]
name = John Doe
organization = Acme Widgets Inc.

`
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestSaveToFile(t *testing.T) {
	i := Iniparser{}
	i.LoadFromFile("test.ini")
	t.Run("Save to correct file name", func(t *testing.T) {
		err := i.SaveToFile("test2.ini")
		if err != nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}
	})

	t.Run("Save to incorrect file name", func(t *testing.T) {
		err := i.SaveToFile("")
		if err == nil {
			t.Fatal("expected an error and recieved nil")
		}
	})

}
