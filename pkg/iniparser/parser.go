package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	ErrSectionNotExists = errors.New("Section doesn't exist")
	ErrKeyNotExists     = errors.New("cKey doesn't exist")
)

type Iniparser struct {
	sections map[string]map[string]string
}

// a method to load the data from the a string that holds the ini structure into the iniparser
func (i *Iniparser) LoadFromString(config string) {
	i.sections = map[string]map[string]string{}
	lines := strings.Split(config, "\n")
	section := ""
	m := map[string]string{}

	for _, line := range lines {
		line := strings.TrimSpace(line)
		if len(line) != 0 {
			if line[0] == '[' {
				section = strings.Replace(line, "[", "", 1)
				section = strings.Replace(section, "]", "", 1)
				m = make(map[string]string)
			} else if line[0] != '#' && line[0] != ';' {
				pair := strings.Split(line, " = ")
				m[pair[0]] = pair[1]
				i.sections[section] = m
			}
		}

	}
}

// a method to load the data from an ini file into the iniparser
// it return an error if wrong filename is passed
func (i *Iniparser) LoadFromFile(filename string) error {
	config, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("can't read data from file: %s", filename)
	}
	lines := string(config)
	i.LoadFromString(lines)
	return nil
}

// a method that return the sections names as a slice of string
func (i *Iniparser) GetSectionNames() []string {
	var sectionnames []string
	for section := range i.sections {
		sectionnames = append(sectionnames, section)
	}
	sort.Strings(sectionnames)
	return sectionnames

}

// a method that return the sections names as a map whose key is the section name
// and the value as a map of keys and values of type string
func (i *Iniparser) GetSections() map[string]map[string]string {
	return i.sections
}

// a method that recieves section name and key then returns the corresponding value
// it returns an error if the section name or the key doesn't exist
func (i *Iniparser) Get(sectionname, key string) (string, error) {
	section, exist := i.sections[sectionname]
	if exist {
		value, exist := section[key]
		if exist {
			return value, nil
		}
		return "", ErrKeyNotExists
	}
	return "", ErrSectionNotExists
}

// a method that recieves section name, key and value then add the key and value to this section
// it returns an error if the section name doesn't exist
func (i *Iniparser) Set(sectionname, key, value string) error {
	section, exist := i.sections[sectionname]
	if exist {
		section[key] = value
		return nil
	}
	return ErrSectionNotExists
}

// a method that returns the data stored in the iniparser as a string of sections, keys and values
func (i *Iniparser) ToString() string {
	lines := ""
	sections := i.GetSectionNames()
	for _, section := range sections {
		lines += ("[" + section + "]\n")

		m := i.sections[section]
		keys := []string{}
		for key := range m {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			lines += (key + " = " + m[key] + "\n")
		}
		lines += "\n"
	}
	return lines

}

// a method that saves the data stored in the iniparser in an ini file giving its name
// it creates a file with the given name and returns an error if writing to the file fails
func (i *Iniparser) SaveToFile(filename string) error {
	config := i.ToString()
	err := os.WriteFile(filename, []byte(config), 0666)
	if err != nil {
		return fmt.Errorf("can't save to file with name: %s due to error: %w", filename, err)
	}
	return nil

}
