// main package
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tree struct {
	Root *Node
}

type Node struct {
	parent   *Node
	children []*Node
	size     int
	fileName string
	isFile   bool
}

func (p *Node) print() {
	if !(p.parent == nil) {
		fmt.Printf("Parent: %v\n", p.parent.fileName)
	}
	if p.isFile {
		fmt.Printf("File: %v\t%d\n", p.fileName, p.size)
	} else {
		fmt.Printf("Directory: %v\t%d\n", p.fileName, p.size)
	}

	if len(p.children) > 0 {
		fmt.Println("Children")
		for _, el := range p.children {
			el.print()
		}
	}
}

func (p *Node) printChildren() string {
	var childs string
	for _, el := range p.children {
		if !el.isFile {
			childs += fmt.Sprintf("%s, ", el.fileName)
		}
	}
	childs += fmt.Sprintf("\n")
	return childs
}

func (p *Node) calculateDirSize() int {
	if len(p.children) == 0 {
		return p.size
	} else {
		for _, el := range p.children {
			p.size += el.calculateDirSize()
		}
	}
	return p.size
}

func (p *Node) sumUnder100K() int {
	var sum int
	for _, el := range p.children {
		if !el.isFile {
			sum += el.sumUnder100K()
		}
		if !el.isFile && el.size < 100000 {
			sum += el.size
		}
	}
	return sum
}

func (p *Node) findSmallestNeed(n int) int {
	smallest := math.MaxInt
	for _, el := range p.children {
		if !el.isFile && el.size > n && len(el.children) > 0 {
			smol_child := el.findSmallestNeed(n)
			if smol_child < el.size {
				if smol_child < smallest {
					smallest = smol_child
				}
			}
		}
		if !el.isFile && el.size > n {
			if el.size < smallest {
				smallest = el.size
			}
		}
	}
	return smallest
}

func (p *Node) findNodeByName(n string) *Node {
	for _, el := range p.children {
		if el.fileName == n {
			return el
		}
	}
	return nil
}

func part1(f string) int {
	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan first line and initialize tree root
	scanner.Scan()
	t := Tree{
		Root: &Node{
			parent:   nil,
			children: nil,
			size:     0,
			fileName: "/",
			isFile:   false,
		},
	}

	pwd := t.Root

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "$ ls") {
			for scanner.Scan() {
				line = scanner.Text()
				if strings.Contains(line, "$") {
					break
				}
				fields := strings.Split(line, " ")
				if strings.Contains(fields[0], "dir") {
					node := Node{
						parent:   pwd,
						children: nil,
						size:     0,
						fileName: fields[1],
						isFile:   false,
					}
					pwd.children = append(pwd.children, &node)
				} else {
					fSize, _ := strconv.Atoi(fields[0])
					node := Node{
						parent:   pwd,
						children: nil,
						size:     fSize,
						fileName: fields[1],
						isFile:   true,
					}
					pwd.children = append(pwd.children, &node)
				}
			}
		}

		if strings.Contains(line, "$ cd ..") {
			pwd = pwd.parent
		} else if strings.Contains(line, "$ cd") {
			dir := strings.Split(line, " ")[2]
			pwd = pwd.findNodeByName(dir)
		}
	}

	t.Root.calculateDirSize()

	return t.Root.sumUnder100K()
}

func part2(f string) int {
	file, _ := os.Open(f)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan first line and initialize tree root
	scanner.Scan()
	t := Tree{
		Root: &Node{
			parent:   nil,
			children: nil,
			size:     0,
			fileName: "/",
			isFile:   false,
		},
	}

	pwd := t.Root

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "$ ls") {
			for scanner.Scan() {
				line = scanner.Text()
				if strings.Contains(line, "$") {
					break
				}
				fields := strings.Split(line, " ")
				if strings.Contains(fields[0], "dir") {
					node := Node{
						parent:   pwd,
						children: nil,
						size:     0,
						fileName: fields[1],
						isFile:   false,
					}
					pwd.children = append(pwd.children, &node)
				} else {
					fSize, _ := strconv.Atoi(fields[0])
					node := Node{
						parent:   pwd,
						children: nil,
						size:     fSize,
						fileName: fields[1],
						isFile:   true,
					}
					pwd.children = append(pwd.children, &node)
				}
			}
		}

		if strings.Contains(line, "$ cd ..") {
			pwd = pwd.parent
		} else if strings.Contains(line, "$ cd") {
			dir := strings.Split(line, " ")[2]
			pwd = pwd.findNodeByName(dir)
		}
	}

	t.Root.calculateDirSize()

	free := 70000000 - t.Root.size
	need := 30000000 - free

	return t.Root.findSmallestNeed(need)
}

func main() {
	ta := time.Now()
	log.Printf("Test input solution: %d\n", part2("test.txt"))
	tb := time.Now()
	log.Printf("Test input took: %v\n", tb.Sub(ta))
	log.Printf("Part 1 solution: %d\n", part1("input.txt"))
	ta = time.Now()
	log.Printf("Part 1 took: %v\n", ta.Sub(tb))
	log.Printf("Part 2 solution: %d\n", part2("input.txt"))
	tb = time.Now()
	log.Printf("Part 2 took: %v\n", tb.Sub(ta))
}
