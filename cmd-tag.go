package main

import (
	"strings"

	gotag "github.com/Apfel/GOtag"
	gotaglib "github.com/Apfel/GOtag/libraries"
	"github.com/MikeModder/anpan"
)

var parser = gotag.NewBuilder().AddMethods(gotaglib.Arguments()).Build()

func tagcmd(ctx anpan.Context, args []string) {
	ctx.Reply(parser.Parse(strings.Join(args, " ")))
}
