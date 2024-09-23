package test

import (
	"reflect"
	"testing"

	"github.com/codescalersinternships/INI-Parser-Fatma-Ebrahim/pkg/iniparser"
)

func TestLoadFromString(t *testing.T) {
	i := iniparser.Parser{}
	t.Run("valid ini string", func(t *testing.T) {
		err := i.LoadFromString(`[owner]
	name = John Doe
	organization = Acme Widgets Inc.`)
		want := "John Doe"
		got := i.Sections["owner"]["name"]
		if err != nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}
		if want != got {
			t.Errorf("expected: %s, got: %s", want, got)
		}
	})
	t.Run("valid ini string with comments", func(t *testing.T) {
		err := i.LoadFromString(`[owner]
		#comment1
	name = John Doe
	;comment2
	organization = Acme Widgets Inc.`)
		want := "John Doe"
		got := i.Sections["owner"]["name"]
		if err != nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}
		if want != got {
			t.Errorf("expected: %s, got: %s", want, got)
		}
	})
	t.Run("invalid ini string section name with no suffix ]", func(t *testing.T) {
		err := i.LoadFromString(`[owner
		name = John Doe
		organization = Acme Widgets Inc.`)
		if err == nil {
			t.Errorf("expected: %s, got: %s", err, "nil")
		}

	})
	t.Run("invalid ini string section name with no prefix [ or suffix ]", func(t *testing.T) {
		err := i.LoadFromString(`owner
		name = John Doe
		organization = Acme Widgets Inc.`)
		if err == nil {
			t.Errorf("expected: %s, got: %s", err, "nil")
		}

	})
}

func TestLoadFromFile(t *testing.T) {
	i := iniparser.Parser{}
	t.Run("Load from existing valid file", func(t *testing.T) {
		err := i.LoadFromFile("../testdata/validtest.ini")
		if err != nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}

	})
	t.Run("Load from existing invalid file", func(t *testing.T) {
		err := i.LoadFromFile("../testdata/invalidtest.ini")
		if err == nil {
			t.Errorf("expected: %s, got: %s", err, "nil")
		}

	})
	t.Run("Load from non existing file", func(t *testing.T) {
		err := i.LoadFromFile("nottest.ini")
		if err == nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}

	})

}

func TestGetSectionNames(t *testing.T) {
	i := iniparser.Parser{}
	err := i.LoadFromFile("../testdata/validtest.ini")
	if err != nil {
		t.Fatal(err)
	}
	got := i.GetSectionNames()
	want := []string{"owner"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestGetSections(t *testing.T) {
	i := iniparser.Parser{}
	err := i.LoadFromFile("../testdata/validtest.ini")
	if err != nil {
		t.Fatal(err)
	}
	got := i.GetSections()
	want := map[string]map[string]string{"owner": {"name": "John Doe", "organization": "Acme Widgets Inc."}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestGet(t *testing.T) {
	i := iniparser.Parser{}
	err := i.LoadFromFile("../testdata/validtest.ini")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Get existing key", func(t *testing.T) {
		got, exists := i.Get("owner", "name")
		want := "John Doe"
		if want != got {
			t.Errorf("expected: %s, got: %s", want, got)
		}
		if !exists {
			t.Errorf("expected key to exist")
		}

	})
	t.Run("Get non existing section", func(t *testing.T) {
		_, exists := i.Get("database", "name")
		if exists {
			t.Errorf("expected key to not exist")
		}

	})
	t.Run("Get non existing key", func(t *testing.T) {
		_, exists := i.Get("owner", "age")
		if exists {
			t.Errorf("expected key to not exist")
		}
	})
}

func TestSet(t *testing.T) {
	i := iniparser.Parser{}
	err := i.LoadFromFile("../testdata/validtest.ini")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Set a key to existing section", func(t *testing.T) {
		i.Set("owner", "name", "Alex Doe")
		if err != nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}
		got, exists := i.Get("owner", "name")
		want := "Alex Doe"
		if want != got {
			t.Errorf("expected: %s, got: %s", want, got)
		}
		if !exists {
			t.Errorf("expected key to exist")
		}

	})
	t.Run("Set a key to non existing section", func(t *testing.T) {
		i.Set("database", "name", "Alex Doe")
		if err != nil {
			t.Errorf("expected: %s, got: %s", "nil", err)
		}
		got, exists := i.Get("database", "name")
		want := "Alex Doe"
		if want != got {
			t.Errorf("expected: %s, got: %s", want, got)
		}
		if !exists {
			t.Errorf("expected key to exist")
		}
	})
}

func TestToString(t *testing.T) {
	i := iniparser.Parser{}
	err := i.LoadFromFile("../testdata/validtest.ini")
	if err != nil {
		t.Fatal(err)
	}
	got := i.String()
	want := `[owner]
name = John Doe
organization = Acme Widgets Inc.

`
	if got != want {
		t.Errorf("expected: %s, got: %s", want, got)
	}
}

func TestSaveToFile(t *testing.T) {
	i := iniparser.Parser{}
	err := i.LoadFromFile("../testdata/validtest.ini")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Save to correct file name", func(t *testing.T) {
		err := i.SaveToFile("../testdata/test2.ini")
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
