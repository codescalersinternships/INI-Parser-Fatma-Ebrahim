package iniparser

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	ErrSectionNotExists = errors.New("section doesn't exist")
	ErrKeyNotExists     = errors.New("key doesn't exist")
)

type Parser struct {
	Sections map[string]map[string]string
}

// LoadFromString loads the data from the a string that holds the ini structure into the Parser
func (i *Parser) LoadFromString(config string) error {
	i.Sections = make(map[string]map[string]string)
	lines := strings.Split(config, "\n")
	section := ""
	m := make(map[string]string)

	for _, line := range lines {
		line := strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		if !strings.HasPrefix(line, "[") && !strings.HasSuffix(line, "]") {
			pair := strings.Split(line, " = ")
			if len(pair) != 2 {
				return fmt.Errorf("invalid line: %s", line)
			}
			m[pair[0]] = pair[1]
			i.Sections[section] = m
		} else if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			section = line[1 : len(line)-1]
			m = make(map[string]string)
		} else {
			return fmt.Errorf("invalid line: %s", line)
		}

	}
	return nil
}

// LoadFromFile loads the data from an ini file into the Parser
// it return an error if wrong filename is passed
func (i *Parser) LoadFromFile(filename string) error {
	config, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("err: %w", err)
	}
	lines := string(config)
	err = i.LoadFromString(lines)
	return err
}

// GetSectionNames returns the sections names as a slice of string
func (i *Parser) GetSectionNames() []string {
	var sectionNames []string
	for section := range i.Sections {
		sectionNames = append(sectionNames, section)
	}
	sort.Strings(sectionNames)
	return sectionNames

}

// GetSections returns the sections names as a map whose key is the section name
// and the value as a map of keys and values of type string
func (i *Parser) GetSections() map[string]map[string]string {
	return i.Sections
}

// Get recieves section name and key then returns the corresponding value
// it returns an error if the section name or the key doesn't exist
func (i *Parser) Get(sectionName, key string) (string, bool) {
	val, exists := i.Sections[sectionName][key]
	return val, exists
}

// Set recieves section name, key and value then add the key and value to this section
// it returns an error if the section name doesn't exist
func (i *Parser) Set(sectionName, key, value string) {
	section, exist := i.Sections[sectionName]
	if exist {
		section[key] = value

	} else {
		m := make(map[string]string)
		m[key] = value
		i.Sections[sectionName] = m
	}

}

// ToString returns the data stored in the Parser as a string of sections, keys and values
func (i *Parser) String() string {
	builder := strings.Builder{}
	sections := i.GetSectionNames()
	for _, section := range sections {
		builder.WriteString(fmt.Sprintf("[%s]\n", section))
		m := i.Sections[section]
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			builder.WriteString(fmt.Sprintf("%s = %s\n", key, m[key]))
		}
		builder.WriteString("\n")
	}
	return builder.String()

}

// SaveToFile saves the data stored in the Parser in an ini file giving its name
// it creates a file with the given name and returns an error if writing to the file fails
func (i *Parser) SaveToFile(filename string) error {
	config := i.String()
	err := os.WriteFile(filename, []byte(config), 0666)
	if err != nil {
		return fmt.Errorf("can't save to file with name: %s due to error: %w", filename, err)
	}
	return nil

}
