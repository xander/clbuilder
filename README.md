# clbuilder
client side asset builder for Go


Directory structure

```bash
root/
	pages/
		my_page.html
	views/
		nav/
			nav.html 
			nav.js
			nav.css
	libs/
		jquery.js
	models/
		person.js

```

Example:

```go

type MyPage struct {
	Template: "my_page.html"
	Views: []string{
		"nav"
	}
	Libs:[]string{
		"jquery.js"
	}
	Models:[]string{
		"person"
	}
}

func (mp *MyPage) ServerHTTP(w ResponseWriter, r *Request){
	var data interface{}	//	data to pass to the template

	Builder.Execute(w, data, mp.Template, mp.Views, mp.Libs, mp.Models)
}
```

```html
<!DOCTYPE html>
<html lang="en">
	<head>
	    <title>Dashboard - Ceres HQ</title>
	    {{PrintCSS}}
	</head>
	<body>
		{{View "nav" .}}
		{{PrintJS}}
	</body>
</html>
```