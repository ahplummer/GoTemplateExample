package main

import (
	"testing"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred: #{err}")
	}
}
func TestGetUIDisplay(t *testing.T) {
	actual := UIDisplay()
	expected := `
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
<h1>Go Website from template</h1>
</ul>
</body>
</html>
`
	if actual != expected {
		t.Errorf("Actual: \n%s\n=====\nExpected:\n%s",actual, expected)
	}
}


