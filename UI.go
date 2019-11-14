package main

import (
	"github.com/zserge/lorca"
	"log"
	"net/url"
)

var (
	viewtitle = "Title"

	
	inputlabel     = "輸入"
	bottonname        = "執行"
	vieweight  int = 520
	viewheight int = 320
	
)

func main() {

	Inputtitle("視窗標題")	
	Inputlabel("標題")
	Inputbotton("按鈕名稱")
	Makeview(500, 400)

	ui, err := lorca.New("", "", vieweight, viewheight)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	html := buildhtml()

	ui.Load("data:text/html," + url.PathEscape(html))
	ui.Bind("execute", func(input string) {
		ui.Eval(`document.querySelector('.done').innerText = '` + input + ` done'`)
	})
	<-ui.Done()
}

//out Call

func Inputtitle(title string){
	viewtitle = title
}

func Makeview(width int, height int) {
	vieweight = width
	viewheight = height
}

func Inputlabel(label string) {
	inputlabel = label
}
func Inputbotton(botton string) {
	bottonname = botton
}

//in call
func buildhtml() string {
	htmlstr := `<html>
	<head>
	<title>`+viewtitle+`</title>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	</head>
		<body>			
			<div class="field half"> `

	htmlend := `<div class="done"></div>		
	</body>
	</html> 
	`

	return htmlstr+
	 `
	 	<label for="name" style="font-size:16px;">` + inputlabel + `</label>
		<input id="input" type="text" value=""  SIZE=40  height="35" style="font-size:16px;">
		</div>
		<input type="button" onclick="execute(document.querySelector('#input').value)" style="width:100px;height:30px;font-size:16px;" value="` + bottonname + `">
	`+ htmlend
}
