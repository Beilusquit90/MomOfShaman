// MomOfShaman project main.go
package main

import (
	"fmt"
	_ "image/png"
	"math/rand"
	"os/exec"
	"strconv"

	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/fogleman/gg"
)

func main() {
	xx := make(map[string]string)
	xx["Пиманов А.С."] = "Пиманов.png"
	xx["Агапкин В.А."] = "Агапкин.png"
	xx["Якимчев М.Л."] = "Якимчев.png"
	xx["Мархинов С.А."] = "Мархинов.png"
	xx["Ульев А.В."] = "Ульев.png"
	xx["Алатинов С.В."] = "Алатинов.png"
	xx["Кашановский В.В."] = "Кашановский.png"
	xx["Короктев А.А."] = "Короктев.png"
	xx["Смиреев С.С."] = "Смиреев.png"
	xx["Галчин А.А."] = "Галчин.png"
	fmt.Println("Hello World!")

	xlre, err := excelize.OpenFile("1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := [][]string{}

	for count := 3; len(xlre.GetCellValue("Sheet1", "F"+strconv.Itoa(count))) > 0; count += 5 {
		tdata := []string{}

		tdata = append(tdata, xlre.GetCellValue("Sheet1", "A"+strconv.Itoa(count)))
		tdata = append(tdata, xlre.GetCellValue("Sheet1", "B"+strconv.Itoa(count)))
		tdata = append(tdata, xlre.GetCellValue("Sheet1", "C"+strconv.Itoa(count)))
		tdata = append(tdata, xlre.GetCellValue("Sheet1", "D"+strconv.Itoa(count)))
		tdata = append(tdata, xlre.GetCellValue("Sheet1", "D"+strconv.Itoa(count+2)))
		tdata = append(tdata, xlre.GetCellValue("Sheet1", "F"+strconv.Itoa(count)))

		data = append(data, tdata)

	}
	p2, err := gg.LoadImage("img/2.png")
	if err != nil {
		fmt.Println(err)
	}
	p3, err := gg.LoadImage("img/1.png")
	if err != nil {
		fmt.Println(err)
	}
	p4, err := gg.LoadImage("img/3.png")
	if err != nil {
		fmt.Println(err)
	}
	p5, err := gg.LoadImage("img/4.png")
	if err != nil {
		fmt.Println(err)
	}
	/*p1, err := gg.LoadImage("img/3.png")
	if err != nil {
		fmt.Println(err)
	}
	p2, err := gg.LoadImage("img/3.png")
	if err != nil {
		fmt.Println(err)
	}*/
	count := 0
	for _, dd := range data {
		ph, err := gg.LoadImage("img/" + xx[dd[2]])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(dd)

		im, err := gg.LoadImage("img/Scan.jpg")
		if err != nil {
			fmt.Println(err)
		}
		const S = 1024

		dc := gg.NewContext(1654, 2338)
		dc.SetRGB(1, 1, 1)
		dc.Clear()
		dc.SetRGB(0, 0, 0)
		if err := dc.LoadFontFace("font/arial.ttf", 20); err != nil {
			fmt.Println(err)
		}
		dc.DrawRoundedRectangle(0, 0, 1654, 2338, 0)
		dc.DrawImage(im, 0, 0)
		dc.DrawStringAnchored("Дата: "+dd[0], 822, 267, 0, 0)
		dc.DrawStringAnchored(dd[5], 1265, 267, 0, 0)
		dc.DrawStringAnchored(dd[0], 95, 1180, 0, 0)
		dc.DrawStringAnchored(dd[0], 95, 1948, 0, 0)
		dc.DrawStringAnchored(dd[0], 821, 1180, 0, 0)
		dc.DrawStringAnchored(dd[0], 95, 1229, 0, 0)
		dc.DrawStringAnchored(dd[0], 503, 1230, 0, 0)
		dc.DrawStringAnchored(dd[0], 821, 1231, 0, 0)
		dc.DrawStringAnchored(dd[0], 1236, 1231, 0, 0)

		dc.DrawStringAnchored(dd[2], 95, 1507, 0, 0)
		dc.DrawImageAnchored(ph, 210, 1480, 0, 0.5)
		dc.DrawStringAnchored(dd[2], 821, 1507, 0, 0)
		dc.DrawImageAnchored(ph, 932, 1470, 0, 0.5)

		dc.DrawStringAnchored(dd[3], 95, 1133, 0, 0)
		dc.DrawStringAnchored(dd[4], 821, 1135, 0, 0)

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		rx := (r1.Intn(30))
		ry := (r1.Intn(30))
		dc.DrawImageAnchored(p2, 1420+rx, 1925+ry, 0.5, 0.5)

		dc.Clip()
		dc.SavePNG("data/" + strconv.Itoa(count) + "-2.png")

		im2, err := gg.LoadImage("img/Scan2.jpg")
		if err != nil {
			fmt.Println(err)
		}

		dc2 := gg.NewContext(1654, 2338)
		dc2.SetRGB(1, 1, 1)
		dc2.Clear()
		dc2.SetRGB(0, 0, 0)
		if err := dc2.LoadFontFace("font/arial.ttf", 20); err != nil {
			fmt.Println(err)
		}
		dc2.DrawRoundedRectangle(0, 0, 1654, 2338, 0)
		dc2.DrawImage(im2, 0, 0)

		dc2.DrawStringAnchored(dd[2], 95, 300, 0, 0)
		dc2.DrawImageAnchored(ph, 210, 280, 0, 0.5)
		r1 = rand.New(s1)
		rx = (r1.Intn(30))
		ry = (r1.Intn(30))
		dc2.DrawImageAnchored(p2, 1435+rx, 1505+ry, 0.5, 0.5)
		dc2.DrawImageAnchored(p3, 715+rx, 1505+ry, 0.5, 0.5)
		dc2.DrawStringAnchored(dd[0], 495, 1520, 0, 0)
		dc2.DrawStringAnchored(dd[0], 1255, 1520, 0, 0)
		dc2.DrawStringAnchored(dd[1], 1105, 475, 0, 0)
		r1 = rand.New(s1)
		rx = (r1.Intn(25))
		ry = (r1.Intn(25))
		dc2.DrawImageAnchored(p4, 246+rx, 1640+ry, 0.5, 0.5)
		r1 = rand.New(s1)
		rx = (r1.Intn(25))
		ry = (r1.Intn(25))
		dc2.DrawImageAnchored(p5, 987+rx, 1640+ry, 0.5, 0.5)

		dc2.Clip()
		dc2.SavePNG("data/" + strconv.Itoa(count) + ".png")
		count++
		out, err := exec.Command("convert", "data/*.png", "data/pdf/"+dd[5]+".pdf").Output()

		out2, err2 := exec.Command("/bin/sh", "-c", "sudo rm data/*.png").Output()

		fmt.Println(out, err, out2, err2)

	}
}
