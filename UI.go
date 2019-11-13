package UI

import (
	_ "bytes"
	_ "fmt"
	_ "io"
	"io/ioutil"
	"log"
	_ "net/http"
	"net/url"
	_ "os"
	_ "strconv"
	_ "strings"
	"unsafe"

	"github.com/zserge/lorca"
)

var (
	fname string
)

func Html() {
	ui, err := lorca.New("", "", 520, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	html := makehtml()

	ui.Load("data:text/html," + url.PathEscape(html))
	ui.Bind("download", func(filename string) {
		ui.Eval(`document.querySelector('.done').innerText = '` + filename + ` done'`)
	})
	<-ui.Done()	
}

//call
func makehtml() string {

	input := `
			<label for="name" style="font-size:16px;">請輸入網址</label>
				<input id="URL" type="text" value=""  SIZE=40  height="35" style="font-size:16px;">
			</div>
			<input type="button" onclick="download(document.querySelector('#URL').value)" style="width:100px;height:30px;font-size:16px;" value="Download">
	`

	html1 := Read("html1")
	html2 := Read("html2")
	return html1 + input + html2
}

//讀取文件
func Read(filename string) string {
	html, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return String(html)
}

//將[]byte轉成string
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
