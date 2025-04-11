package app

import (
	// "optima-app/views"

	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	myApp := app.NewWithID("com.dohd.optima-app")
	myApp.Settings().SetTheme(theme.LightTheme())
	w := myApp.NewWindow("OPTIMA BUSINESS")
	wSize := fyne.NewSize(600, 400)
	w.Resize(wSize)

	// login form
	username := widget.NewEntry()
	username.SetPlaceHolder("Enter Username")
	password := widget.NewEntry()
	password.SetPlaceHolder("Enter Password")
	loginBtn := widget.NewButtonWithIcon("Login", theme.LoginIcon(), nil)
	loginHeader := canvas.NewText("Welcome", color.Black)
	loginHeader.Alignment = fyne.TextAlignCenter // Horizontal centering
	loginHeader.TextSize = 24
	form := container.NewVBox(
		container.NewPadded(loginHeader),
		username,
		password,
		loginBtn,
	)
	formBorder := canvas.NewRectangle(color.RGBA{200, 200, 200, 255}) // dark gray
	formBorder.CornerRadius = 5
	formBorder.SetMinSize(fyne.NewSize(wSize.Width * 0.5, wSize.Height * 0.5))
	borderedCtn := container.NewStack(
		formBorder, 
		container.NewHBox(
			container.NewGridWrap(fyne.NewSize(formBorder.MinSize().Width * 0.08, formBorder.MinSize().Height)),
			container.NewGridWrap(fyne.NewSize(formBorder.MinSize().Width * 0.84, formBorder.MinSize().Height), form),
			container.NewGridWrap(fyne.NewSize(formBorder.MinSize().Width * 0.08, formBorder.MinSize().Height)),
		),
	)


	// products & config
	// Sample product data
	type Product struct {
		Name  string
		Price float64
		Stock int
	}
	products := []Product{
		{"Laptop", 1200.50, 5},
		{"Smartphone", 699.99, 10},
		{"Tablet", 399.99, 7},
		{"Headphones", 129.99, 15},
		{"Smartwatch", 199.99, 8},
	}

	// Convert products to a 2D array (for table display)
	data := make([][]string, len(products))
	for i, p := range products {
		data[i] = []string{p.Name, "$" + strconv.FormatFloat(p.Price, 'f', 2, 64), strconv.Itoa(p.Stock)}
	}

	table := widget.NewTable(
		func() (int, int) { return len(data) + 1, 3 }, // Rows + Header, Columns
		func() fyne.CanvasObject {
			return widget.NewLabel("") // Cells as labels
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			if tci.Row == 0 {
				// Header row
				headers := []string{"Product", "Price", "Stock"}
				label.SetText(headers[tci.Col])
				label.TextStyle.Bold = true
			} else {
				// Data rows
				label.SetText(data[tci.Row-1][tci.Col])
			}
		},
	)

	// Set column widths (adjust values as needed)
	tableBorderSize := fyne.NewSize(wSize.Width * 0.7, wSize.Height * 0.7)
	table.SetColumnWidth(0, tableBorderSize.Width * 0.4) // Product Name
	table.SetColumnWidth(1, tableBorderSize.Width * 0.3)  // Price
	table.SetColumnWidth(2, tableBorderSize.Width * 0.1)  // Stock

	tableBorder := canvas.NewRectangle(theme.Color(theme.ColorNameBackground)) 
	tableBorder.CornerRadius = 5
	tableBorder.SetMinSize(tableBorderSize)
	tableCtn := container.NewStack(
		tableBorder, 
		container.NewBorder(nil, nil, nil, nil, table),
	)
	
	headerBorder := canvas.NewRectangle(color.Transparent)
	createItem := widget.NewButton("Create", nil)
	headerBorder.CornerRadius = 5
	headerBorder.SetMinSize(fyne.NewSize(wSize.Width * 0.7, createItem.MinSize().Height))
	headerCtn := container.NewStack(
		headerBorder, 
		container.NewHBox(
			widget.NewLabelWithStyle("Products & Services", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			layout.NewSpacer(),
			createItem,
		),
	)


	content := widget.NewLabel("Welcome! Select an option from the menu")
	contentArea := container.NewCenter(content)

	menu := container.NewVBox(
		widget.NewButtonWithIcon("Dashboard", theme.HomeIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(container.NewCenter(widget.NewLabel("Dashboard View")))
		}),
		widget.NewButtonWithIcon("Sales\t\t", theme.FileIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(container.NewCenter(widget.NewLabel("Sales View")))
		}),
		widget.NewButtonWithIcon("Products\t", theme.GridIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(
				container.NewVBox(
					headerCtn,
					container.NewGridWrap(fyne.NewSize(headerCtn.MinSize().Width, 20), layout.NewSpacer()),
					tableCtn,
				),
			)
		}),
		widget.NewButtonWithIcon("Expenses\t", theme.ListIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(container.NewCenter(widget.NewLabel("Expenses View")))
		}),
		widget.NewButtonWithIcon("Settings\t", theme.SettingsIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(container.NewCenter(widget.NewLabel("Settings View")))
		}),
		widget.NewButtonWithIcon("Profile\t", theme.AccountIcon(), func() {
			contentArea.RemoveAll()
			contentArea.Add(borderedCtn)
		}),
	)

	split := container.NewHSplit(menu, contentArea)
	split.SetOffset(0.2)

	w.SetContent(split)
	w.ShowAndRun()
}
