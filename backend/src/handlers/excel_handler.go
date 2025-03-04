package handlers

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	//"time"
)

func CreateExcel(input string) (*excelize.File, error) {
	excelFile := excelize.NewFile()
	sheet := "Sheet1"

	SupapData := GetAttributes(input)

	headerStyle, err := excelFile.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
		Font: &excelize.Font{
			Bold: true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"ADD8E6"}, Pattern: 1},
	})

	cellStyle, err := excelFile.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})

	if err != nil {
		panic("Problem styling excel")
	}

	// Write data to Excel
	excelFile.SetCellValue(sheet, "A1", "SUPAR")
	excelFile.SetCellValue(sheet, "B1", "MAKE-MODEL")
	excelFile.SetCellValue(sheet, "C1", "IN/EX")
	excelFile.SetCellValue(sheet, "D1", "OEM")

	excelFile.SetCellStyle(sheet, "A1", "A1", headerStyle)
	excelFile.SetCellStyle(sheet, "B1", "B1", headerStyle)
	excelFile.SetCellStyle(sheet, "C1", "C1", headerStyle)
	excelFile.SetCellStyle(sheet, "D1", "D1", headerStyle)

	var lastCol int
	for i, s := range SupapData {
		row := i + 2 // Start from row 2 (row 1 is for headers)
		excelFile.SetCellValue(sheet, fmt.Sprintf("A%d", row), s.SuparCode)
		excelFile.SetCellValue(sheet, fmt.Sprintf("B%d", row), s.MakeModel)
		excelFile.SetCellValue(sheet, fmt.Sprintf("C%d", row), s.Type)
		excelFile.SetCellValue(sheet, fmt.Sprintf("D%d", row), s.OriginalCode)

		excelFile.SetCellStyle(sheet, fmt.Sprintf("A%d", row), fmt.Sprintf("A%d", row), cellStyle)
		excelFile.SetCellStyle(sheet, fmt.Sprintf("B%d", row), fmt.Sprintf("B%d", row), cellStyle)
		excelFile.SetCellStyle(sheet, fmt.Sprintf("C%d", row), fmt.Sprintf("C%d", row), cellStyle)
		excelFile.SetCellStyle(sheet, fmt.Sprintf("D%d", row), fmt.Sprintf("D%d", row), cellStyle)
		lastCol = row
	}

	for col := 1; col <= lastCol; col++ {
		colLetter, _ := excelize.ColumnNumberToName(col)
		excelFile.SetColWidth(sheet, colLetter, colLetter, 15) // Set width (20 is a reasonable default)
	}
	excelFile.SetColWidth(sheet, "A", "D", 35)

	// Save the file
	//filePath := "Spesifikasyon_" + time.Now().Format("2025-01-02") + " .xlsx"
	////filePath := "output.xlsx"
	//if err := excelFile.SaveAs(filePath); err != nil {
	//	return , err
	//}

	return excelFile, nil
}
