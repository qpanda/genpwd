genpwd
================
A small Go script to generate passwords that confirm to password complexity rules.

## Installation
Pre-built **windows/amd64** and **darwin/amd64** binaries can be downloaded from the [release page](https://github.com/qpanda/genpwd/releases)

To run **genpwd** or build binaries using Go download the latest source package from the [release page](https://github.com/qpanda/genpwd/releases) and extract it.
  
	$ cd genpwd-x.y.z/genpwd
	$ go run main.go
    $ go build

**Note:** **genpwd** has been developed and tested with Go 1.11.2.

## Using genpwd
Run ```genpwd``` to get usage information. All parameters have default values.

    $ genpwd -h
	Usage of ..\bin\genpwd.exe:
	  -l uint
			password length (default 16)
	  -lower uint
			minimum number of lower case characters (default 2)
	  -numeric uint
			minimum number of numeric characters (default 2)
	  -special uint
			minimum number of special characters (default 2)
	  -specialCharSet string
			special character set "all", "restricted", or "limited" (default "limited")
	  -upper uint
			minimum number of upper case characters (default 2)

## License
**genpwd** is licensed under the BSD-3-Clause license.