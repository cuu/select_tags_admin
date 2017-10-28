{{ template "head.html" .}}
<body>
	<div id="main" class="container">
			 <h1>{{.Title }} </h1>

			 <a href="/" >Home</a> <br />

			 {{.Count}}

			 <br />
			 <button onclick="menu_add()" > Create a menu </button>



			 <table border="1">
			 <tr>
				<th> Date </th>
				<th> Booked </th>
				<th> Extras </th>
				<th> Operations </th>
			</tr>

			{{ range .Menus }}
			<tr>
				<td> <div> {{.SpecifyDate | datetime }} </div> </td>
				<td>
					<div>
								{{with .Booked }}
									 {{.String}}
								{{end}}
					</div>
				</td>
				<td>
					<div>
								{{with .Extras }}
											 {{.String}}
								{{end}}
					</div>
				</td>
				<td><div> <a href="#" onclick="menu_delete({{.Id}})">Delete </a> or <a href="#" onclick="menu_edit({{.Id}});"> Edit </a> </div> </td>
			</tr>

			{{end}}
			</table>

</div>
</body>

{{template "foot.html" . }}