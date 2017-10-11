{{ template "head.html" . }}


<form method="POST" action="/nur/edit/{{.Id}}">
{{ .xsrf_html }} {{ .once_html }}

{{template "form/fields.html" .NurFormSets }}

<button class="btn btn-primary"> Submit <i class="icon-chevron-sign-right"></i></button>

</form>


{{ template "foot.html" . }}
