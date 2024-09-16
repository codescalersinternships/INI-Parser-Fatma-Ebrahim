# INI-Parser-Fatma-Ebrahim

This repository contains an implementation of INI parser package that provides functionality to load, retrieve and modify data from `.ini` configuration files. INI files are commonly used for configuration purposes where data is organized into sections, keys, and values.

## Features

1. **Load Configuration from String**: Load data directly from a string containing the INI structure.
2. **Load Configuration from File**: Load data from an `.ini` file.
3. **Get Section Names**: Get a list of all sections present in the configuration.
4. **Get All Sections**: Get all sections with their keys and values.
5. **Get Key-Value**: Get a specific value based on the section and key.
6. **Set Key-Value**: Add or update a key-value pair in a specific section.
7. **Save to File**: Save the current state of the INI parser to a file.
8. **ToString**: Convert the current state of the INI parser to a string format that follows the INI file structure.


### `Iniparser` Struct

- `sections`: A map representing sections of the INI file. Each section maps to another map of key-value pairs.

## Functions

### `LoadFromString(config string)`
Loads the INI structure from a string and populates the `Iniparser` object.

### `LoadFromFile(filename string) error`
Reads an `.ini` file and loads the data into the `Iniparser` object. Returns an error if the file cannot be read.

### `GetSectionNames() []string`
Returns a list of all section names present in the INI file.

### `GetSections() map[string]map[string]string`
Returns a map containing all sections and their respective key-value pairs.

### `Get(sectionname, key string) (string, error)`
Returns a value based on a given section and key. Returns an error if the section or key doesn't exist.

### `Set(sectionname, key, value string) error`
Sets or updates a key-value pair for a given section. Returns an error if the section doesn't exist.

### `ToString() string`
Returns a string representation of the INI structure.

### `SaveToFile(filename string) error`
Saves the INI structure to a file. Returns an error if file writing fails.

## How to Use:
  ### Step 1: Install the Package Using `go get`

  ```bash
  go get github.com/codescalersinternships/INI-Parser-Fatma-Ebrahim
  ```

  This command fetches the package and adds it to your project's `go.mod` file.

  ### Step 2: Import and Use the Package in Your Code

  After running `go get`, you can import the package into your project and use the functions as described:

  ```go
  package main

  import (
      "github.com/codescalersinternships/INI-Parser-Fatma-Ebrahim" 
  )

  func main() {
      i := iniparser.Iniparser{}
     //Now you can use all the methods in the package
  }
  ```
