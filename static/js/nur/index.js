function add_nur() {
		window.location.href="/nur/add";
}

function edit_nur(id) {
		window.location.href="/nur/edit/"+id;
}

function delete_nur(id){
		var r = confirm("Delete Nur ?");
		if (r == true) {
				window.location.href="/nur/delete/"+id;
		}
}



