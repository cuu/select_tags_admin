<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, max    imum-scale=1.0, user-scalable=no">

	{{compress_css "lib"}}
	{{compress_css "app"}}
	{{compress_js  "lib"}}
	{{compress_js  "app"}}

</head>

<body>
  <div class="container">
		<h1 class="logo">Welcome</h1>
		
    <div class="description">
		<div>
			{{ if .InPost }}
			{{.PostData}}
			{{end }}
		</div>
			
    </div>

		<div>
			<div> <a href="http://127.0.0.1:8081">Front Page</a></div>
			<hr />
			<br />
			<ul class="list-group">
				<li class="list-group-item"><a href="/menu"> Menus </a></li>
				<li class="list-group-item"><a href="/dish"> Dishes </a></li>
				<li class="list-group-item"><a href="/ingredient">  Ingredients </a></li>
				<li class="list-group-item"><a href="/nur">  Nutritions </a></li>
			</ul>
		</div>
		
			
  </div>
 
		
</body>
</html>
