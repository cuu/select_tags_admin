{{ template "head.html" .}}
<body>
	<div id="main" class="container">	

<nav class="breadcrumb">
  <a class="breadcrumb-item" href="/">Home</a>
  <a class="breadcrumb-item" href="#" onclick="menu_add()">Create a Menu</a>
  <span class="breadcrumb-item active">{{.Title}}</span>
</nav>


			 <table class="table">
			 <thead class="thead-default">
			 <tr>
				<th> Date </th>
				<th> Booked </th>
				<th> Extras </th>
				<th> Operations </th>
			</tr>
			</thead>
			<tbody>
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
			</tbody>
			</table>

</div>
</body>

{{template "foot.html" . }}