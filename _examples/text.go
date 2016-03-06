package main

import (
	"log"
	"time"

	"github.com/mephux/sensehat"
)

func main() {
	device, err := sensehat.GetDevice()

	if err != nil {
		log.Fatal(err)
	}

	// spew.Dump(device)

	// device.Pack(sensehat.Pixel{255, 255, 255})

	// pp := sensehat.Pixel{56, 252, 0}
	// fmt.Println(pp.Pack())

	// p, err := device.GetPixel(4, 2)

	// if err != nil {
	// log.Fatal(err)
	// }

	// spew.Dump(p)

	// p, err = device.SetPixel(4, 2, &sensehat.Pixel{
	// R: 0,
	// G: 0,
	// B: 0,
	// })

	// if err != nil {
	// log.Fatal(err)
	// }

	// spew.Dump(p)

	// p, err = device.GetPixel(4, 2)

	// if err != nil {
	// log.Fatal(err)
	// }

	// spew.Dump(p)

	// device.SetAllPixels(&sensehat.Pixel{
	// R: 0,
	// G: 0,
	// B: 0,
	// })

	device.SetAllPixels(&sensehat.Pixel{
		R: 0,
		G: 0,
		B: 0,
	})

	var y = 0
	for range time.Tick(time.Second * 10) {
		for i := 0; i < 8; i++ {

			rr, gg, bb := colorful.FastHappyColor().RGB255()
			device.SetPixel(y, i, &sensehat.Pixel{
				R: int(rr),
				G: int(gg),
				B: int(bb),
			})

			time.Sleep(1 * time.Second)
		}

		y++
		// device.SetAllPixels(&sensehat.Pixel{
		// R: 0,
		// G: 0,
		// B: 0,
		// })
	}

	// if list, err := device.GetPixels(); err != nil {
	// log.Fatal(err)
	// } else {

	// go func() {
	// for range time.Tick(time.Second * 3) {
	// fmt.Println("CLEAR ALL")
	// device.SetAllPixels(&sensehat.Pixel{
	// R: 0,
	// G: 0,
	// B: 0,
	// })
	// }
	// }()

	// for range time.Tick(time.Second * 5) {
	// fmt.Println("SET COLOR")
	// if err := device.SetPixels(list); err != nil {
	// fmt.Println("ERROR::::", err)
	// }

	// }
	// }
}
