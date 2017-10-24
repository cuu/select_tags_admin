(function($){

		$(function(){
				$('.select2-admin-model,[rel=select2-admin-model] ').each(function(_,e){
						var $e = $(e);
						//console.log($e);
						var model = $e.data('model');
						$e.select2({
								placeholder:"search",
								delay:250,
								minimumInputLength: 2,
								ajax: {
										url: '/model/select',
										type: 'POST',
										
										data: function (params) {
												return {
														'search': params.term,
														'model': model
												};
										},
										processResults: function(data,params){
												var results = [];
												if(data.success && data.data){
														var data = data.data;
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
								} // end ajax
								/*
								initSelection: function(ele,callback) {
										//console.log(JSON.stringify(ele));
										
								}
								*/						 
								
						});
				});
		});


		$( ".select2-admin-model" ).on( "select2:open", function() {
        if ( $( this ).parents( "[class*='has-']" ).length ) {
						var classNames = $( this ).parents( "[class*='has-']" )[ 0 ].className.split( /\s+/ );
						
						for ( var i = 0; i < classNames.length; ++i ) {
								if ( classNames[ i ].match( "has-" ) ) {
										$( "body > .select2-container" ).addClass( classNames[ i ] );
								}
						}
        }
    });
		
})(jQuery);
