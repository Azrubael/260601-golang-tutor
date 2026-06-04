Підказки для роботи зі стилями MSExcel

```go
    var err_txt string = "Помилка створення стилю для клітинок "
    style2, err2 := vopi_xlsx.NewStyle(&excelize.Style{
        Font:      &excelize.Font{Family: "Times New Roman", Size: 10},
        Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: false},
    })
    if err2 != nil {
        fmt.Println(err_txt, "зі званням:", err2)
        os.Exit(6)
    }
    style3, err3 := vopi_xlsx.NewStyle(&excelize.Style{
        Font:      &excelize.Font{Family: "Times New Roman", Size: 11},
        Alignment: &excelize.Alignment{Horizontal: "left", Vertical: "center", WrapText: false},
    })
    if err3 != nil {
        fmt.Println(err_txt, "з іменами:", err3)
        os.Exit(7)
    }
    style4, err4 := vopi_xlsx.NewStyle(&excelize.Style{
        Font:      &excelize.Font{Family: "Times New Roman", Size: 11},
        Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: false},
    })
    if err4 != nil {
        fmt.Println(err_txt, "з підрозділом:", err4)
        os.Exit(8)
    }
    style5, err5 := vopi_xlsx.NewStyle(&excelize.Style{
        Font:      &excelize.Font{Family: "Times New Roman", Size: 11},
        Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center", WrapText: true},
    })
    if err5 != nil {
        fmt.Println(err_txt, "з телефоном:", err5)
        os.Exit(9)
    }
	
	vopi_xlsx.SetCellStyle(vopi_sheet, coord_rank, coord_rank, style2)
	vopi_xlsx.SetCellStyle(vopi_sheet, coord_name, coord_name, style3)
	vopi_xlsx.SetCellStyle(vopi_sheet, coord_dep, coord_dep, style4)
	vopi_xlsx.SetCellStyle(vopi_sheet, coord_tel, coord_tel, style5)
	
```

```go

// припускаємо: f *excelize.File, sheet string, coord string, val interface{}
styleID, err := f.GetCellStyle(sheet, coord)
if err != nil {
    // якщо не вдалося — просто пишемо значення
    _ = f.SetCellValue(sheet, coord, val)
} else {
    if err := f.SetCellValue(sheet, coord, val); err != nil {
        return err
    }
    // знову застосувати стиль, щоби зберегти бордери/формат
    if err := f.SetCellStyle(sheet, coord, coord, styleID); err != nil {
        return err
    }
}
```