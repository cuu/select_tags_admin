{{ template "head.html" . }}


<div id="main" class="container">
<div class="row">
		 <div id="content">

<form method="POST" action="/dish/add">
{{ .xsrf_html }} {{ .once_html }}

{{template "form/fields.html"  .DishFormSets}}

<button class="btn btn-primary"> Add <i class="icon-chevron-sign-right"></i></button>

</form>

		</div>
</div>
</div>

{{ template "foot.html" . }}

