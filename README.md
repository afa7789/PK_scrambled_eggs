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
cli-tool --date 2023-10-05 --constant 1 --string "HelloWorld"
```

## Example

Given the date `2023-10-05`:
1. Convert to float: `0.10-05-2023`
2. Divide π by the float value.
3. Extract digits up to the specified length.
4. Shift characters in the string "HelloWorld" using the extracted digits.

## Installation

To install the CLI tool, run:

```sh
pip install cli-tool
```

## License

This project is licensed under the MIT License.