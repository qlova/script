package compiler

import "bytes"
import "text/scanner"

//A cache is a storage container that contains code. It can be compiled at a later point in time.
type Cache struct {
	bytes.Buffer
}

//Create a new cache within open and close characters such that they should match.
// eg. NewCache("{", "}") will scan a single code block including any children blocks.
func (c *Compiler) NewCache(open, close string) Cache {
	var cache Cache
	
	var depth = 1
	for {
		c.Scanners[len(c.Scanners)-1].Scan()
		tok := c.Scanners[len(c.Scanners)-1].TokenText()
		
		if tok == close {
			depth--
			if depth == 0 {
				break
			}
		} else if tok == open {
			depth++
		}
		
		cache.Write([]byte(tok))
	}
	
	return cache
}

//Compile the given cache, pretending the filename is name and the line number is line.
func (c *Compiler) CompileCache(cache Cache, name string, line int) {
	
	
	var s scanner.Scanner
	s.Init(&cache)
	s.Filename = name
	s.Whitespace = 0

	c.Scanners = append(c.Scanners, &s)
	c.CurrentLines = append(c.CurrentLines, "")
	c.LineOffset = line
	
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
func (c *Compiler) LoadCache(cache Cache, name string, line int) {
	var s scanner.Scanner
	s.Init(&cache)
	s.Filename = name
	s.Whitespace = 0

	c.Scanners = append(c.Scanners, &s)
	c.CurrentLines = append(c.CurrentLines, "")
	c.LineOffset = line
}
