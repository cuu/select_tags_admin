<head>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0, max    imum-scale=1.0, user-scalable=no">

<meta name="_xsrf" content="{{.xsrf_token}}" />

{{compress_css "lib"}}
{{compress_css "app"}}
{{compress_js  "lib"}}
{{compress_js  "app"}}

{{compress_js  "admin"}}
{{compress_js  "user"}}


{{compress_blueimp_js "lib"}}
{{compress_blueimp_css "lib"}}


{{compress_jqueryupload_css "lib"}}
{{compress_jqueryupload_js "lib"}}


<!-- CSS adjustments for browsers with JavaScript disabled -->
<noscript><link rel="stylesheet" href="static/thirdpart/jquery-file-upload/css/jquery.fileupload-noscript.css"></noscript>
<noscript><link rel="stylesheet" href="static/thirdpart/jquery-file-upload/css/jquery.fileupload-ui-noscript.css"></noscript>


{{assets_js "/static/js/dish/main.js"}}

</head>


<div id="main" class="container">
<div class="row">
		 <div id="content">

<form id="fileupload" method="POST" action="/dish/add" enctype="multipart/form-data">
{{ .xsrf_html }} {{ .once_html }}

{{template "form/fields.html"  .DishFormSets}}


        <!-- Redirect browsers with JavaScript disabled to the origin page -->
        <noscript><input type="hidden" name="redirect" value="https://blueimp.github.io/jQuery-File-Upload/"></noscript>
        <!-- The fileupload-buttonbar contains buttons to add/delete files and start/cancel the upload -->
        <div class="row fileupload-buttonbar">
            <div class="col-lg-7">
                <!-- The fileinput-button span is used to style the file input field as button -->
                <span class="btn btn-success fileinput-button">
                    <i class="glyphicon glyphicon-plus"></i>
                    <span>添加文件</span>
                    <input type="file" name="files[]" multiple>
                </span>
                <button type="submit" class="btn btn-primary start">
                    <i class="glyphicon glyphicon-upload"></i>
                    <span>Start upload</span>
                </button>
                <button type="reset" class="btn btn-warning cancel">
                    <i class="glyphicon glyphicon-ban-circle"></i>
                    <span>Cancel upload</span>
                </button>
                <button type="button" class="btn btn-danger delete">
                    <i class="glyphicon glyphicon-trash"></i>
                    <span>Delete</span>
                </button>
                <input type="checkbox" class="toggle">
                <!-- The global file processing state -->
                <span class="fileupload-process"></span>
            </div>
            <!-- The global progress state -->
            <div class="col-lg-5 fileupload-progress fade">
                <!-- The global progress bar -->
                <div class="progress progress-striped active" role="progressbar" aria-valuemin="0" aria-valuemax="100">
                    <div class="progress-bar progress-bar-success" style="width:0%;"></div>
                </div>
                <!-- The extended global progress state -->
                <div class="progress-extended">&nbsp;</div>
            </div>
        </div>
        <!-- The table listing the files available for upload/download -->
        <table role="presentation" class="table table-striped"><tbody class="files"></tbody></table>




<button class="btn btn-primary"> 添加  <i class="icon-chevron-sign-right"></i></button>

</form>

		</div>
</div>
</div>



<!-- The blueimp Gallery widget -->
<div id="blueimp-gallery" class="blueimp-gallery blueimp-gallery-controls" data-filter=":even">
    <div class="slides"></div>
    <h3 class="title"></h3>
    <a class="prev">‹</a>
    <a class="next">›</a>
    <a class="close">×</a>
    <a class="play-pause"></a>
    <ol class="indicator"></ol>
</div>
<!-- The template to display files available for upload -->
<script id="template-upload" type="text/x-tmpl" >
{{str2html .template_upload}}
</script>

<!-- The template to display files available for download -->
<script id="template-download" type="text/x-tmpl">
{{str2html .template_download}}
</script>


{{ template "foot.html" . }}

