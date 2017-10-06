{{ template "head.html" .}}

<h1>{{.Title }} </h1>

{{.Count}}
<br />
<button onclick="add_nur()" > Add Nurition </button>



{{template "foot.html" . }}