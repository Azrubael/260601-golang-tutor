package main

import (
	"strings"

	"github.com/xuri/excelize/v2"
)

// EnsureWrapText встановлює wrap_text для ячейки, зберігши інші властивості
func EnsureWrapText(f *excelize.File, sheet, coord string) error {
    // Отримати поточний styleID (0 якщо немає)
    styleID, err := f.GetCellStyle(sheet, coord)
    if err != nil {
        return err
    }

    if styleID == 0 { // Якщо немає стилю — застосуємо простий wrap-only стиль
        newStyle, err := f.NewStyle(&excelize.Style{
            Alignment: &excelize.Alignment{WrapText: true},
        })
        if err != nil {
            return err
        }
        return f.SetCellStyle(sheet, coord, coord, newStyle)
    }

    origStyle, err := f.GetStyle(styleID)   // Отримати повний опис стилю
    if err != nil {
        // fallback: застосуємо wrap-only стиль
        newStyle, err2 := f.NewStyle(&excelize.Style{
            Alignment: &excelize.Alignment{WrapText: true},
        })
        if err2 != nil {
            return err2
        }
        return f.SetCellStyle(sheet, coord, coord, newStyle)
    }

    // deep copy структури Style і втановлення WrapText = true.
    s := &excelize.Style{}

    if origStyle.Font != nil {
        tmp := *origStyle.Font
        s.Font = &tmp
    }
    tmpFill := origStyle.Fill
    if tmpFill.Pattern < 0 || tmpFill.Pattern > 18 {
        tmpFill.Pattern = 0
    }
    if tmpFill.Color != nil {
        col := make([]string, len(tmpFill.Color))
        copy(col, tmpFill.Color)
        tmpFill.Color = col
    }
    // If Fill was entirely zero-value, this still copies zero-value which is fine.
    s.Fill = tmpFill

    if len(origStyle.Border) > 0 {
        b := make([]excelize.Border, len(origStyle.Border))
        copy(b, origStyle.Border)
        s.Border = b
    }
    if len(origStyle.Border) > 0 {
        b := make([]excelize.Border, len(origStyle.Border))
        copy(b, origStyle.Border)
        s.Border = b
    }
    if origStyle.Protection != nil {
        p := *origStyle.Protection
        s.Protection = &p
    }
    if origStyle.NumFmt != 0 {
        s.NumFmt = origStyle.NumFmt
    }
    // Alignment: копіюємо або створюємо нову і встановимо WrapText = true
    if origStyle.Alignment != nil {
        a := *origStyle.Alignment
        a.WrapText = true
        s.Alignment = &a
    } else {
        s.Alignment = &excelize.Alignment{WrapText: true}
    }

    // Якщо будуть інші елементи стилю, можна додати їх тут.

    // Створюємо новий стиль на основі скопійованого опису
    newStyleID, err := f.NewStyle(s)
    if err != nil {
        // Якщо все ще помилка — як останній захід застосуємо простий wrap-only стиль
        fallback, ferr := f.NewStyle(&excelize.Style{Alignment: &excelize.Alignment{WrapText: true}})
        if ferr != nil {
            return err
        }
        return f.SetCellStyle(sheet, coord, coord, fallback)
    }

    // Застосовуємо новий стиль до клітинки
    return f.SetCellStyle(sheet, coord, coord, newStyleID)
}


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