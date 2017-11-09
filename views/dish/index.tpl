{{ template "head.html" .}}
<body>
	<div id="main" class="container">

<nav class="breadcrumb">
  <a class="breadcrumb-item" href="/">Home</a>
  <a class="breadcrumb-item" href="#" onclick="dish_add()">Create a Dish</a>
  <span class="breadcrumb-item active">{{.Title}}</span>
</nav>

			 <table class="table">
			 <thead class="thead-default">
			 		<tr>	
						 	<th>Avatar</th>
							<th> Name </th>
      				<th> First Class </th>
			      	<th> Ingredients </th>
     			  	<th> Operation</th>
		    	</tr>
				</thead>
				<tbody>
			{{ range .Dishes }}
			<tr>
				<td><div><img src="{{.Image1 | ThumbnailURL }}" />  </div></td>
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
			
			</tbody>
			</table>

</div>
</body>

{{template "foot.html" . }}