package main

import (
	"strings"

	gotag "github.com/Apfel/gotag"
	gotaglib "github.com/Apfel/gotag/libraries"
	"github.com/MikeModder/anpan"
)

var parser = gotag.NewBuilder().AddMethods(gotaglib.Arguments()).AddMethods(gotaglib.Functional()).AddMethods(gotaglib.Miscellaneous()).AddMethods(gotaglib.Strings()).AddMethods(gotaglib.Time()).AddMethods(gotaglib.Variables()).Build()

func tagcmd(ctx anpan.Context, args []string) {
	ctx.Reply(parser.Parse(strings.Join(args, " ")))
}
