{{ template "head.html" .}}
<body>
	<div id="main" class="container">
<nav class="breadcrumb">
  <a class="breadcrumb-item" href="/">Home</a>
  <a class="breadcrumb-item" href="#" onclick="ingredient_add()">Add Ingredient</a>
  <span class="breadcrumb-item active">{{.Title}}</span>
</nav>

			 <table class="table">
			 <thead class="thead-default">
			 <tr>
				<th> Name </th>
				<th> Nutritions </th>
				<th> Operation</th>
			</tr>
			</thead>
			<tbody>
			{{ range .Ingredients }}
			<tr>
				<td> <div> {{.Name }} </div> </td>
				<td>
					<div>
								{{with .Nutritions }}
									 {{.String}}
								{{end}}
					</div>
				</td>
				<td><div> <a href="#" onclick="ingredient_delete({{.Id}})">Delete </a> || <a href="#" onclick="ingredient_edit({{.Id}});"> Edit </a> </div> </td>
			</tr>

			{{end}}
			</tbody>
			</table>

</div>
</body>

{{template "foot.html" . }}