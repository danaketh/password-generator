# Password Generator
Simple program for securely generating passwords whenever you're in need of that
and have for any reason can't use a password manager.

## Requirements
* Go 1.22+

## Options
- `--length` - set the length of the password
- `--number` - number of passwords to generate
- `--upper` - include uppercase letters
- `--lower` - include lowercase letters
- `--digits` - include digits
- `--special` - include special characters

## Defaults
* `--length` is `16` characters
* `--number` is `1`
* `--upper` is `true`
* `--lower` is `true`
* `--digits` is `true`
* `--special` is `true`

So, if all you're looking for is somewhat secure password, you can just run the program as is, no options required.
