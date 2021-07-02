package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	cc "github.com/skoona/homie-service/pkg/utils"
	"image/color"
	"log"
)

var Green color.Color = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
//text1 := canvas.NewText("Hello", green)

// SknCanvasSVGImageFromPath loads as an canvas.Image
func SknCanvasSVGImageFromPath(filename string) *canvas.Image  {
	image := SknLoadImageFromURI(storage.NewFileURI( cc.LocateFile(filename))).(*canvas.Image)
	image.Resource = theme.NewThemedResource(image.Resource)
	image.FillMode = canvas.ImageFillContain
	return image
}

// SknLoadImageFromPath returns a CanvasObject
func SknLoadImageFromPath(path string) fyne.CanvasObject {
	return SknLoadImageFromURI(storage.NewFileURI(path))
}

// SknLoadImageFromURI returns a CanvasObject
func SknLoadImageFromURI(u fyne.URI) fyne.CanvasObject {
	read, err := storage.Reader(u)
	if err != nil {
		log.Println("Error opening image", err)
		return canvas.NewRectangle(color.Black)
	}
	res, err :=	storage.LoadResourceFromURI(read.URI())
	if err != nil {
		log.Println("Error reading image", err)
		return canvas.NewRectangle(color.Black)
	}
	img := canvas.NewImageFromResource(res)
	//img.FillMode = canvas.ImageFillContain
	img.FillMode = canvas.ImageFillOriginal
	return img
}
