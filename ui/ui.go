package ui

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"

	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ktappdev/filesync/models"
	"github.com/ktappdev/filesync/ui/topbar"
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
	ui.window = ui.app.NewWindow("Ableton Livesync")
	ui.window.Resize(fyne.NewSize(800, 500))

	searchEntry := widget.NewEntry()
	searchEntry.ActionItem = widget.NewButtonWithIcon("", theme.CancelIcon(), func() { searchEntry.SetText("") })
	searchEntry.SetPlaceHolder("Search Projects...")
	searchEntry.OnChanged = func(text string) {
		ui.updateFileList(text)
	}

	ui.fileList = widget.NewList(
		func() int {
			return len(ui.filteredFiles)
		},
		func() fyne.CanvasObject {
			icon := widget.NewIcon(resourceAbletonIcon512Jpg)
			label := widget.NewLabel("")
			labelNumber := widget.NewLabel("")

			return container.NewHBox(labelNumber, icon, label)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// Since the object is now a container, we need to get the label part of it to set the text
			container := o.(*fyne.Container)
			labelNumber := container.Objects[0].(*widget.Label)
			labelNumber.Text = fmt.Sprint(i + 1)
			labelNumber.TextStyle.Bold = false
			projectLabel := container.Objects[2].(*widget.Label)
			projectLabel.SetText(ui.filteredFiles[i].Name)
			projectLabel.TextStyle.Bold = true
			labelNumber.Refresh()
		},
	)

	ui.fileList.OnSelected = func(id widget.ListItemID) {
		ui.updateDetailView(id)
	}
	numFiles := len(ui.files)

	infoLbl := widget.NewLabel(fmt.Sprintf("%d Projects", numFiles))
	ui.detailContainer = container.NewVBox()
	// Align all children to the center vertically

	scrollableFileList := container.NewVScroll(ui.fileList)

	listContainer := container.NewBorder(searchEntry, infoLbl, nil, nil, scrollableFileList)

	// Top bar
	topBar := topbar.DisplayTopBar()

	split := container.NewHSplit(listContainer, ui.detailContainer)
	split.Offset = 0.5

	borderLayout := layout.NewBorderLayout(topBar, nil, nil, nil)
	allContent := container.New(borderLayout, topBar, split)
	ui.window.SetContent(allContent)

	// Initialize the file list with all files
	ui.updateFileList("")
}

// updateFileList filters the file list based on the search query.
func (ui *FileManagerUI) updateFileList(query string) {
	query = strings.ToLower(query)
	ui.filteredFiles = make([]models.FileInfo, 0, len(ui.files))

	for _, file := range ui.files {
		if strings.Contains(strings.ToLower(file.Name), query) {
			ui.filteredFiles = append(ui.filteredFiles, file)
		}
	}

	if len(ui.filteredFiles) == 0 {
		ui.fileList.Refresh()
		return
	}

	ui.fileList.Refresh()
	ui.detailContainer.Objects = nil
}

// updateDetailView updates the detail view based on the selected file.
func (ui *FileManagerUI) updateDetailView(id widget.ListItemID) {
	spacer := layout.NewSpacer()

	if id >= len(ui.filteredFiles) {
		return
	}
	file := ui.filteredFiles[id]

	// Clear existing content
	ui.detailContainer.Objects = nil

	projectName := canvas.NewText(file.Name, color.Black)
	projectName.TextSize = 12.0
	projectName.Alignment = fyne.TextAlignCenter
	projectName.TextStyle = fyne.TextStyle{Bold: true}
	projectNameWithPadding := container.NewPadded(projectName)

	// ui.detailContainer.Add(widget.NewSeparator())
	fileSizeFloat := float64(file.Size) / (1024 * 1024)
	formattedSize := fmt.Sprintf("%.2f MB", fileSizeFloat) // (fileSizeFloat)
	fileSize := container.NewGridWithColumns(3, widget.NewLabel("Size:"), widget.NewLabel(formattedSize), spacer)

	bpmLabel := widget.NewLabel(fmt.Sprintf("%.2f", file.BPM))
	bppmHBox := container.NewGridWithColumns(3, widget.NewLabel("BPM:"), bpmLabel, spacer)

	releaseDate := widget.NewLabel(file.UpdatedAt.Format("Jan 02, 2006"))
	releasedHBox := container.NewGridWithColumns(3, widget.NewLabel("Release Date:"), releaseDate, spacer)

	createdAt := widget.NewLabel(file.CreatedAt.Format("Jan 02, 2006"))
	createdAtHBox := container.NewGridWithColumns(3, widget.NewLabel("Created At:"), createdAt, spacer)

	updatedAt := widget.NewLabel(file.UpdatedAt.Format("Jan 02, 2006"))
	updateHBox := container.NewGridWithColumns(3, widget.NewLabel("Updated At:"), updatedAt, spacer)

	genreSelect := widget.NewSelect([]string{"Hip-Hop", "Jazz"}, func(value string) { file.Genre = value })
	genreSelect.SetSelected(file.Genre)
	genreS := container.NewGridWithColumns(3, widget.NewLabel("Genre:"), genreSelect, spacer)

	statusSelect := widget.NewSelect([]string{"WIP", "Upcoming"}, func(value string) { file.Status = value })
	statusSelect.SetSelected(file.Status)
	statusS := container.NewGridWithColumns(3, widget.NewLabel("Status:"), statusSelect, spacer)

	gradeSelect := widget.NewSelect([]string{"S", "D"}, func(value string) { file.Grade = value })
	gradeSelect.SetSelected(file.Grade)
	gradeS := container.NewGridWithColumns(3, widget.NewLabel("Grade:"), gradeSelect, spacer)

	keyLabel := widget.NewLabel(fmt.Sprint(file.Key, "Major"))
	keyS := container.NewGridWithColumns(3, widget.NewLabel("Key:"), keyLabel, spacer)

	innerDetailsContainer := container.NewVBox(
		keyS,
		fileSize,
		bppmHBox,
		releasedHBox,
		createdAtHBox,
		updateHBox,
		genreS,
		statusS,
		gradeS,
	)
	// detailsRectangle := canvas.NewRectangle(&color.RGBA{R: 192, G: 192, B: 192, A: 255})
	detailsBorder := container.NewBorder(projectNameWithPadding, nil, nil, nil, innerDetailsContainer)
	// detailsStack := container.NewVSplit(detailsRectangle, detailsBorder)
	ui.detailContainer.Add(detailsBorder)

	// Refresh the container to display the new content
	ui.detailContainer.Refresh()

}

// Run starts the application, displaying the main window.
func (ui *FileManagerUI) Run() {
	ui.window.ShowAndRun()
}

// Implementing the required method for secure restorable state
func (ui *FileManagerUI) ApplicationSupportsSecureRestorableState() bool {
	return true
}
