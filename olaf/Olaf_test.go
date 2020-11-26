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
	if o.NodeID != nodeId {
		t.Fatalf("Invalid Olaf instance, expected NodeId: %d, actual NodeId: %d.", nodeId, o.NodeID)
	}
	if o.Epoch != Epoch {
		t.Fatalf("Invalid Olaf instance, expected Epoch: %d, actual Epoch: %d.", Epoch, o.Epoch)
	}
}

func TestNewOlafWithEpoch(t *testing.T) {
	nodeId := int64(1981)
	epoch := int64(123456789)
	o := NewOlafWithEpoch(nodeId, epoch)
	if o.NodeID != nodeId {
		t.Fatalf("Invalid Olaf instance, expected NodeId: %d, actual NodeId: %d.", nodeId, o.NodeID)
	}
	if o.Epoch != epoch {
		t.Fatalf("Invalid Olaf instance, expected Epoch: %d, actual Epoch: %d.", epoch, o.Epoch)
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
	items := make([]uint64, numItems)
	for i := 0; i < numItems; i++ {
		id := o.Id64()
		items[i] = id
		if i > 0 && items[i] <= items[i-1] {
			t.Fatalf("Generated ID is invalid: items[%d]=%d must be less than items[%d]=%d.", i-1, items[i-1], i, items[i])
		}
	}
}

func TestOlaf_Id64MultiThreads(t *testing.T) {
	numThreads := _numThreads()
	numItems := _numItems()
	numItemsPerThread := numItems / numThreads
	numItems = numItemsPerThread * numThreads
	var wg sync.WaitGroup
	wg.Add(numThreads)
	items := make([]uint64, numItems)
	o := NewOlaf(1981)
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
}

func TestOlaf_Id64Hex(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]string, numItems)
	for i := 0; i < numItems; i++ {
		id := o.Id64Hex()
		items[i] = id
		if i > 0 && items[i] <= items[i-1] {
			t.Fatalf("Generated ID is invalid: items[%d]=%s must be less than items[%d]=%s.", i-1, items[i-1], i, items[i])
		}
	}
}

func TestOlaf_Id64Ascii(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]string, numItems)
	for i := 0; i < numItems; i++ {
		id := o.Id64Ascii()
		items[i] = id
		if i > 0 && items[i] <= items[i-1] {
			t.Fatalf("Generated ID is invalid: items[%d]=%s must be less than items[%d]=%s.", i-1, items[i-1], i, items[i])
		}
	}
}

func TestOlaf_Id128(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]*big.Int, numItems)
	for i := 0; i < numItems; i++ {
		id := o.Id128()
		items[i] = id
		if i > 0 && items[i].Cmp(items[i-1]) <= 0 {
			t.Fatalf("Generated ID is invalid: items[%d]=%s must be less than items[%d]=%s.", i-1, items[i-1].String(), i, items[i].String())
		}
	}
}

func TestOlaf_Id128MultiThreads(t *testing.T) {
	numThreads := _numThreads()
	numItems := _numItems()
	numItemsPerThread := numItems / numThreads
	numItems = numItemsPerThread * numThreads
	var wg sync.WaitGroup
	wg.Add(numThreads)
	items := make([]*big.Int, numItems)
	o := NewOlaf(1981)
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
}

func TestOlaf_Id128Hex(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]string, numItems)
	for i := 0; i < numItems; i++ {
		id := o.Id128Hex()
		items[i] = id
		if i > 0 && items[i] <= items[i-1] {
			t.Fatalf("Generated ID is invalid: items[%d]=%s must be less than items[%d]=%s.", i-1, items[i-1], i, items[i])
		}
	}
}

func TestOlaf_Id128Ascii(t *testing.T) {
	o := NewOlaf(1981)
	numItems := _numItems()
	items := make([]string, numItems)
	for i := 0; i < numItems; i++ {
		id := o.Id128Ascii()
		items[i] = id
		if i > 0 && items[i] <= items[i-1] {
			t.Fatalf("Generated ID is invalid: items[%d]=%s must be less than items[%d]=%s.", i-1, items[i-1], i, items[i])
		}
	}
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
