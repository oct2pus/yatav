package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/dhowden/tag"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// SIZE is the widn
const SIZE = 400

func main() {
	md := make(chan tag.Metadata)

	go func(md chan tag.Metadata) {
		for {
			cmd := "cmus-remote -Q | grep \"file\" | sed -e \"s:file ::\""
			out, err := exec.Command("bash", "-c", cmd).Output()
			check(err)
			f, err := os.Open(strings.Trim(string(out), "\n"))
			check(err)
			t, err := tag.ReadFrom(f)
			check(err)
			md <- t
			f.Close()
			time.Sleep(1 * time.Second)
		}
	}(md)

	gtk.Init(nil)
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	check(err)
	img, err := gtk.ImageNew()
	check(err)
	win.Add(img)
	win.SetDefaultSize(SIZE, SIZE)
	win.SetResizable(false)
	win.ShowAll()
	go func() {
		for {
			var pic []byte
			pbl, err := gdk.PixbufLoaderNew()
			check(err)
			meta := <-md
			pic = meta.Picture().Data
			win.SetTitle(meta.Title())
			pix, err := pbl.WriteAndReturnPixbuf(pic)
			check(err)
			pix, err = pix.ScaleSimple(SIZE, SIZE, gdk.INTERP_BILINEAR)
			check(err)
			img.SetFromPixbuf(pix)
			win.Resize(SIZE, SIZE)
			pbl.Close()
		}
	}()
	gtk.Main()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
