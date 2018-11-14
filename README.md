genpwd
================
A small Go script to generate passwords that confirm to password complexity rules.

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