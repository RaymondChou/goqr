package main

import (
	_ "bytes"
	"flag"
	"fmt"
	"github.com/loggi/goqr/pkg"
	"image"
	_ "image/color"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var ch chan int

var qrinput = flag.String("data", "", "please use -data input qr data string.")

var server = flag.Bool("server", false, "Use Server")

var port = flag.Int("port", 8889, "Listening")

var debug = flag.Bool("debug", false, "Debug, may print sensitive data.")

var logo qr.Logo

const shown = 8

func main() {

	flag.Parse()

	img := loadImg("input/loggi_bit.png")
	msk := loadImg("input/loggi_mask.png")
	logo = qr.Logo{img, msk}

	if *server {

		http.HandleFunc("/", api)

		err := http.ListenAndServe(":"+fmt.Sprintf("%d", *port), nil)

		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		} else {
			log.Println("Server started at port: " + fmt.Sprintf("%d", *port))
		}

	} else {

		if *qrinput == "" {
			fmt.Println("please use -data input qr data string.")
			return
		}

		qrarray := strings.Split(*qrinput, ",")

		begintime := time.Now().Unix()

		ch = make(chan int)

		os.Mkdir("output", 0755)

		for i, qrdata := range qrarray {

			fmt.Println("QREncoding >>>>>> " + qrdata)

			go output(qrdata, i, true)

			<-ch

		}

		endtime := time.Now().Unix()

		fmt.Println("completed time in seconds : " + fmt.Sprintf("%d", endtime-begintime))

	}

}

func loadImg(path string) (image.Image) {
	f, _ := os.Open(path)
	defer f.Close()
	loaded, _, err := image.Decode(f)
	if err != nil {
		log.Println(err)
	}
	return loaded
}

func output(data string, i int, goroutine bool) {

	c, err := qr.Encode(data, qr.H, logo)

	if err != nil {
		log.Println(err)
	}

	pngdat := c.PNG()

	if true {
		ioutil.WriteFile("output/"+fmt.Sprint(i+1)+".png", pngdat, 0666)
	}

	// m, err := png.Decode(bytes.NewBuffer(pngdat))

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// gm := m.(*image.Gray)

	// scale := c.Scale
	// siz := c.Size
	// nbad := 0

	// for y := 0; y < scale*(8+siz); y++ {

	// 	for x := 0; x < scale*(8+siz); x++ {

	// 		v := byte(255)

	// 		if c.Black(x/scale-4, y/scale-4) {
	// 			v = 0
	// 		}

	// 		if gv := gm.At(x, y).(color.Gray).Y; gv != v {
	// 			fmt.Println("%d,%d = %d, want %d", x, y, gv, v)
	// 			if nbad++; nbad >= 20 {
	// 				fmt.Println("too many bad pixels")
	// 			}
	// 		}

	// 	}

	// }
	if goroutine == true {
		ch <- 1
	}
}

func api(w http.ResponseWriter, r *http.Request) {

	var valid bool

	r.ParseForm()

	for k, v := range r.Form {

		if k == "data" {
			var inputed = strings.Join(v, "")
			if !*debug {
				inputed = mask(inputed)
			}
			log.Printf("%v: %v\n", k, inputed)

			data := strings.Join(v, "")

			c, err := qr.Encode(data, qr.H, logo)
			if err != nil {
				log.Println(err)
			}
			pngdat := c.PNG()

			w.Header().Set("Content-Type", "image/png")
			w.Write(pngdat)

			defer func() {
				valid = true
			}()
		}

	}

	if valid == false {
		fmt.Fprintf(w, "Please input `data` using get method!")
	}
}


func mask(inputed string) string {
	var res = make([]byte, len(inputed))
	for i := range res {
		if i < shown {
			res[i] = inputed[i]
		} else {
			res[i] = '*'
		}
	}
	return string(res)
}