/*
	This tool generates a language interface and associated language templates.
*/
package main

import "math/big"

type Flag *big.Int
func Flags(flags ...Flag) Flag {
	var final Flag = big.NewInt(1)
	for _, flag := range flags {
		(*big.Int)(final).Mul((*big.Int)(final), (*big.Int)(flag))
	}
	return final
} 

func main() {
	GenerateLanguageInterface()
	GenerateLanguageTypes()
	
	GenerateLanguageTemplate("Example")
}
