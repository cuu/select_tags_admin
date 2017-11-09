{{ template "head.html" .}}
<body>
<div id="main" class="container">
<nav class="breadcrumb">
  <a class="breadcrumb-item" href="/">Home</a>
  <a class="breadcrumb-item" href="#" onclick="add_nur()">Add Ingredient</a>
  <span class="breadcrumb-item active">{{.Title}}</span>
</nav>

<table class="table">
<thead class="thead-default">
<tr>
<th> Name </th>
<th> Everyday Dosage </th>
<th> Indication </th>
<th> Operation</th>
</tr>
</thead>
<tbody>
{{ range .Nurs }}
<tr>
<td> <div> {{.Name }} </div> </td>
<td> <div> {{.Everyday}} mg </div></td>
<td><div> {{.Indication }} </div></td>
<td><div> <a href="#" onclick="delete_nur({{.Id}})">Delete </a> || <a href="#" onclick="edit_nur({{.Id}});"> Edit </a> </div> </td>
</tr>

{{end}}
</tbody>
</table>

{{template "foot.html" . }}