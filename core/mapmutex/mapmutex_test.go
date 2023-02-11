package mapmutex_test

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/mapmutex"
)

// ExampleMapMutex provides an example implementation of a map mutex.
func ExampleStringMapMutex() {
	mapMutex := mapmutex.NewStringMapMutex()
	lock1 := mapMutex.Lock("lock1")
	lock2 := mapMutex.Lock("lock2")

	lock1.Unlock()
	lock2.Unlock()
}

func ExampleStringerMapMutex() {
	vitalik := common.HexToAddress("0xab5801a7d398351b8be11c439e05c5b3259aec9b")
	tether := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")

	mapMutex := mapmutex.NewStringerMapMutex()
	// second variable will be true
	vitalikLock, _ := mapMutex.TryLock(vitalik)
	tetherLock := mapMutex.Lock(tether)

	vitalikLock.Unlock()
	tetherLock.Unlock()
}

func ExampleIntMapMutex() {
	mapMutex := mapmutex.NewIntMapMutex()
	// second variable will be true
	lock0, _ := mapMutex.TryLock(0)
	lock1 := mapMutex.Lock(1)

	lock0.Unlock()
	lock1.Unlock()
}

func (s MapMutexSuite) TestExampleMapMutex() {
	NotPanics(s.T(), ExampleStringMapMutex)
	NotPanics(s.T(), ExampleStringerMapMutex)
	NotPanics(s.T(), ExampleStringMapMutex)
}

func (s MapMutexSuite) TestMapMutex() {
	//nolint:gosec
	r := rand.New(rand.NewSource(42))

	m := mapmutex.NewTestMapMutex(s.T())
	_ = m

	keyCount := 20
	iCount := 10000
	out := make(chan string, iCount*2)

	// run a bunch of concurrent requests for various keys,
	// the idea is to have a lot of lock contention
	var wg sync.WaitGroup
	wg.Add(iCount)
	for i := 0; i < iCount; i++ {
		go func(rn int) {
			defer wg.Done()
			key := strconv.Itoa(rn)

			// you can prove the test works by commenting the locking out and seeing it fail
			l := m.Lock(key)
			_, obtained := m.TryLock(key)
			False(s.T(), obtained)
			defer l.Unlock()

			out <- key + " A"
			time.Sleep(time.Microsecond) // make 'em wait a mo'
			out <- key + " B"
		}(r.Intn(keyCount))
	}
	wg.Wait()
	close(out)

	// verify the map is empty now
	if l := len(m.GetMa()); l != 0 {
		s.T().Errorf("unexpected map length at test end: %v", l)
	}

	// confirm that the output always produced the correct sequence
	outLists := make([][]string, keyCount)
	for so := range out {
		sParts := strings.Fields(so)
		kn, err := strconv.Atoi(sParts[0])
		if err != nil {
			s.T().Fatal(err)
		}
		outLists[kn] = append(outLists[kn], sParts[1])
	}
	for kn := 0; kn < keyCount; kn++ {
		l := outLists[kn] // list of output for this particular key
		for i := 0; i < len(l); i += 2 {
			if l[i] != "A" || l[i+1] != "B" {
				s.T().Errorf("For key=%v and i=%v got unexpected values %v and %v", kn, i, l[i], l[i+1])
				break
			}
		}
	}
	if s.T().Failed() {
		s.T().Logf("Failed, outLists: %#v", outLists)
	}
}

func BenchmarkMapMutex(b *testing.B) {
	m := mapmutex.NewTestMapMutex(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// run uncontended lock/unlock - should be quite fast
		m.Lock(i).Unlock()
	}
}
