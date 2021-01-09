package main

import (
	"bytes"
	"html/template"
)

func GetUIHTML() string {
	html := `
<html>
<head>
<style>
html {
	background-color: #39424e;
}
h1, h2, h3 {
	margin-top: 2em;
	margin-bottom: .5em;
	font-family: didact gothic,sans-serif;
	opacity: .6;
}
body {
  	font-size: 16px;
	font-family: didact gothic,sans-serif;
	color: #fff;
	line-height: 2rem;
	letter-spacing: 1.5px;
	text-shadow: none;
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	align-items: center;
	opacity: 1;
}
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
}

</style>
</head>
<body>
<h1>{{.PageTitle}}</h1>
</ul>
</body>
</html>
`
	return html
}

type UIData struct {
	PageTitle string
}
func UIDisplay() string {

	tmpl := template.Must(template.New("myname").Parse(GetUIHTML()))

	data := UIData{
		PageTitle: "Go Website from template",
	}

	//tmpl.Execute(w, data)
	//this is for a simple string return
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, data)
	return tpl.String()
}