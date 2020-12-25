package main

import (
	"github.com/mobyvb/temp-conversion/convert1"
	"github.com/mobyvb/temp-conversion/convert2"
	"github.com/mobyvb/temp-conversion/convert3"
	"github.com/mobyvb/temp-conversion/convert4"
	"github.com/mobyvb/temp-conversion/convert5"
	"github.com/mobyvb/temp-conversion/convert6"
	"github.com/mobyvb/temp-conversion/convert7"
	"github.com/mobyvb/temp-conversion/convertmany"
	"github.com/mobyvb/temp-conversion/tempunit"
)

func main() {
	c1 := 30.0
	f1 := 86.0
	k1 := 303.15
	n1 := 9.9

	convert1.Convert(c1, f1, k1, n1)
	convert2.Convert(c1, f1, k1, n1)
	convert3.Convert(c1, f1, k1, n1)
	convert4.Convert(c1, f1, k1, n1)
	convert5.Convert(c1, f1, k1, n1)
	convert6.Convert(c1, f1, k1, n1)
	convert7.Convert(c1, f1, k1, n1)
	convertmany.Convert(tempunit.C, tempunit.F, 30, 0, 20, 10)
}
