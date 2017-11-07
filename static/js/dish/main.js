/*
 * jQuery File Upload Plugin JS Example
 * https://github.com/blueimp/jQuery-File-Upload
 *
 * Copyright 2010, Sebastian Tschan
 * https://blueimp.net
 *
 * Licensed under the MIT license:
 * https://opensource.org/licenses/MIT
 */

/* global $, window */

$(function () {
    'use strict';
		var hash = {};
		var server_url_base="http://localhost:9073";
		
		function refresh_filelist(){
				//$("#fileupload .up_file").remove();
				var index = 1;
				$.each(hash,function(idx,h) {
						if (h == true) {
								$("#fileupload #DishForm-Image"+index).val(idx.replace("delete","image") );
								//										$('#fileupload').append('<input class="up_file" type="text" name="fileName" value="'+idx.replace("delete","image")+'"  readonly/>');
						}
						index++;
				});				
		}
    // Initialize the jQuery File Upload widget:
    $('#fileupload').fileupload({
        // Uncomment the following to send cross-domain cookies:
        xhrFields: {withCredentials: true},
        url: server_url_base+'/files',
				// Enable file resume
				// Chunk size in bytes
				maxNumberOfFiles: 6,
				autoUpload:false,
				maxChunkSize: 1000000,
				acceptFileTypes: /(\.|\/)(jpe?g|png)$/i,
				change : function (e, data) {
						if(data.files.length >= 6){
								alert("Max 6 files are allowed")
								return false;
						}
				},
				add: function (e, data) {
						var that = this;
						$.ajax({
								url: server_url_base+'/resume',
								xhrFields: {withCredentials: true},
								data: {file: data.files[0].name}
						}).done(function(result) {
								var file = result.file;
								data.uploadedBytes = file && file.size;
								$.blueimp.fileupload.prototype.options.add.call(that, e, data);
						});
				},
				success:function(result,textStatus,jqXHR){
						
						//console.log(result);
						$.each(result.files,function(index,file) {
								hash[file.deleteUrl] = true
						});
						refresh_filelist();
						
				}
    });

    // Enable iframe cross-domain access via redirect option:
    $('#fileupload').fileupload(
        'option',
        'redirect',
        window.location.href.replace(
            /\/[^\/]*$/,
            '/cors/result.html?%s'
        )
    );

    if (window.location.hostname === 'blueimp.github.io') {
        // Demo settings:
        $('#fileupload').fileupload('option', {
            url: '//jquery-file-upload.appspot.com/',
            // Enable image resizing, except for Android and Opera,
            // which actually support image resizing, but fail to
            // send Blob objects via XHR requests:
            disableImageResize: /Android(?!.*Chrome)|Opera/
                .test(window.navigator.userAgent),
            maxFileSize: 999000,
            acceptFileTypes: /(\.|\/)(gif|jpe?g|png)$/i
        });
        // Upload server status check for browsers with CORS support:
        if ($.support.cors) {
            $.ajax({
                url: '//jquery-file-upload.appspot.com/',
                type: 'HEAD'
            }).fail(function () {
                $('<div class="alert alert-danger"/>')
                    .text('Upload server currently unavailable - ' +
                            new Date())
                    .appendTo('#fileupload');
            });
        }
    } else {
        // Load existing files:
        $('#fileupload').addClass('fileupload-processing');
        $.ajax({
            // Uncomment the following to send cross-domain cookies:
            //xhrFields: {withCredentials: true},
            url: $('#fileupload').fileupload('option', 'url'),
            dataType: 'json',
            context: $('#fileupload')[0]
        }).always(function () {
            $(this).removeClass('fileupload-processing');
        }).done(function (result) {
						console.log(result);
            $(this).fileupload('option', 'done')
                .call(this, $.Event('done'), {result: result});
        });
    }

		
		$('#fileupload').on("fileuploaddone", function (e, data) {
				/*
				$.each(data.files, function (index, file) {
						console.log(file);
				});
				*/
		});

		$('#fileupload').on("fileuploadcompleted", function (e, data) {
				/*
				$.each(data.files, function (index, file) {
						console.log(file);
				});
				*/
				
		});
		
		$("#fileupload").on("fileuploadadd",function(e,data) {
				/*
				console.log("fileuploadadd: ");
				console.log(data.files);
				console.log( $("#fileupload").serializeArray() );
				*/
		});
		
		$("#fileupload").on("fileuploaddestroyed",function(e,data) {
				
			
			
				hash[data.url]=false;
				refresh_filelist();
			
		});
		
});
