// Code generated by https://github.com/gagliardetto/codebox. DO NOT EDIT.

package main

import "strconv"

func TaintStepTest_StrconvAppendQuote_B0I0O0(sourceCQL interface{}) interface{} {
	fromByte656 := sourceCQL.([]byte)
	intoByte414 := strconv.AppendQuote(fromByte656, "")
	return intoByte414
}

func TaintStepTest_StrconvAppendQuote_B0I1O0(sourceCQL interface{}) interface{} {
	fromString518 := sourceCQL.(string)
	intoByte650 := strconv.AppendQuote(nil, fromString518)
	return intoByte650
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	fromByte784 := sourceCQL.([]byte)
	intoByte957 := strconv.AppendQuoteToASCII(fromByte784, "")
	return intoByte957
}

func TaintStepTest_StrconvAppendQuoteToASCII_B0I1O0(sourceCQL interface{}) interface{} {
	fromString520 := sourceCQL.(string)
	intoByte443 := strconv.AppendQuoteToASCII(nil, fromString520)
	return intoByte443
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	fromByte127 := sourceCQL.([]byte)
	intoByte483 := strconv.AppendQuoteToGraphic(fromByte127, "")
	return intoByte483
}

func TaintStepTest_StrconvAppendQuoteToGraphic_B0I1O0(sourceCQL interface{}) interface{} {
	fromString989 := sourceCQL.(string)
	intoByte982 := strconv.AppendQuoteToGraphic(nil, fromString989)
	return intoByte982
}

func TaintStepTest_StrconvQuote_B0I0O0(sourceCQL interface{}) interface{} {
	fromString417 := sourceCQL.(string)
	intoString584 := strconv.Quote(fromString417)
	return intoString584
}

func TaintStepTest_StrconvQuoteToASCII_B0I0O0(sourceCQL interface{}) interface{} {
	fromString991 := sourceCQL.(string)
	intoString881 := strconv.QuoteToASCII(fromString991)
	return intoString881
}

func TaintStepTest_StrconvQuoteToGraphic_B0I0O0(sourceCQL interface{}) interface{} {
	fromString186 := sourceCQL.(string)
	intoString284 := strconv.QuoteToGraphic(fromString186)
	return intoString284
}

func TaintStepTest_StrconvUnquote_B0I0O0(sourceCQL interface{}) interface{} {
	fromString908 := sourceCQL.(string)
	intoString137, _ := strconv.Unquote(fromString908)
	return intoString137
}

func TaintStepTest_StrconvUnquoteChar_B0I0O0(sourceCQL interface{}) interface{} {
	fromString494 := sourceCQL.(string)
	_, _, intoString873, _ := strconv.UnquoteChar(fromString494, 0)
	return intoString873
}

func RunAllTaints_Strconv() {
	{
		source := newSource(0)
		out := TaintStepTest_StrconvAppendQuote_B0I0O0(source)
		sink(0, out)
	}
	{
		source := newSource(1)
		out := TaintStepTest_StrconvAppendQuote_B0I1O0(source)
		sink(1, out)
	}
	{
		source := newSource(2)
		out := TaintStepTest_StrconvAppendQuoteToASCII_B0I0O0(source)
		sink(2, out)
	}
	{
		source := newSource(3)
		out := TaintStepTest_StrconvAppendQuoteToASCII_B0I1O0(source)
		sink(3, out)
	}
	{
		source := newSource(4)
		out := TaintStepTest_StrconvAppendQuoteToGraphic_B0I0O0(source)
		sink(4, out)
	}
	{
		source := newSource(5)
		out := TaintStepTest_StrconvAppendQuoteToGraphic_B0I1O0(source)
		sink(5, out)
	}
	{
		source := newSource(6)
		out := TaintStepTest_StrconvQuote_B0I0O0(source)
		sink(6, out)
	}
	{
		source := newSource(7)
		out := TaintStepTest_StrconvQuoteToASCII_B0I0O0(source)
		sink(7, out)
	}
	{
		source := newSource(8)
		out := TaintStepTest_StrconvQuoteToGraphic_B0I0O0(source)
		sink(8, out)
	}
	{
		source := newSource(9)
		out := TaintStepTest_StrconvUnquote_B0I0O0(source)
		sink(9, out)
	}
	{
		source := newSource(10)
		out := TaintStepTest_StrconvUnquoteChar_B0I0O0(source)
		sink(10, out)
	}
}
