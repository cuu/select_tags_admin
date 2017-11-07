<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, max    imum-scale=1.0, user-scalable=no">
	
  <style type="text/css">
    *,body {
    }

    body {
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
    margin-left: auto;
    margin-right: auto;
		marign:10px;
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
      z-index: -1;
      top: 0px;
      left: 0px;
    }
		ul.list{
		list-style:none;
		}
		ul.list li {
		cursor:pointer;
		marign-top:10px;
		margin-bottom:10px;
		font-size:26px;
		
		}
  </style>
</head>

<body>
  <header>
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
			<ul class="list">
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
