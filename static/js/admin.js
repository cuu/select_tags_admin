(function($){

		$(function(){
//				$('[rel=select2-admin-model]').each(function(_,e){
				$('.select2-admin-model').each(function(_,e){
						var $e = $(e);
						console.log($e);
						var model = $e.data('model');
						$e.select2({
								width:null,
								minimumInputLength: 3,
								ajax: {
										url: '/model/select',
										type: 'POST',
										data: function (query, page) {
												return {
														'search': query,
														'model': model
												};
										},
										results: function(d){
												var results = [];
												if(d.success && d.data){
														var data = d.data;
														$.each(data, function(i,v){
																results.push({
																		'id': v[0],
																		'text': v[1]
																});
														});
												}
												return {'results': results};
										},
										cache:true
								},
								initSelection: function(elm, cbk){
										var id = parseInt($e.val(), 10);
										if(id){
												$.post('/model/pick', {'id': id, 'model': model}, function(d){
														if(d.success){
																if(d.data && d.data.length){
																		cbk({
																				'id': d.data[0],
																				'text': d.data[1]
																		});
																}
														}
												});
										}
								}
						});
				});
		});
		
})(jQuery);
