{{ template "head.html" . }}



<form method="POST" action="/nur/add">
{{ .xsrf_html }} {{ .once_html }}

{{template "form/fields.html" .AddNurFormSets }}

<button class="btn btn-primary"> Add <i class="icon-chevron-sign-right"></i></button>

</form>


{{ template "foot.html" . }}
