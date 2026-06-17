package main

import (
	"strings"

	"github.com/xuri/excelize/v2"
)

// SetRowHeightXlsx встановлює висоту рядка, зберігши інші властивості
func SetRowHeightXlsx(f *excelize.File, sheet string, row int, height float64, txt string) error {
    
    wrap_lines := len(strings.Fields(txt))
    if wrap_lines == 1 {
        return nil
    }
    required_height := (float64(wrap_lines)) * height
    err_height := f.SetRowHeight(sheet, row, required_height)
    
    return err_height
}