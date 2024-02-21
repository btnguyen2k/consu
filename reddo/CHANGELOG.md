# consu/reddo release notes

## 2024-02-21 - v0.1.9

### Fixed/Improvement

- Fix(CodeQL): Off-by-one comparison against length.

## 2023-05-04 - v0.1.8

- Bug fixes & enhancements: `ToMap` & `ToSlice` handle complex/nested/nil cases better.

## 2021-01-12 - v0.1.7

- Bug fix: `ToMap` causes panic if map key or value is nil.
- Bug fix: `ToSlice` causes panic if an element is nil.
- Add `ZeroMode`:
    - If `reddo.ZeroMode=true`: `zero` value is returned when input is `nil`.
    - If `reddo.ZeroMode=false`: error is returned by `ToBool`, `ToFloat`, `ToInt`, `ToUint`, `ToString`, `ToTime`/`ToTimeWithLayout` and `ToStruct`

## 2019-04-12 - v0.1.6

- Return `zero` value when input is `nil`:
    - `reddo.ToBool(...)` returns `false`
    - `reddo.ToFloat(...)`, `reddo.ToInt(...)` and `reddo.ToUint(...)` returns `0`
    - `reddo.ToStrirng(...)` returns `""`
    - `reddo.ToTime(...)` and `reddo.ToTimeWithLayout(...)` returns `time.Time{}`
    - `reddo.ToSlice(...)` returns `nil`
    - `reddo.ToMap(...)` returns `nil`
    - `reddo.ToPointer(...)` returns `nil`

## 2019-04-02 - v0.1.5

- `reddo.ToString(...)`: handle the case converting `[]byte` to `string`.
- `reddo.ToSlice(...)`: handle the case converting `string` to `[]byte`.

## 2019-04-01 - v0.1.4

- Migrate to Go modular design.

## 2019-03-07 - v0.1.3

- New function `ToTimeWithLayout(v interface{}, layout string) (time.Time, error)`

## 2019-03-05 - v0.1.2

- Refactoring:
    - `ToBool(...)` now returns `(bool, error)`
    - `ToFloat(...)` now returns `(float64, error)`
    - `ToInt(...)` now returns `(int64, error)`
    - `ToUint(...)` now returns `(uint64, error)`
    - `ToString(...)` now returns `(string, error)`
    - `ToStruct(...)` changes its parameters to `ToStruct(v interface{}, t reflect.Type)`. Supplied target type can be slice, array or an element or array/slice.
    - `ToMap(...)` changes its parameters to `ToMap(v interface{}, t reflect.Type)`.
    - `Convert(...)` changes its parameters to `Convert(v interface{}, t reflect.Type)`.
- Remove `Zero...`, add `Type...`
- Other fixes and enhancements

## 2019-02-12 - v0.1.1.2

- New (semi)constants `ZeroMap` and `ZeroSlice`
- Fix: to solve the case "convert to `interface{}`"
    - Function `Convert(v interface{}, t interface{}) (interface{}, error)` returns `(v, nil)` if `t` is `nil`

## 2019-02-11 - v0.1.1.1

- New constant `ZeroUint64`

## 2019-01-15 - v0.1.1

- `ToStruct(interface{}, interface{}) (interface{}, error)` & new function `ToTime(interface{}) (time.Time, error)`:
    - Add special case when converting to `time.Time`
    - Add global value `ZeroTime`
    - Fix a bug when converting a unexported field

## 2019-01-12 - v0.1.0

- Convert primitive types (`bool`, `float*`, `int*`, `uint*`, `string`)
- Convert `struct`, `array/slice` and `map`
- Convert pointer
