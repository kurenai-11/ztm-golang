//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32
type AverageResult float32

type BandwidthUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) Average() AverageResult {
	var totalBytes Bytes = 0
	for _, v := range b.amount {
		totalBytes += v
	}
	return AverageResult(float32(totalBytes) / float32(len(b.amount)))
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) Average() AverageResult {
	var totalCelcius Celcius = 0
	for _, v := range c.temp {
		totalCelcius += v
	}
	return AverageResult(totalCelcius / Celcius(len(c.temp)))
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) Average() AverageResult {
	var totalBytes Bytes = 0
	for _, v := range m.amount {
		totalBytes += v
	}
	return AverageResult(float32(totalBytes) / float32(len(m.amount)))
}

type Dashboard struct {
	bandwidth BandwidthUsage
	cpu       CpuTemp
	memory    MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}
	dashboard := Dashboard{bandwidth, temp, memory}
	fmt.Println(dashboard.bandwidth.Average())
	fmt.Println(dashboard.cpu.Average())
	fmt.Println(dashboard.memory.Average())
}
