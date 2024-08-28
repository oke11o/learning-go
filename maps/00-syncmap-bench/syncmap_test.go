package main

import (
	"strconv"
	"sync"
	"testing"
)

func benchmarkSyncMap(b *testing.B, readRatio int, writeRatio int) {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		m.Store(strconv.Itoa(i), i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(readRatio + writeRatio)

		for j := 0; j < readRatio; j++ {
			go func() {
				defer wg.Done()
				m.Load("1")
			}()
		}

		for j := 0; j < writeRatio; j++ {
			go func() {
				defer wg.Done()
				m.Store("1", 1)
			}()
		}

		wg.Wait()
	}
}

func benchmarkMapWithMutex(b *testing.B, readRatio int, writeRatio int) {
	m := make(map[string]int)
	var mu sync.RWMutex
	var wg sync.WaitGroup

	mu.Lock()
	for i := 0; i < 100; i++ {
		m[strconv.Itoa(i)] = i
	}
	mu.Unlock()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(readRatio + writeRatio)

		for j := 0; j < readRatio; j++ {
			go func() {
				defer wg.Done()
				mu.RLock()
				_ = m["1"]
				mu.RUnlock()
			}()
		}

		for j := 0; j < writeRatio; j++ {
			go func() {
				defer wg.Done()
				mu.Lock()
				m["1"] = 1
				mu.Unlock()
			}()
		}

		wg.Wait()
	}
}

func BenchmarkRead1000Write1SyncMap(b *testing.B)      { benchmarkSyncMap(b, 1000, 1) }
func BenchmarkRead1000Write1MapWithMutex(b *testing.B) { benchmarkMapWithMutex(b, 1000, 1) }

func BenchmarkReadWriteOneSyncMap(b *testing.B)      { benchmarkSyncMap(b, 1, 1) }
func BenchmarkReadWriteOneMapWithMutex(b *testing.B) { benchmarkMapWithMutex(b, 1, 1) }

func BenchmarkRead1Write1000SyncMap(b *testing.B)      { benchmarkSyncMap(b, 1, 1000) }
func BenchmarkRead1Write1000MapWithMutex(b *testing.B) { benchmarkMapWithMutex(b, 1, 1000) }
