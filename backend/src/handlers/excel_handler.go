package handlers

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	_ "image/jpeg"
	_ "image/png"
)

func CreateExcel(input string) (*excelize.File, error) {
	excelFile := excelize.NewFile()
	sheet := "Sheet1"

	SupapData := GetSpecification(input)

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

func CreateProforma(input string) (*excelize.File, error) {
	excelFile := excelize.NewFile()
	sheet := "Sheet1"
	hideGrid := false
	excelFile.SetSheetView(sheet, 0, &excelize.ViewOptions{
		ShowGridLines: &hideGrid,
	})

	SupapData := GetSpecification(input)

	thickHeaderStyle, err := excelFile.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
		Font: &excelize.Font{
			Bold: true,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 5},
			{Type: "top", Color: "000000", Style: 5},
			{Type: "bottom", Color: "000000", Style: 5},
			{Type: "right", Color: "000000", Style: 5},
		},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#000000"}, Pattern: 4},
	})

	refStyle, err := excelFile.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
		},
		Font: &excelize.Font{
			Bold: true,
		},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#000000"}, Pattern: 4},
	})

	outlineStyle, err := excelFile.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 5},
			{Type: "top", Color: "000000", Style: 5},
			{Type: "bottom", Color: "000000", Style: 5},
			{Type: "right", Color: "000000", Style: 5},
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
	})

	excelFile.SetDefaultFont("Arial")

	dataCellStyle, err := excelFile.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "Left",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
		},
		Font: &excelize.Font{
			Family: "Arial",
			Bold:   true,
		},
	})

	picErr := excelFile.AddPicture("Sheet1", "A1", "./assets/images/supar_logo.png", &excelize.GraphicOptions{
		ScaleX:  1.2, // Scale the image width to 50%
		OffsetY: 10,
	})
	picErr = excelFile.AddPicture("Sheet1", "E1", "./assets/images/tuv_nord.jpg", &excelize.GraphicOptions{
		ScaleX:  0.4,
		ScaleY:  0.4,
		OffsetY: 10,
	})
	if picErr != nil {
		panic(picErr)
	}

	if err != nil {
		panic("Problem styling excel")
	}

	excelFile.SetColWidth(sheet, "A", "A", 13)
	excelFile.SetColWidth(sheet, "B", "B", 30)
	excelFile.SetColWidth(sheet, "C", "C", 10)
	excelFile.SetColWidth(sheet, "D", "D", 15)
	excelFile.SetColWidth(sheet, "E", "F", 10)
	excelFile.SetColWidth(sheet, "G", "G", 13)
	excelFile.SetRowHeight(sheet, 2, 30)
	excelFile.SetRowHeight(sheet, 1, 1)

	// Merge cells
	excelFile.MergeCell(sheet, "A9", "C9")
	excelFile.MergeCell(sheet, "A10", "C10")
	excelFile.MergeCell(sheet, "A11", "C11")
	excelFile.MergeCell(sheet, "A12", "C12")
	excelFile.MergeCell(sheet, "D12", "G12")

	excelFile.MergeCell(sheet, "B14", "D14")
	excelFile.MergeCell(sheet, "B15", "D15")
	excelFile.MergeCell(sheet, "E14", "G14")
	excelFile.MergeCell(sheet, "E15", "G15")

	excelFile.SetCellValue(sheet, "A2", "         SUPAR SUPAP VE PARCA SAN TİC A.S.")
	applyStyle(excelFile, sheet, "A2", &excelize.Style{
		Font: &excelize.Font{
			Size: 16,
			Bold: true,
		},
	})

	excelFile.SetCellStyle(sheet, "A14", "G14", thickHeaderStyle)
	excelFile.SetCellValue(sheet, "A14", "Mode of Transport")
	applyStyle(excelFile, sheet, "A14", &excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 8,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 5},
			{Type: "top", Color: "000000", Style: 5},
			{Type: "bottom", Color: "000000", Style: 5},
			{Type: "right", Color: "000000", Style: 5},
		},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#000000"}, Pattern: 4},
	})

	excelFile.SetCellStyle(sheet, "A15", "A15", outlineStyle)
	excelFile.SetCellStyle(sheet, "B15", "D15", outlineStyle)
	excelFile.SetCellStyle(sheet, "E15", "G15", outlineStyle)

	applyStyle(excelFile, sheet, "A9", &excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 5},
			{Type: "top", Color: "000000", Style: 5},
			{Type: "right", Color: "000000", Style: 5},
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
	})
	applyStyle2Cells(excelFile, sheet, "A9", "C9", &excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 5},
			{Type: "top", Color: "000000", Style: 5},
			{Type: "right", Color: "000000", Style: 5},
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
	})
	applyStyle2Cells(excelFile, sheet, "A10", "C10", &excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 5},
			{Type: "right", Color: "000000", Style: 5},
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
	})
	applyStyle2Cells(excelFile, sheet, "A11", "C11", &excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 5},
			{Type: "right", Color: "000000", Style: 5},
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
	})
	applyStyle2Cells(excelFile, sheet, "A12", "C12", &excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 5},
			{Type: "bottom", Color: "000000", Style: 5},
			{Type: "right", Color: "000000", Style: 5},
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
	})

	//excelFile.SetCellStyle(sheet, "A9", "C12", outlineStyle)
	excelFile.SetCellValue(sheet, "A3", "Factory: 1. Organize Sanayi Bölgesi Horozluhan Mahallesi Güven Caddesi No:150")
	excelFile.SetCellValue(sheet, "A4", "42120 Selçuklu-Konya/TÜRKİYE")
	excelFile.SetCellValue(sheet, "A5", "Tel: (90) (332) 248 2394  ")
	excelFile.SetCellValue(sheet, "A6", "Fax: (90) (332) 2487521")
	excelFile.SetCellValue(sheet, "A7", "e-mail: export@supar.com")

	excelFile.SetCellValue(sheet, "B14", "Payment")
	excelFile.SetCellValue(sheet, "E14", "Terms Of Delivery")

	excelFile.SetRowHeight(sheet, 16, 6)
	// ROW 17
	excelFile.SetCellValue(sheet, "A17", "Order #")
	excelFile.SetCellValue(sheet, "B17", "Description of Goods")
	excelFile.SetCellValue(sheet, "C17", "Supar Code")
	excelFile.SetCellValue(sheet, "D17", "Item No")
	excelFile.SetCellValue(sheet, "E17", "Qty (Pcs.)")
	excelFile.SetCellValue(sheet, "F17", "Unit Price")
	excelFile.SetCellValue(sheet, "G17", "TOTAL USD")

	cols := []string{"A", "B", "C", "D", "E", "F", "G"}
	for i := range cols {
		applyStyle(excelFile, sheet, fmt.Sprintf("%s17", cols[i]), &excelize.Style{
			Font: &excelize.Font{
				Size: 9,
				Bold: true,
			},
			Alignment: &excelize.Alignment{
				Horizontal: "center",
			},
			Border: []excelize.Border{
				{Type: "left", Color: "000000", Style: 2},
				{Type: "top", Color: "000000", Style: 2},
				{Type: "bottom", Color: "000000", Style: 6},
				{Type: "right", Color: "000000", Style: 2},
			},
			Fill: excelize.Fill{Type: "pattern", Color: []string{"#000000"}, Pattern: 4},
		})
	}

	//excelFile.SetCellStyle(sheet, "C1", "E1", thickHeaderStyle)

	var lastCol int
	for i, s := range SupapData {
		row := i + 20 // Start from row 2 (row 1 is for headers)
		excelFile.SetCellValue(sheet, fmt.Sprintf("B%d", row), s.OrderNo)
		excelFile.SetCellValue(sheet, fmt.Sprintf("B%d", row), s.MakeModel)
		excelFile.SetCellValue(sheet, fmt.Sprintf("C%d", row), s.SuparCode)
		excelFile.SetCellValue(sheet, fmt.Sprintf("D%d", row), s.ItemNo)
		excelFile.SetCellValue(sheet, fmt.Sprintf("E%d", row), s.Quantity)
		excelFile.SetCellValue(sheet, fmt.Sprintf("F%d", row), s.UnitPrice)

		lastCol = row + 5
	}

	excelFile.SetCellStyle(sheet, "A18", fmt.Sprintf("G%d", lastCol), dataCellStyle)

	excelFile.SetCellValue(sheet, "E9", "Date :")
	excelFile.SetCellValue(sheet, "E10", "Your Order :")
	excelFile.SetCellValue(sheet, "E11", "Our Ref : ")

	excelFile.SetCellStyle(sheet, "E9", "E9", refStyle)
	excelFile.SetCellStyle(sheet, "E10", "E10", refStyle)
	excelFile.SetCellStyle(sheet, "E11", "E11", refStyle)

	excelFile.SetRowHeight(sheet, 8, 6)
	excelFile.SetRowHeight(sheet, 13, 6)
	excelFile.SetCellValue(sheet, "D12", "COMMERCIAL INVOICE")
	applyStyle(excelFile, sheet, "D12", &excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
		},
		Font: &excelize.Font{
			Size:   14,
			Bold:   true,
			Family: "Arial",
		},
	})

	//for col := 1; col <= lastCol; col++ {
	//	colLetter, _ := excelize.ColumnNumberToName(col)
	//	excelFile.SetColWidth(sheet, colLetter, colLetter, 15) // Set width (20 is a reasonable default)
	//}
	//excelFile.SetColWidth(sheet, "A", "D", 35)

	return excelFile, nil
}

func applyStyle(f *excelize.File, sheet, cell string, styleDef *excelize.Style) error {
	styleID, err := f.NewStyle(styleDef)
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, cell, cell, styleID)
}

func applyStyle2Cells(f *excelize.File, sheet, cell string, cell2 string, styleDef *excelize.Style) error {
	styleID, err := f.NewStyle(styleDef)
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, cell, cell2, styleID)
}
