package generator

import "testing"

func TestGeneratFromURL(t *testing.T) {
	t.Run("TestGeneratFromURL", func(t *testing.T) {
		url := "https://www.baidu.com/"
		if err := GeneratFromURL(url, "./url.pdf"); err != nil {
			t.Errorf("TestGeneratFromURL() error = %v", err)
		}
	})
}

func TestGeneratFromHTML(t *testing.T) {
	t.Run("TestGeneratFromHTML", func(t *testing.T) {
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="utf-8">
			<title>菜鸟教程(runoob.com)</title>
		</head>
		<body>
		<h4><a href="https://www.baidu.com/">百度一下</a></h4>
		<img border="0" src="https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png" alt="Pulpit rock" width="304" height="228">
		</body>
		</html>`
		if err := GeneratFromHTML(html, "./html.pdf"); err != nil {
			t.Errorf("GeneratFromHTML() error = %v ", err)
		}
	})
}
