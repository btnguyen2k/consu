package olaf

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestUnixMilliseconds(t *testing.T) {
	ms := UnixMilliseconds()
	now := time.Now()
	delta := now.UnixNano() - ms*1000000
	if delta < 0 || delta >= 1000000 {
		t.Fatalf("TestUnixMilliseconds failed, ms: %d, now: %d, delta: %d.", ms, now.UnixNano(), delta)
	}
}

func TestWaitTillNextMillisec(t *testing.T) {
	start := time.Now()
	nextMs := WaitTillNextMillisec(start.UnixNano() / 1000000)
	end := time.Now()
	startMs := start.UnixNano() / 1000000
	endMs := end.UnixNano() / 1000000
	delta := nextMs - startMs
	if 0 >= delta {
		t.Fatalf("Next milliseconds was incorrect, prevMs: %d, nextMs: %d, delta: %d", startMs, nextMs, delta)
	}
	if endMs < nextMs {
		t.Fatalf("Next milliseconds must not greater than now, nextMs: %d, nowMs: %d.", nextMs, endMs)
	}
}

func TestNewOlaf(t *testing.T) {
	nodeId := int64(1981)
	o := NewOlaf(nodeId)
	if o.nodeId != nodeId {
		t.Fatalf("Invalid Olaf instance, expected NodeId: %d, actual NodeId: %d.", nodeId, o.nodeId)
	}
	if o.epoch != Epoch {
		t.Fatalf("Invalid Olaf instance, expected epoch: %d, actual epoch: %d.", Epoch, o.epoch)
	}
}

func TestNewOlafWithEpoch(t *testing.T) {
	nodeId := int64(1981)
	epoch := int64(123456789)
	o := NewOlafWithEpoch(nodeId, epoch)
	if o.nodeId != nodeId {
		t.Fatalf("Invalid Olaf instance, expected NodeId: %d, actual NodeId: %d.", nodeId, o.nodeId)
	}
	if o.epoch != epoch {
		t.Fatalf("Invalid Olaf instance, expected epoch: %d, actual epoch: %d.", epoch, o.epoch)
	}
}

func _numItems() int {
	numItems, err := strconv.Atoi(os.Getenv("OLAF_NUM_ITEMS"))
	if err != nil || numItems < 1000 {
		numItems = 1000
	}
	fmt.Println("\tOlaf num items:", numItems)
	return numItems
}

func _numThreads() int {
	numThreads, err := strconv.Atoi(os.Getenv("OLAF_NUM_THREADS"))
	if err != nil || numThreads < 1 {
		numThreads = 1
	}
	fmt.Println("\tOlaf num threads:", numThreads)
	return numThreads
}

func TestOlaf_Id64(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]uint64, numItems, numItems)
	tstart := time.Now()
	for i := 0; i < numItems; i++ {
		items[i] = o.Id64()
	}
	for i := 1; i < numItems; i++ {
		if items[i-1] >= items[i] {
			t.Fatalf("Generated ID is invalid: id[%d]=%d must be less than id[%d]=%d.", i-1, items[i-1], i, items[i])
		}
	}
	d := time.Now().UnixNano() - tstart.UnixNano()
	rate := float64(numItems) / float64(d) * 1e9
	fmt.Printf("\t[INFO] generated %d IDs in %d ms (rate: %.2f/sec)\n", numItems, d/1e6, rate)
}

func TestOlaf_Id64MultiThreads(t *testing.T) {
	numThreads := _numThreads()
	numItems := _numItems()
	numItemsPerThread := numItems / numThreads
	numItems = numItemsPerThread * numThreads
	var wg sync.WaitGroup
	wg.Add(numThreads)
	items := make([]uint64, numItems, numItems)
	o := NewOlaf(1981)
	tstart := time.Now()
	for i := 0; i < numThreads; i++ {
		go func(threadIndex int, wg *sync.WaitGroup) {
			startIndex := threadIndex * numItemsPerThread
			for i := 0; i < numItemsPerThread; i++ {
				items[startIndex+i] = o.Id64()
			}
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
	itemsMap := make(map[uint64]bool)
	for _, v := range items {
		itemsMap[v] = true
		if v == 0 {
			t.Fatalf("Invalid ID: %d", v)
		}
	}
	if len(items) != numItems || len(itemsMap) != numItems {
		t.Fatalf("Expected %d but generated %d (%d unique)", numItems, len(items), len(itemsMap))
	}
	d := time.Now().UnixNano() - tstart.UnixNano()
	rate := float64(numItems) / float64(d) * 1e9
	fmt.Printf("\t[INFO] generated %d IDs in %d ms (rate: %.2f/sec)\n", numItems, d/1e6, rate)
}

func TestOlaf_Id64Hex(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]string, numItems, numItems)
	tstart := time.Now()
	for i := 0; i < numItems; i++ {
		items[i] = o.Id64Hex()
	}
	for i := 1; i < numItems; i++ {
		if items[i-1] >= items[i] {
			t.Fatalf("Generated ID is invalid: id[%d]=%s must be less than id[%d]=%s.", i-1, items[i-1], i, items[i])
		}
	}
	d := time.Now().UnixNano() - tstart.UnixNano()
	rate := float64(numItems) / float64(d) * 1e9
	fmt.Printf("\t[INFO] generated %d IDs in %d ms (rate: %.2f/sec)\n", numItems, d/1e6, rate)
}

func TestOlaf_Id64Ascii(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]string, numItems, numItems)
	tstart := time.Now()
	for i := 0; i < numItems; i++ {
		items[i] = o.Id64Ascii()
	}
	for i := 1; i < numItems; i++ {
		if items[i-1] >= items[i] {
			t.Fatalf("Generated ID is invalid: id[%d]=%s must be less than id[%d]=%s.", i-1, items[i-1], i, items[i])
		}
	}
	d := time.Now().UnixNano() - tstart.UnixNano()
	rate := float64(numItems) / float64(d) * 1e9
	fmt.Printf("\t[INFO] generated %d IDs in %d ms (rate: %.2f/sec)\n", numItems, d/1e6, rate)
}

func TestOlaf_Id128(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]*big.Int, numItems, numItems)
	tstart := time.Now()
	for i := 0; i < numItems; i++ {
		items[i] = o.Id128()
	}
	for i := 1; i < numItems; i++ {
		if items[i].Cmp(items[i-1]) <= 0 {
			t.Fatalf("Generated ID is invalid: id[%d]=%s must be less than id[%d]=%s.", i-1, items[i-1].String(), i, items[i].String())
		}
	}
	d := time.Now().UnixNano() - tstart.UnixNano()
	rate := float64(numItems) / float64(d) * 1e9
	fmt.Printf("\t[INFO] generated %d IDs in %d ms (rate: %.2f/sec)\n", numItems, d/1e6, rate)
}

func TestOlaf_Id128MultiThreads(t *testing.T) {
	numThreads := _numThreads()
	numItems := _numItems()
	numItemsPerThread := numItems / numThreads
	numItems = numItemsPerThread * numThreads
	var wg sync.WaitGroup
	wg.Add(numThreads)
	items := make([]*big.Int, numItems, numItems)
	o := NewOlaf(1981)
	tstart := time.Now()
	for i := 0; i < numThreads; i++ {
		go func(threadIndex int, wg *sync.WaitGroup) {
			startIndex := threadIndex * numItemsPerThread
			for i := 0; i < numItemsPerThread; i++ {
				items[startIndex+i] = o.Id128()
			}
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
	itemsMap := make(map[string]bool)
	for _, v := range items {
		itemsMap[v.String()] = true
		if v.Cmp(big.NewInt(0)) == 0 {
			t.Fatalf("Invalid ID: %d", v)
		}
	}
	if len(items) != numItems || len(itemsMap) != numItems {
		t.Fatalf("Expected %d but generated %d (%d unique)", numItems, len(items), len(itemsMap))
	}
	d := time.Now().UnixNano() - tstart.UnixNano()
	rate := float64(numItems) / float64(d) * 1e9
	fmt.Printf("\t[INFO] generated %d IDs in %d ms (rate: %.2f/sec)\n", numItems, d/1e6, rate)
}

func TestOlaf_Id128Hex(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]string, numItems, numItems)
	tstart := time.Now()
	for i := 0; i < numItems; i++ {
		items[i] = o.Id128Hex()
	}
	for i := 1; i < numItems; i++ {
		if items[i-1] >= items[i] {
			t.Fatalf("Generated ID is invalid: id[%d]=%s must be less than id[%d]=%s.", i-1, items[i-1], i, items[i])
		}
	}
	d := time.Now().UnixNano() - tstart.UnixNano()
	rate := float64(numItems) / float64(d) * 1e9
	fmt.Printf("\t[INFO] generated %d IDs in %d ms (rate: %.2f/sec)\n", numItems, d/1e6, rate)
}

func TestOlaf_Id128Ascii(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]string, numItems, numItems)
	tstart := time.Now()
	for i := 0; i < numItems; i++ {
		items[i] = o.Id128Ascii()
	}
	for i := 1; i < numItems; i++ {
		if items[i-1] >= items[i] {
			t.Fatalf("Generated ID is invalid: id[%d]=%s must be less than id[%d]=%s.", i-1, items[i-1], i, items[i])
		}
	}
	d := time.Now().UnixNano() - tstart.UnixNano()
	rate := float64(numItems) / float64(d) * 1e9
	fmt.Printf("\t[INFO] generated %d IDs in %d ms (rate: %.2f/sec)\n", numItems, d/1e6, rate)
}

func TestOlaf_ExtractTime64(t *testing.T) {
	o := NewOlaf(1981)
	now := time.Now()
	id64 := o.Id64()
	time := o.ExtractTime64(id64)
	delta := time.UnixNano() - now.UnixNano()
	if delta > 1000000 || delta < -1000000 {
		// delta must be within 1ms
		t.Fatalf("Invalid extracted time %s (delta nano: %d).", time, delta)
	}
}

func TestOlaf_ExtractTime64Hex(t *testing.T) {
	o := NewOlaf(1981)
	now := time.Now()
	id64Hex := o.Id64Hex()
	time := o.ExtractTime64Hex(id64Hex)
	delta := time.UnixNano() - now.UnixNano()
	if delta > 1000000 || delta < -1000000 {
		// delta must be within 1ms
		t.Fatalf("Invalid extracted time %s (delta nano: %d).", time, delta)
	}
}

func TestOlaf_ExtractTime64Ascii(t *testing.T) {
	o := NewOlaf(1981)
	now := time.Now()
	id64Ascii := o.Id64Ascii()
	time := o.ExtractTime64Ascii(id64Ascii)
	delta := time.UnixNano() - now.UnixNano()
	if delta > 1000000 || delta < -1000000 {
		// delta must be within 1ms
		t.Fatalf("Invalid extracted time %s (delta nano: %d).", time, delta)
	}
}

func TestOlaf_ExtractTime128(t *testing.T) {
	o := NewOlaf(1981)
	now := time.Now()
	id128 := o.Id128()
	time := o.ExtractTime128(id128)
	delta := time.UnixNano() - now.UnixNano()
	if delta > 1000000 || delta < -1000000 {
		// delta must be within 1ms
		t.Fatalf("Invalid extracted time %s (delta nano: %d).", time, delta)
	}
}

func TestOlaf_ExtractTime128Hex(t *testing.T) {
	o := NewOlaf(1981)
	now := time.Now()
	id128Hex := o.Id128Hex()
	time := o.ExtractTime128Hex(id128Hex)
	delta := time.UnixNano() - now.UnixNano()
	if delta > 1000000 || delta < -1000000 {
		// delta must be within 1ms
		t.Fatalf("Invalid extracted time %s (delta nano: %d).", time, delta)
	}
}

func TestOlaf_ExtractTime128Ascii(t *testing.T) {
	o := NewOlaf(1981)
	now := time.Now()
	id128Ascii := o.Id128Ascii()
	time := o.ExtractTime128Ascii(id128Ascii)
	delta := time.UnixNano() - now.UnixNano()
	if delta > 1000000 || delta < -1000000 {
		// delta must be within 1ms
		t.Fatalf("Invalid extracted time %s (delta nano: %d).", time, delta)
	}
}
