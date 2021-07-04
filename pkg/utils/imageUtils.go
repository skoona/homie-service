package utils

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

func SknSelectResource(alias string) fyne.Resource {
	return SknImageByName(alias, false).Resource
}
func SknSelectThemedResource(alias string) fyne.Resource {
	return SknImageByName(alias, true).Resource
}

func SknSelectImage(alias string) *canvas.Image {
	return SknImageByName(alias, false)
}
func SknSelectThemedImage(alias string) *canvas.Image {
	return SknImageByName(alias, true)
}

func SknImageByName(alias string, themed bool) *canvas.Image {
	var selected fyne.Resource

	switch (alias) {
	case "homieBanner":
		selected = resourceHomieBannerPng
	case "insertDriveFile_o":
		selected = resourceInsertDriveFileMdo24pxSvg
	case "insertDriveFile_r":
		selected = resourceInsertDriveFileMdr24pxSvg
	case "notificationAlert_o":
		selected = resourceNotificationAlertMbo24pxSvg
	case "notificationAlert_r":
		selected = resourceNotificationAlertMbr24pxSvg
	case "sensorOff_o":
		selected = resourceSensorsOffMbo24pxSvg
	case "sensorOff_r":
		selected = resourceSensorsOffMbr24pxSvg
	case "sensorOn_o":
		selected = resourceSensorsOnMbo24pxSvg
	case "sensorOn_r":
		selected = resourceSensorsOnMbr24pxSvg
	case "thumbsDown_o":
		selected = resourceThumbsDownMdo24pxSvg
	case "thumbsDown_r":
		selected = resourceThumbsDownMdr24pxSvg
	case "thumbsDown_s":
		selected = resourceThumbsDownMds24pxSvg
	case "thumbsUp_o":
		selected = resourceThumbsUpMdo24pxSvg
	case "thumbsUp_r":
		selected = resourceThumbsUpMdr24pxSvg
	case "thumbsUp_s":
		selected = resourceThumbsUpMds24pxSvg
	case "timeLapse_o":
		selected = resourceTimeLapseMbo24pxSvg
	case "timeLapse_r":
		selected = resourceTimeLapseMbr24pxSvg
	case "toggleOff_o":
		selected = resourceToggleOffMbo24pxSvg
	case "toggleOff_r":
		selected = resourceToggleOffMbr24pxSvg
	case "toggleOn_o":
		selected = resourceToggleOnMbo24pxSvg
	case "toggleOn_r":
		selected = resourceToggleOnMbr24pxSvg
	case "wallClock":
		selected = resourceWallClockSvgrepoComSvg
	case "worksWithHomie":
		selected = resourceWorksWithHomieSvg
	default:
		selected = resourceHomieBannerPng
	}
	image := canvas.NewImageFromResource(selected)
	if themed {
		image.Resource = theme.NewThemedResource(selected)
	}
	image.FillMode = canvas.ImageFillContain
	image.ScaleMode = canvas.ImageScaleSmooth
	return image
}

// SknImageResourceFromPath loads a image, icon, or svg from filesystem
func SknImageResourceFromPath(path string) fyne.Resource {
	uri, err := storage.LoadResourceFromURI(storage.NewFileURI(LocateFile(path)))
	if err != nil {
		fmt.Println("Error opening image:",path, err.Error())
		return nil
	}
	return uri
}

// SknThemedCanvasImageFromPath loads a image, icon, or svg from filesystem with Themeing
func SknThemedCanvasImageFromPath(path string) *canvas.Image  {
	return SknLoadImageFromURI(storage.NewFileURI( LocateFile(path)), true).(*canvas.Image)
}

// SknImageFromPath loads a image, icon, or svg from filesystem
func SknImageFromPath(path string) fyne.CanvasObject {
	return SknLoadImageFromURI(storage.NewFileURI(path), false)
}

// SknLoadImageFromURI loads an image or svg from filesystem
func SknLoadImageFromURI(u fyne.URI, themed bool) fyne.CanvasObject {
	read, err := storage.Reader(u)
	if err != nil {
		fmt.Println("Error opening image", err)
		return canvas.NewRectangle(color.Black)
	}
	res, err :=	storage.LoadResourceFromURI(read.URI())
	if err != nil {
		fmt.Println("Error reading image", err)
		return canvas.NewRectangle(color.Black)
	}
	img := canvas.NewImageFromResource(res)
	img.FillMode = canvas.ImageFillContain
	if themed {
		img.Resource = theme.NewThemedResource(img.Resource)
	}
	return img
}

