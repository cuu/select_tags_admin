{{ template "head.html" .}}
<body>
	<div id="main" class="container">
			 <h1>{{.Title }} </h1>

			 <a href="/" >Home</a> <br />

			 {{.Count}}

			 <br />
			 <button onclick="add_dish()" > Add dish </button>



			 <table border="1">
			 <tr>
				<th> Name </th>
				<th> Nurs </th>
				<th> Operation</th>
			</tr>

			{{ range .Dishes }}
			<tr>
				<td> <div> {{.Name }} </div> </td>
				<td> <div> {{.Nurs}} mg </div></td>
				<td><div> <a href="#" onclick="delete_dish({{.Id}})">Delete </a> || <a href="#" onclick="edit_dish({{.Id}});"> Edit </a> </div> </td>
			</tr>

			{{end}}
			</table>

</div>
</body>

{{template "foot.html" . }}