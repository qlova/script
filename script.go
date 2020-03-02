package script

import (
	"encoding/base64"
	"math/big"
)

//Ctx is the script ctx.
type Ctx struct {
	Language
	*ctx
}

type AnyCtx interface {
	RootCtx() Ctx
}

func (q Ctx) RootCtx() Ctx {
	return q
}

func NewCtx() Ctx {
	var ctx = Ctx{
		nil, new(ctx),
	}
	ctx.id = make(map[string]int64)
	ctx.True = ctx.Bool(true)
	ctx.False = ctx.Bool(false)
	return ctx
}

type ctx struct {
	//Ctx variables.
	id map[string]int64

	defining  bool
	variables []string

	False Bool
	True  Bool
}

func (ctx Ctx) ID(prefix string) string {
	ctx.id[prefix]++
	return prefix + base64.RawURLEncoding.EncodeToString(big.NewInt(ctx.id[prefix]).Bytes())
}

//Script is a script.
type Script func(Ctx)
