{{ template "head.html" . }}


<div id="main" class="container">
<div class="row">
		 <div id="content">

<form method="POST" action="/menu/edit/{{.Id}}">
{{ .xsrf_html }} {{ .once_html }}

{{template "form/fields.html"  .MenuFormSets}}


<button class="btn btn-primary"> Update <i class="icon-chevron-sign-right"></i></button>

</form>

		</div>
</div>
</div>

{{ template "foot.html" . }}