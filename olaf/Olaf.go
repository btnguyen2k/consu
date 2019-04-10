/*
Package olaf provides methods to generate unique ID using Twitter Snowflake algorithm.

Sample usage:

	package main

	import (
		"fmt"
		"github.com/btnguyen2k/consu/olaf"
	)

	func main() {
		// use default epoch
		o := olaf.NewOlaf(1981)

		// use custom epoch (note: epoch is in milliseconds)
		// o := olaf.NewOlafWithEpoch(103, 1546543604123)

		id64 := o.Id64()
		id64Hex := o.Id64Hex()
		id64Ascii := o.Id64Ascii()
		fmt.Println("ID 64-bit (int)   : ", id64, " / Timestamp: ", o.ExtractTime64(id64))
		fmt.Println("ID 64-bit (hex)   : ", id64Hex, " / Timestamp: ", o.ExtractTime64Hex(id64Hex))
		fmt.Println("ID 64-bit (ascii) : ", id64Ascii, " / Timestamp: ", o.ExtractTime64Ascii(id64Ascii))

		id128 := o.Id128()
		id128Hex := o.Id128Hex()
		id128Ascii := o.Id128Ascii()
		fmt.Println("ID 128-bit (int)  : ", id128.String(), " / Timestamp: ", o.ExtractTime128(id128))
		fmt.Println("ID 128-bit (hex)  : ", id128Hex, " / Timestamp: ", o.ExtractTime128Hex(id128Hex))
		fmt.Println("ID 128-bit (ascii): ", id128Ascii, " / Timestamp: ", o.ExtractTime128Ascii(id128Ascii))
	}
*/
package olaf

import (
	"math/big"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// Version defines version number of this package
	Version = "0.1.1"

	// Epoch is set to 2019-01-01 00:00:00 UTC.
	// You may customize this to set a different epoch for your application.
	Epoch int64 = 1546300800000
)

const (
	maskNodeId64     = 0x3FF  // 10 bits
	maxSequenceId64  = 0x1FFF // 13 bits
	shiftNodeId64    = 13
	shiftTimestamp64 = 23

	maskNodeId128     = 0xFFFFFFFFFFFF // 48 bits
	maxSequenceId128  = 0xFFFF         // 16 bits
	shiftNodeId128    = 16
	shiftTimestamp128 = 64

	maxRadix = 36
)

// Olaf wraps configurations for Twitter Snowflake IDs.
type Olaf struct {
	NodeID     int64 // original node-id
	nodeId64   int64 // node-id  for 64-bit ids
	nodeId128  int64 // node-id  for 128-bit ids
	Epoch      int64 // Twitter snowflake's epoch
	SequenceId int64 // Twitter snowflake's sequence-id
	Timestamp  int64 // last 'touch' UNIX timestamp in milliseconds
}

// NewOlaf creates a new Olaf with default epoch.
func NewOlaf(nodeId int64) *Olaf {
	return NewOlafWithEpoch(nodeId, Epoch)
}

// NewOlafWithEpoch creates a new Olaf with custom epoch.
func NewOlafWithEpoch(nodeId int64, epoch int64) *Olaf {
	olaf := Olaf{}
	olaf.NodeID = nodeId
	olaf.nodeId64 = nodeId & maskNodeId64
	olaf.nodeId128 = nodeId & maskNodeId128
	olaf.Epoch = epoch
	olaf.SequenceId = 0
	olaf.Timestamp = 0
	return &olaf
}

/*----------------------------------------------------------------------*/

// UnixMilliseconds returns current UNIX timestamp in milliseconds.
func UnixMilliseconds() int64 {
	return time.Now().UnixNano() / 1000000
}

// WaitTillNextMillisec waits till clock moves to the next millisecond.
// Returns the "next" millisecond.
func WaitTillNextMillisec(currentMillisec int64) int64 {
	nextMillisec := UnixMilliseconds()
	for ; nextMillisec <= currentMillisec; nextMillisec = UnixMilliseconds() {
		runtime.Gosched()
	}
	return nextMillisec
}

// ExtractTime64 extracts time metadata from a 64-bit id.
func (o *Olaf) ExtractTime64(id64 uint64) time.Time {
	timestamp := id64>>shiftTimestamp64 + uint64(o.Epoch)
	sec := timestamp / 1000
	nsec := (timestamp % 1000) * 1000000
	return time.Unix(int64(sec), int64(nsec))
}

// ExtractTime64Hex extracts time metadata from a 64-bit id in hex (base 16) format.
func (o *Olaf) ExtractTime64Hex(id64Hex string) time.Time {
	id64, _ := strconv.ParseUint(id64Hex, 16, 64)
	return o.ExtractTime64(id64)
}

// ExtractTime64Ascii extracts time metadata from a 64-bit id in ascii (base 36) format.
func (o *Olaf) ExtractTime64Ascii(id64Ascii string) time.Time {
	id64, _ := strconv.ParseUint(id64Ascii, maxRadix, 64)
	return o.ExtractTime64(id64)
}

// Id64 generates a 64-bit id.
func (o *Olaf) Id64() uint64 {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	timestamp := UnixMilliseconds()
	sequence := int64(0)
	for done := false; !done; {
		done = true
		for timestamp < o.Timestamp {
			timestamp = WaitTillNextMillisec(timestamp)
		}
		if timestamp == o.Timestamp {
			// increase sequence
			sequence = atomic.AddInt64(&o.SequenceId, 1)
			if sequence > maxSequenceId64 {
				// reset sequence
				o.SequenceId = 0
				timestamp = WaitTillNextMillisec(timestamp)
				done = false
			}
		}
	}
	o.SequenceId = sequence
	o.Timestamp = timestamp
	result := ((timestamp - o.Epoch) << shiftTimestamp64) | (o.nodeId64 << shiftNodeId64) | sequence
	return uint64(result)
}

// Id64Hex generates a 64-bit id as a hex (base 16) string.
func (o *Olaf) Id64Hex() string {
	return strconv.FormatUint(o.Id64(), 16)
}

// Id64Ascii generates a 64-bit id as an ascii string (base 36).
func (o *Olaf) Id64Ascii() string {
	return strconv.FormatUint(o.Id64(), maxRadix)
}

/*----------------------------------------------------------------------*/

// ExtractTime128 extracts time metadata from a 128-bit id.
func (o *Olaf) ExtractTime128(id128 *big.Int) time.Time {
	timestamp := id128.Rsh(id128, shiftTimestamp128).Int64()
	sec := timestamp / 1000
	nsec := (timestamp % 1000) * 1000000
	return time.Unix(sec, nsec)
}

// ExtractTime128Hex extracts time metadata from a 128-bit id in hex (base 16) format.
func (o *Olaf) ExtractTime128Hex(id128Hex string) time.Time {
	id128 := big.NewInt(0)
	id128.SetString(id128Hex, 16)
	return o.ExtractTime128(id128)
}

// ExtractTime128Ascii extracts time metadata from a 128-bit id in ascii (base 36) format.
func (o *Olaf) ExtractTime128Ascii(id128Ascii string) time.Time {
	id128 := big.NewInt(0)
	id128.SetString(id128Ascii, maxRadix)
	return o.ExtractTime128(id128)
}

// Id128 generates a 128-bit id.
func (o *Olaf) Id128() *big.Int {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	timestamp := UnixMilliseconds()
	sequence := int64(0)
	for done := false; !done; {
		done = true
		for timestamp < o.Timestamp {
			timestamp = WaitTillNextMillisec(timestamp)
		}
		if timestamp == o.Timestamp {
			// increase sequence
			sequence = atomic.AddInt64(&o.SequenceId, 1)
			if sequence > maxSequenceId128 {
				// reset sequence
				o.SequenceId = 0
				timestamp = WaitTillNextMillisec(timestamp)
				done = false
			}
		}
	}
	o.SequenceId = sequence
	o.Timestamp = timestamp
	high := timestamp
	low := (o.nodeId128 << shiftNodeId128) | sequence
	h := big.NewInt(high)
	h.Lsh(h, shiftTimestamp128)
	return h.Add(h, big.NewInt(low))
}

// Id128Hex generates a 128-bit id as a hex (base 16) string.
func (o *Olaf) Id128Hex() string {
	return o.Id128().Text(16)
}

// Id128Ascii generates a 128-bit id as an ascii (base 36) string.
func (o *Olaf) Id128Ascii() string {
	return o.Id128().Text(maxRadix)
}
