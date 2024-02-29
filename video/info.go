package video

import (
	"image"
	"time"
)

type Info struct {
	FPS       float64
	Rectangle image.Rectangle

	Device       interface{} `json:",omitempty"`
	StartedAt    *time.Time  `json:",omitempty"`
	NFrom        int         `json:",omitempty"`
	Grayscaled   bool        `json:",omitempty"`
	FPSDivider   *int        `json:",omitempty"`
	DPMConverted *float64    `json:",omitempty"`
}
