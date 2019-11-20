package main

import (
	"github.com/zserge/lorca"
	"log"
	"net/url"
)

var (
	viewtitle      = "視窗標題"
	inputlabel     = "標題"
	bottonname     = "按鈕名稱"
	vieweight  int = 520
	viewheight int = 320

	ui, err = lorca.New("", "", vieweight, viewheight)
)

func main() {
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

//in call
func buildhtml() string {
	htmlstr := `<html>
	<head>
	<title>` + viewtitle + `</title>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	</head>
		<body>			
			<div class="field half"> `

	htmlend := `<div class="done"></div>		
	</body>
	</html> 
	`

	return htmlstr +
		`
	 	<label for="name" style="font-size:16px;">` + inputlabel + `</label>
		<input id="input" type="text" value=""  SIZE=40  height="35" style="font-size:16px;">
		</div>
		<input type="button" onclick="execute(document.querySelector('#input').value)" style="width:100px;height:30px;font-size:16px;" value="` + bottonname + `">
	` + htmlend
}
