package compiler

import "bytes"
import "text/scanner"

//A cache is a storage container that contains code. It can be compiled at a later point in time.
type Cache struct {
	bytes.Buffer

	filename string
	line     int

	Match string
}

func (c *Compiler) NewCustomCache(begin, end func(*Compiler) bool) Cache {
	var cache Cache

	cache.filename = c.Scanners[len(c.Scanners)-1].Filename
	cache.line = c.Scanners[len(c.Scanners)-1].Line - 1

	var depth = 1
	for {
		c.Scanners[len(c.Scanners)-1].Scan()
		tok := c.Scanners[len(c.Scanners)-1].TokenText()
		c.token = tok

		if end(c) || tok == "" {
			depth--
			if depth == 0 {
				cache.Match = tok
				break
			}
		} else if begin(c) {
			depth++
		}

		cache.Write([]byte(tok))
	}

	cache.Write([]byte("\n"))

	return cache
}

//Create a new cache within open and matching close characters such that they should match.
// eg. NewCache("{", "}") will scan a single code block including any children blocks.
func (c *Compiler) NewCache(open string, matches ...string) Cache {
	return c.NewCustomCache(func(c *Compiler) bool {
		return c.Token() == open
	}, func(c *Compiler) bool {
		for _, match := range matches {
			if c.Token() == match {
				return true
			}
		}
		return false
	})
}

//Compile the given cache, pretending the filename is name and the line number is line.
func (c *Compiler) CompileCache(cache Cache) {

	var s scanner.Scanner
	s.Init(&cache)
	s.Filename = cache.filename
	s.Whitespace = 0

	c.Scanners = append(c.Scanners, &s)
	c.CurrentLines = append(c.CurrentLines, "")
	c.LineOffset = cache.line

	var length = len(c.Scanners)

	for {
		c.ScanStatement()

		if len(c.Scanners) < length {
			c.LineOffset = 0
			return
		}
	}
}

//Load the cache as if it was the next set of scannable tokens.
func (c *Compiler) LoadCache(cache Cache) {
	var s scanner.Scanner
	s.Init(&cache)
	s.Filename = cache.filename
	s.Whitespace = 0

	c.Scanners = append(c.Scanners, &s)
	c.CurrentLines = append(c.CurrentLines, "")
	c.LineOffset = cache.line
}
