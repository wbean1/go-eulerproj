package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Triangle struct {
	a Coord
	b Coord
	c Coord
}

func main() {
	triangles := parseFile("p102_triangles.txt")
	var sum int
	origin := Coord{0, 0}
	for _, t := range triangles {
		if triangleContains(t, origin) {
			sum++
		}
	}
	fmt.Printf("number of triangles containing 0,0: %d", sum)
}

func parseFile(file string) []Triangle {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	triangles := []Triangle{}
	for scanner.Scan() {
		t := parseTriangle(scanner.Text())
		triangles = append(triangles, t)
	}
	return triangles
}

func parseTriangle(line string) Triangle {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, ",")
	p := []int{}
	for _, str := range parts {
		k, _ := strconv.Atoi(str)
		p = append(p, k)
	}
	c1 := Coord{p[0], p[1]}
	c2 := Coord{p[2], p[3]}
	c3 := Coord{p[4], p[5]}
	return Triangle{c1, c2, c3}
}

func triangleContains(t Triangle, c Coord) bool {
	as_x := c.x - t.a.x
	as_y := c.y - t.a.y

	s_ab := (t.b.x-c.x)*as_y-(t.b.y-c.y)*as_x > 0
	if (t.c.x-t.a.x)*as_y-(t.c.y-t.a.y)*as_x > 0 == s_ab {
		return false
	}
	if (t.c.x-t.b.x)*(c.y-t.b.y)-(t.c.y-t.b.y)*(c.x-t.b.x) > 0 != s_ab {
		return false
	}
	return true
}
