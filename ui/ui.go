package ui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ktappdev/filesync/models" // Assuming the FileInfo struct and related logic are in the model package
)

// FileManagerUI holds the components and state for the File Manager UI.
type FileManagerUI struct {
	app             fyne.App
	window          fyne.Window
	files           []models.FileInfo
	filteredFiles   []models.FileInfo
	fileList        *widget.List
	detailContainer *fyne.Container
}

// NewFileManagerUI creates and initializes the File Manager UI.
func NewFileManagerUI(app fyne.App, files []models.FileInfo) *FileManagerUI {
	ui := &FileManagerUI{
		app:   app,
		files: files,
	}
	ui.setupUI()
	return ui
}

// setupUI initializes the UI components and layouts.
func (ui *FileManagerUI) setupUI() {
	ui.window = ui.app.NewWindow("File Manager")
	ui.window.Resize(fyne.NewSize(800, 600))
	// abletonIcon :=

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search files...")
	searchEntry.OnChanged = func(text string) {
		ui.updateFileList(text)
	}

	ui.fileList = widget.NewList(
		func() int {
			return len(ui.filteredFiles)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(abletonIcon)
			label := widget.NewLabel("")
			return container.NewHBox(icon, label)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// Since the object is now a container, we need to get the label part of it to set the text
			container := o.(*fyne.Container)
			label := container.Objects[1].(*widget.Label)
			label.SetText(ui.filteredFiles[i].Name)
		},
	)

	ui.fileList.OnSelected = func(id widget.ListItemID) {
		ui.updateDetailView(id)
	}

	ui.detailContainer = container.NewVBox()

	toolbar := container.NewHBox(
		widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {}),
		widget.NewButtonWithIcon("", theme.ContentCutIcon(), func() {}),
		widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {}),
		widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {}),
		widget.NewSeparator(),
		widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {}),
	)

	ui.detailContainer.Add(toolbar)

	scrollableFileList := container.NewVScroll(ui.fileList)

	listContainer := container.NewBorder(searchEntry, nil, nil, nil, scrollableFileList)

	// Top bar
	topBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MailComposeIcon(), func() {}),
		widget.NewToolbarAction(theme.HelpIcon(), func() {}),
		widget.NewToolbarAction(theme.AccountIcon(), func() {}),
	)

	split := container.NewHSplit(listContainer, ui.detailContainer)
	split.Offset = 0.3

	borderLayout := layout.NewBorderLayout(topBar, nil, nil, nil)
	allContent := container.New(borderLayout, topBar, split)
	ui.window.SetContent(allContent)

	// Initialize the file list with all files
	ui.updateFileList("")
}

// updateFileList filters the file list based on the search query.
func (ui *FileManagerUI) updateFileList(query string) {
	query = strings.ToLower(query)
	ui.filteredFiles = nil // Reset the filtered list

	for _, file := range ui.files {
		if strings.Contains(strings.ToLower(file.Name), query) {
			ui.filteredFiles = append(ui.filteredFiles, file)
		}
	}

	ui.fileList.Refresh()
}

// updateDetailView updates the detail view based on the selected file.
func (ui *FileManagerUI) updateDetailView(id widget.ListItemID) {
	if id >= len(ui.filteredFiles) {
		return
	}
	file := ui.filteredFiles[id]

	// Clear existing content
	ui.detailContainer.Objects = nil

	// Re-add the toolbar to the detail view
	ui.detailContainer.Add(container.NewHBox(
		widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {}),
		widget.NewButtonWithIcon("", theme.ContentCutIcon(), func() {}),
		widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {}),
		widget.NewButtonWithIcon("", theme.ContentPasteIcon(), func() {}),
		widget.NewSeparator(),
		widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {}),
	))

	// Add file details
	ui.detailContainer.Add(widget.NewLabel(fmt.Sprintf("Name: %s", file.Path)))
	ui.detailContainer.Add(widget.NewLabel(fmt.Sprintf("Size: %d", file.Size)))
	ui.detailContainer.Add(widget.NewLabel(fmt.Sprintf("BPM: %.2f", file.BPM)))

	// Genre selection
	genreSelect := widget.NewSelect([]string{"Hip-Hop", "Jazz"}, func(value string) { file.Genre = value })
	genreSelect.SetSelected(file.Genre)
	ui.detailContainer.Add(container.NewHBox(widget.NewLabel("Genre:"), genreSelect))

	// Status selection
	statusSelect := widget.NewSelect([]string{"WIP", "Upcoming"}, func(value string) { file.Status = value })
	statusSelect.SetSelected(file.Status)
	ui.detailContainer.Add(container.NewHBox(widget.NewLabel("Status:"), statusSelect))

	// Key selection
	keySelect := widget.NewSelect([]string{"C#", "Bb"}, func(value string) { file.Key = value })
	keySelect.SetSelected(file.Key)
	ui.detailContainer.Add(container.NewHBox(widget.NewLabel("Key:"), keySelect))

	// Grade selection
	gradeSelect := widget.NewSelect([]string{"S", "D"}, func(value string) { file.Grade = value })
	gradeSelect.SetSelected(file.Grade)
	ui.detailContainer.Add(container.NewHBox(widget.NewLabel("Grade:"), gradeSelect))

	ui.detailContainer.Add(widget.NewLabel(fmt.Sprintf("Release Date: %s", file.ReleaseDate)))
	ui.detailContainer.Add(widget.NewLabel(fmt.Sprintf("Created At: %s", file.CreatedAt.Format("Jan 02, 2006"))))
	ui.detailContainer.Add(widget.NewLabel(fmt.Sprintf("Last Updated At: %s", file.Modified.Format("Jan 02, 2006"))))

	// Refresh the container to display the new content
	ui.detailContainer.Refresh()
}

// Run starts the application, displaying the main window.
func (ui *FileManagerUI) Run() {
	ui.window.ShowAndRun()
}
