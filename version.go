package sensehat

const (
	// SenseHatFbName is the device name
	SenseHatFbName = "RPi-Sense FB"
	// SENSE_HAT_FB_FBIOGET_GAMMA = 61696
	// SENSE_HAT_FB_FBIOSET_GAMMA = 61697
	// SENSE_HAT_FB_FBIORESET_GAMMA = 61698
	// SENSE_HAT_FB_GAMMA_DEFAULT = 0
	// SENSE_HAT_FB_GAMMA_LOW = 1
	// SENSE_HAT_FB_GAMMA_USER = 2

)

type PixelMapType map[int][][]int
type PixelList []*Pixel

var (
	// Version number
	Version = "0.1.0"

	PixelMap0 = [][]int{
		[]int{0, 1, 2, 3, 4, 5, 6, 7},
		[]int{8, 9, 10, 11, 12, 13, 14, 15},
		[]int{16, 17, 18, 19, 20, 21, 22, 23},
		[]int{24, 25, 26, 27, 28, 29, 30, 31},
		[]int{32, 33, 34, 35, 36, 37, 38, 39},
		[]int{40, 41, 42, 43, 44, 45, 46, 47},
		[]int{48, 49, 50, 51, 52, 53, 54, 55},
		[]int{56, 57, 58, 59, 60, 61, 62, 63},
	}

	PixelMap = PixelMapType{
		0:   PixelMap0,
		90:  PixelMap0,
		180: PixelMap0,
		270: PixelMap0,
	}
)
