# PK Scramble Eggs
## CLI Tool for Date-Based Character Shifting

This CLI tool takes a date, transforms it into a float in the format `0.MONTH-DAY-YEAR`, and uses a natural constant (such as π) to perform character shifting on a given string. This can be used to encode secrets in plain text.

## Features

- Accepts a date input.
- Converts the date to a float in the format `0.MONTH-DAY-YEAR`.
- Divides a natural constant you choose by the date float.
- Stores all digits up to a specified length.
- Uses the digits to shift characters in a string.
- this will be used to hide things in plain sight. ( only the person who knows it will be able to decode.)

## Usage

```sh
# Example usage
#  -d, --date string    A date flag (MM-DD-YYYY) (default "01-03-2009")
#  -s, --string string  A string flag (default "default")
#  -i, --constant int   An integer flag (0-5)
go run cmd/main.go --date 2023-10-05 --constant 1 --string "HelloWorld"
```

## Example

Given the date `2023-10-05`:
1. Convert to float: `0.10-05-2023`
2. Divide π by the float value.
3. Extract digits up to the specified length.
4. Shift characters in the string "HelloWorld" using the extracted digits.

## Example output:

``` bash
➜  PK_scrambled_eggs git:(main) ✗ make run
→ Running the command line executable...
Integer: 0
String: default
Date: 01-03-2009
Custom constant: 304.415237813391751341467225577

Scramble -> d -> : g , 100 + 3 = 103 
Scramble -> e -> : e , 101 + 0 = 101 
Scramble -> f -> : j , 102 + 4 = 106 
Scramble -> a -> : e , 097 + 4 = 101 
Scramble -> u -> : v , 117 + 1 = 118 
Scramble -> l -> : q , 108 + 5 = 113 
Scramble -> t -> : v , 116 + 2 = 118 
Scrambled string "default" to "gejevqv"

Unscramble -> d -> : a , 100 - 3 = 97 
Unscramble -> e -> : e , 101 - 0 = 101 
Unscramble -> f -> : b , 102 - 4 = 98 
Unscramble -> a -> : ] , 097 - 4 = 93 
Unscramble -> u -> : t , 117 - 1 = 116 
Unscramble -> l -> : g , 108 - 5 = 103 
Unscramble -> t -> : r , 116 - 2 = 114 
Unscrambled string "default" to "aeb]tgr"

➜  PK_scrambled_eggs git:(main) ✗ go run cmd/main.go --string gejevqv
Integer: 0
String: gejevqv
Date: 01-03-2009
Custom constant: 304.415237813391751341467225577

Scramble -> g -> : j , 103 + 3 = 106 
Scramble -> e -> : e , 101 + 0 = 101 
Scramble -> j -> : n , 106 + 4 = 110 
Scramble -> e -> : i , 101 + 4 = 105 
Scramble -> v -> : w , 118 + 1 = 119 
Scramble -> q -> : v , 113 + 5 = 118 
Scramble -> v -> : x , 118 + 2 = 120 
Scrambled string "gejevqv" to "jeniwvx"

Unscramble -> g -> : d , 103 - 3 = 100 
Unscramble -> e -> : e , 101 - 0 = 101 
Unscramble -> j -> : f , 106 - 4 = 102 
Unscramble -> e -> : a , 101 - 4 = 97 
Unscramble -> v -> : u , 118 - 1 = 117 
Unscramble -> q -> : l , 113 - 5 = 108 
Unscramble -> v -> : t , 118 - 2 = 116 
Unscrambled string "gejevqv" to "default"

```

## License

This project is licensed under the MIT License.