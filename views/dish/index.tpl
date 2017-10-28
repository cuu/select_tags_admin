{{ template "head.html" .}}
<body>
	<div id="main" class="container">
			 <h1>{{.Title }} </h1>

			 <a href="/" >Home</a> <br />

			 {{.Count}}

			 <br />
			 <button onclick="dish_add()" > Add dish </button>



			 <table border="1">
			 <tr>
				<th> Name </th>
				<th> First Class </th>
				<th> Ingredients </th>
				<th> Operation</th>
			</tr>

			{{ range .Dishes }}
			<tr>
				<td> <div> {{.Name }} </div> </td>
				<td> <div> {{.FirstClass }}</div></td>
				<td>
					<div>
								{{with .Ingredients }}
									 {{.String}}
								{{end}}
					</div>
				</td>
				<td><div> <a href="#" onclick="dish_delete({{.Id}})">Delete </a> || <a href="#" onclick="dish_edit({{.Id}});"> Edit </a> </div> </td>
			</tr>

			{{end}}
			</table>

</div>
</body>

{{template "foot.html" . }}