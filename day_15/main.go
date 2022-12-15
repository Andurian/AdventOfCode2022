package main

import (
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"fmt"
	"math"
	"regexp"
	"strings"
)

// TODO: No need to deduplicate sensors and beacons - just represent them as pairs of points with precomputed distance
// Optimize Number of non Beacon Locations per row with the method used to find the missing beacon
type Network struct {
	beacons            []point.Point
	sensors            []point.Point
	closestBeaconIndex map[int]int
	closestSensorIndex map[int]int
}

func NewEmptyNetwork() *Network {
	return &Network{[]point.Point{}, []point.Point{}, make(map[int]int), make(map[int]int)}
}

func (n *Network) AddSensor(sensor point.Point, closestBeacon point.Point) {
	if util.IndexOf(n.sensors, sensor) != -1 {
		return
	}

	n.sensors = append(n.sensors, sensor)
	sensorIndex := len(n.sensors) - 1

	beaconIndex := util.IndexOf(n.beacons, closestBeacon)
	if beaconIndex == -1 {
		n.beacons = append(n.beacons, closestBeacon)
		beaconIndex = len(n.beacons) - 1
	}

	n.closestBeaconIndex[sensorIndex] = beaconIndex
	n.closestSensorIndex[beaconIndex] = sensorIndex
}

func (n *Network) Extent() (min, max point.Point) {
	rowMin := math.MaxInt
	rowMax := math.MinInt
	colMin := math.MaxInt
	colMax := math.MinInt

	for iSensor, sensor := range n.sensors {
		d := point.ManhattanDistance(sensor, n.beacons[n.closestBeaconIndex[iSensor]])

		rowMin = util.Min(rowMin, sensor.Row-d)
		rowMax = util.Max(rowMax, sensor.Row+d)
		colMin = util.Min(colMin, sensor.Col-d)
		colMax = util.Max(colMax, sensor.Col+d)
	}
	return point.Point{Row: rowMin, Col: colMin}, point.Point{Row: rowMax, Col: colMax}
}

func (n *Network) IsPossibleBeaconLoaction(p point.Point) bool {
	if util.IndexOf(n.beacons, p) != -1 {
		return true
	}

	for iSensor, sensor := range n.sensors {
		dBeacon := point.ManhattanDistance(sensor, n.beacons[n.closestBeaconIndex[iSensor]])
		dPoint := point.ManhattanDistance(sensor, p)
		if dPoint <= dBeacon {
			return false
		}
	}

	return true
}

func (n *Network) NonBeaconLocationsInRow(row int) int {
	min, max := n.Extent()
	counter := 0
	for col := min.Col; col <= max.Col; col += 1 {
		p := point.Point{Row: row, Col: col}

		if !n.IsPossibleBeaconLoaction(p) {
			counter += 1
		}
	}
	return counter
}

func (n *Network) FindPossibleBeaconLocations(min, max point.Point) point.Point {

	limitingSensorIds := func(p point.Point) []int {
		ret := []int{}
		for iSensor, sensor := range n.sensors {
			beacon := n.beacons[n.closestBeaconIndex[iSensor]]
			dBeacon := point.ManhattanDistance(sensor, beacon)
			dPoint := point.ManhattanDistance(sensor, p)
			if dPoint <= dBeacon {
				ret = append(ret, iSensor)
			}
		}
		return ret
	}

	rightmostCoveredPoint := func(row int, limitingSensors []int) point.Point {
		rightmostCol := min.Col
		for _, i := range limitingSensors {
			sensor := n.sensors[i]
			beacon := n.beacons[n.closestBeaconIndex[i]]
			dBeacon := point.ManhattanDistance(sensor, beacon)
			rightmostCoveredBySensor := sensor.Col + (dBeacon - util.Abs(row-sensor.Row))
			rightmostCol = util.Max(rightmostCol, rightmostCoveredBySensor)
		}
		return point.Point{Row: row, Col: rightmostCol + 1}
	}

	for row := min.Row; row <= max.Row; row += 1 {
		if row%10000 == 0 {
			fmt.Printf("Row: %d\n", row)
		}
		p := point.Point{Row: row, Col: min.Col}
		for p.Col <= max.Col {
			i := limitingSensorIds(p)
			if len(i) == 0 {
				return p
			}
			p = rightmostCoveredPoint(p.Row, i)
		}
	}
	panic("Did not find")
}

func (n *Network) String() string {
	min, max := n.Extent()
	s := ""
	for row := min.Row; row <= max.Row; row += 1 {
		for col := min.Col; col <= max.Col; col += 1 {
			p := point.Point{Row: row, Col: col}
			if util.IndexOf(n.sensors, p) != -1 {
				s += "S"
			} else if util.IndexOf(n.beacons, p) != -1 {
				s += "B"
			} else if !n.IsPossibleBeaconLoaction(p) {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func NetworkFromString(s string) *Network {
	n := NewEmptyNetwork()

	var re = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	for _, line := range strings.Split(s, "\n") {
		result := re.FindStringSubmatch(line)
		sensor := point.Point{Row: util.AtoiSafe(result[2]), Col: util.AtoiSafe(result[1])}
		closestBeacon := point.Point{Row: util.AtoiSafe(result[4]), Col: util.AtoiSafe(result[3])}
		n.AddSensor(sensor, closestBeacon)
	}

	return n
}

func Task2(network *Network) int {
	minV := 0
	maxV := 4000000
	min := point.Point{Row: minV, Col: minV}
	max := point.Point{Row: maxV, Col: maxV}
	p := network.FindPossibleBeaconLocations(min, max)
	return p.Col*4000000 + p.Row
}

func main() {
	input := util.ReadSafe("input.txt")
	network := util.PreprocessTimed(func() *Network { return NetworkFromString(input) })
	util.ExecuteTimed(15, 1, func() int { return network.NonBeaconLocationsInRow(2000000) })
	util.ExecuteTimed(15, 2, func() int { return Task2(network) })
}
