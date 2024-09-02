package excel

import "github.com/xuri/excelize/v2"

type HeaderTitleCol struct {
	topLeft     string
	bottomRight string
	title       string
}

type ColumnSheetStyle struct {
	col   string
	width float64
}

func createHeaders(f *excelize.File, headers []HeaderTitleCol, sheetName string) error {
	headerStyle, err := createHeaderStyle(f)
	if err != nil {
		return err
	}
	for _, header := range headers {
		if err := createHeaderTitle(f, headerStyle, sheetName, header); err != nil {
			return err
		}
	}

	return nil
}
func createHeaderStyle(f *excelize.File) (int, error) {
	borderStyle := []excelize.Border{
		{Type: "left", Color: "000000", Style: 1},
		{Type: "top", Color: "000000", Style: 1},
		{Type: "bottom", Color: "000000", Style: 1},
		{Type: "right", Color: "000000", Style: 1},
	}
	alignmentStyle := &excelize.Alignment{
		Horizontal: "center",
		Vertical:   "center",
		WrapText:   true,
	}
	fontStyle := &excelize.Font{
		Family: "Arial Narrow",
		Size:   8,
	}
	return f.NewStyle(&excelize.Style{
		Alignment: alignmentStyle,
		Border:    borderStyle,
		Font:      fontStyle,
	})
}
func createHeaderTitle(f *excelize.File, style int, sheetName string, props HeaderTitleCol) error {
	if err := f.MergeCell(sheetName, props.topLeft, props.bottomRight); err != nil {
		return err
	}
	if err := f.SetCellStyle(sheetName, props.topLeft, props.bottomRight, style); err != nil {
		return err
	}
	return f.SetCellStr(sheetName, props.topLeft, props.title)
}

func setSheetStyles(f *excelize.File, columns []ColumnSheetStyle, sheetName string) error {
	for _, col := range columns {
		if err := f.SetColWidth(sheetName, col.col, col.col, col.width); err != nil {
			return err
		}
	}

	return nil
}
