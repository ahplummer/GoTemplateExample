package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func setupMain(){
}
func teardownMain(){
}
func TestMain(m *testing.M){
	setupMain()
	code := m.Run()
	teardownMain()
	os.Exit(code)
}
func TestNewRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	req.RequestURI = "/"
	http.HandlerFunc(NewRootHandler).ServeHTTP(rr, req)
	expectedStatus := http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Status code differs. Expected %d, got %d", expectedStatus, status )
	}
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
	if rr.Body.String() !=  expected{
		t.Errorf("Got %s", rr.Body.String())
	}

}


