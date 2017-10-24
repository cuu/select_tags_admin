{{ template "head.html" .}}
<body>
	<div id="main" class="container">
			 <h1>{{.Title }} </h1>

			 <a href="/" >Home</a> <br />

			 {{.Count}}

			 <br />
			 <button onclick="ingredient_add()" > Add ingredient</button>



			 <table border="1">
			 <tr>
				<th> Name </th>
				<th> Nutritions </th>
				<th> Operation</th>
			</tr>

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
			</table>

</div>
</body>

{{template "foot.html" . }}