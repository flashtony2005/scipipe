package main

import (
	sp "github.com/samuell/scipipe"
)

func main() {
	sp.InitLogWarning()

	p := sp.Sh("ls", "ls -l > {o:out}")
	p.PathFormatters["out"] = func(p *sp.SciTask) string {
		return "hej.txt"
	}
	p.Prepend = "echo"
	p.Run()
}