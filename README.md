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
	Views: [
		"nav"
	],
	Libs:[
		"jquery.js"
	],
	Models:[
		"person"
	]
	Builder // embed the builder
}

func (sub *SubscriberHome) ServerHTTP(r, w){

}
```

```html
<!DOCTYPE html>
<html lang="en">
	<head>
	    <title>Dashboard - Ceres HQ</title>
	    {{printCSS}}
	</head>
	<body>
		{{View "nav" .}}
		{{printJS}}
	</body>
</html>
```