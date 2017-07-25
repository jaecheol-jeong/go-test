package main

import "github.com/tealeg/xlsx"

func main() {

	xlfile := xlsx.NewFile()
	sheet, err := xlfile.AddSheet("Sheet1")
	if err != nil {
		panic(err)
	}

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.HMerge = 3
	cell.SetValue("이베이쇼핑 무이자 할부 미노출건")
	style := xlsx.NewStyle()
	cell.SetStyle(style)

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetValue(1)
	cell = row.AddCell()
	cell.SetValue(2)

	cell = row.AddCell()
	cell.SetFormula("=sum(A2:B2)")

	err = xlfile.Save("./test.xlsx")
	if err != nil {
		panic(err)
	}
}
