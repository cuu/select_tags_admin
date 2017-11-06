<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    header {
      padding: 10px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
		  margin-top:30px;
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style>
</head>

<body>
  <header>
		{{ if .InPost }} 
    <h1 class="logo">Welcome to Beego Post Data Test Page</h1>
		{{else }}
		<h1 class="logo">Welcome to Beego</h1>
		{{end}}
		
    <div class="description">
		<div>
			{{ if .InPost }}
			{{.PostData}}
			{{end }}
		</div>
			
    </div>

		<div>
			<ul>
				<li><a href="/menu"> Menus </a></li>
				<li><a href="/dish"> Dishes </a></li>
				<li><a href="/ingredient">  Ingredients </a></li>
				<li><a href="/nur">  Nutritions </a></li>
			</ul>
		</div>
		
			
  </header>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
		
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
