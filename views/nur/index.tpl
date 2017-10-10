{{ template "head.html" .}}

<h1>{{.Title }} </h1>

{{.Count}}
<br />
<button onclick="add_nur()" > Add Nurition </button>



<table border="1">
<tr>
<th> Name </th>
<th> Everyday Dosage </th>
<th> Indication </th>
<th> Operation</th>
</tr>

{{ range .Nurs }}
<tr>
<td> <div> {{.Name }} </div> </td>
<td> <div> {{.Everyday}} mg </div></td>
<td><div> {{.Indication }} </div></td>
<td><div> <a href="#" onclick="delete_nur({{.Id}})">Delete </a> || <a href="#" onclick="edit_nur({{.Id}});"> Edit </a> </div> </td>
</tr>

{{end}}
</table>

{{template "foot.html" . }}