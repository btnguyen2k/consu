# consu/checksum changelog

## 2023-08-14 - v1.0.0

### Changed

- Changed algorithm for calculating map and struct checksum

### Fixed/Improvement

- Fix: value is not unwrapped correctly sometimes
- Fix: wrongly calculate checksum for time.Time

## 2020-11-20 - v0.1.2

If a struct is `time.Time`, use its nanosecond to calculate checksum.

## 2019-10-26 - v0.1.1

If a struct has function `Checksum()`, use it to calculate checksum instead of reflecting through struct's fields.

## 2019-10-17 - v0.1.0

First release:
- Calculate checksum of scalar types (`bool`, `int*`, `uint*`, `float*`, `string`) as well as `slice/array` and `map/struct`.
- Supported hash functions: `CRC32`, `MD5`, `SHA1`, `SHA256`, `SHA512`.
