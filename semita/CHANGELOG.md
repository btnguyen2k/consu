# consu/semita changelog

## 2021-01-13 - v0.1.5

- `PathSeparator` is now configurable (default value is `.`).

## 2019-04-12 - v0.1.4.1

- Upgrade to `consu/reddo-v0.1.6`:
    - Return `zero` value when input is `nil`.

## 2019-04-04 - v0.1.4

- Migrate to Go modular design.

## 2019-03-07 - v0.1.2

- Upgrade to `consu/reddo-v0.1.3`:
    - New functions `GetTime(path string) (time.Time, error)` and `GetTimeWithLayout(path, layout string) (time.Time, error)`

## 2019-03-05 - v0.1.1

- Compatible with `consu/reddo-v0.1.2`

## 2019-02-22 - v0.1.0

First release:
- Supported nested arrays, slices, maps and structs.
- Struct's un-exported fields can be read, but not written.
- Unaddressable structs and arrays are read-only.
