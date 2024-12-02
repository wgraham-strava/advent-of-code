package main

import (
    "fmt"
    "time"
    "os"
    "bufio"
    "math"
    "strings"
    "strconv"
)

type Conversion struct {
    Source int
    Dest int
    Range int
}

type Seed struct {
    Start int
    Range int
}

var seed_slice = make([]int, 0)
var seeds *[]Seed
var seeds_soil *[]Conversion
var soil_fertilizer *[]Conversion
var fertilizer_water *[]Conversion
var water_light *[]Conversion
var light_temp *[]Conversion
var temp_humidity *[]Conversion
var humidity_location *[]Conversion

func lookup(m *[]Conversion, v int) int {
    for _, c := range *m {
        if v >= c.Source && v < c.Source + c.Range {
            return c.Dest + (v - c.Source)
        }
    }
    return v
}

func build_seed_list (s_slice []int) *[]Seed {
    s := make([]Seed, 0)
    for i := 0; i < len(s_slice); i += 2 {
        s = append(s, Seed{
            Start: s_slice[i],
            Range: s_slice[i+1],
        })
    }
    return &s
}

func build_conversion_list(s *bufio.Scanner) *[]Conversion {
    var c = make([]Conversion, 0)
    for s.Scan() {
        line := s.Text()
        if len(line) == 0 {
            return &c
        }
        dest, _ := strconv.Atoi(strings.Fields(line)[0])
        source, _ := strconv.Atoi(strings.Fields(line)[1])
        count, _ := strconv.Atoi(strings.Fields(line)[2])

        c = append(c, Conversion{
            Source: source,
            Dest: dest,
            Range: count,
        })
    }
    return &c
}

func seed_to_location(i int) int {
    return lookup(humidity_location, lookup(temp_humidity, lookup(light_temp, lookup(water_light, lookup(fertilizer_water, lookup(soil_fertilizer, lookup(seeds_soil, i)))))))
}

func part1() int {
    var ans int = math.MaxInt

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	for lineScanner.Scan() {
        line := lineScanner.Text()

        if strings.Contains(line, "seeds"){
            for _, seed := range strings.Fields(strings.Split(line, ":")[1]) {
                s, _ := strconv.Atoi(seed)
                seed_slice = append(seed_slice, s)
            }
        }
        if strings.Contains(line, "seed-to-soil") {
            seeds_soil = build_conversion_list(lineScanner)
        }
        if strings.Contains(line, "soil-to-fertilizer") {
            soil_fertilizer = build_conversion_list(lineScanner)
        }
        if strings.Contains(line, "fertilizer-to-water") {
            fertilizer_water = build_conversion_list(lineScanner)
        }
        if strings.Contains(line, "water-to-light") {
            water_light = build_conversion_list(lineScanner)
        }
        if strings.Contains(line, "light-to-temperature") {
            light_temp = build_conversion_list(lineScanner)
        }
        if strings.Contains(line, "temperature-to-humidity") {
            temp_humidity = build_conversion_list(lineScanner)
        }
        if strings.Contains(line, "humidity-to-location") {
            humidity_location = build_conversion_list(lineScanner)
        }
    }
    for _, i := range seed_slice {
        loc := seed_to_location(i)
        // fmt.Printf("Seed %d -> Location %d\n", i, loc)
        ans = int(math.Min(float64(ans), float64(loc)))
    }

    return ans
}

func part2() int {
    var ans int = math.MaxInt

    seeds = build_seed_list(seed_slice)

    for _, seed := range *seeds {
        fmt.Printf("Considering all seeds between %d - %d\n", seed.Start, seed.Start + seed.Range)
        for i := seed.Start; i < seed.Start + seed.Range; i++ {
            location := seed_to_location(i)
            ans = int(math.Min(float64(ans), float64(location)))
        }
    }

    return ans
}

func main() {
	start := time.Now()

	fmt.Println("Part 1 Solution:", part1())
    p1_time := time.Since(start)
	fmt.Println("Part 1 took", p1_time)

    fmt.Println("Part 2 Solution:", part2())
	fmt.Println("Part 2 took", time.Since(start) - p1_time)
}
