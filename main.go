package main

import (
	"./pkg"
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var ch chan int

func main() {

	qrinput := flag.String("data", "", "please input qr data string.")

	flag.Parse()

	qrarray := strings.Split(*qrinput, ",")

	begintime := time.Now().Unix()

	ch = make(chan int)

	os.Mkdir("output", 0755)

	for i, qrdata := range qrarray {

		fmt.Println("QREncoding >>>>>> " + qrdata)

		go output(qrdata, i)

		<-ch

	}

	endtime := time.Now().Unix()

	fmt.Println("completed time in seconds : " + fmt.Sprintf("%d", endtime-begintime))
}

func output(data string, i int) {

	c, err := qr.Encode(data, qr.M)
	if err != nil {
		fmt.Println(err)
	}
	pngdat := c.PNG()
	if true {
		ioutil.WriteFile("output/"+fmt.Sprint(i+1)+".png", pngdat, 0666)
	}
	m, err := png.Decode(bytes.NewBuffer(pngdat))
	if err != nil {
		fmt.Println(err)
	}
	gm := m.(*image.Gray)

	scale := c.Scale
	siz := c.Size
	nbad := 0

	for y := 0; y < scale*(8+siz); y++ {
		for x := 0; x < scale*(8+siz); x++ {
			v := byte(255)
			if c.Black(x/scale-4, y/scale-4) {
				v = 0
			}
			if gv := gm.At(x, y).(color.Gray).Y; gv != v {
				fmt.Println("%d,%d = %d, want %d", x, y, gv, v)
				if nbad++; nbad >= 20 {
					fmt.Println("too many bad pixels")
				}
			}
		}
	}

	ch <- 1

}
