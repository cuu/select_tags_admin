function menu_add() {
		window.location.href="/menu/add";
}

function menu_edit(id) {
		window.location.href="/menu/edit/"+id;
}

function menu_delete(id){
		var r = confirm("Delete Menu ?");
		if (r == true) {
				window.location.href="/menu/delete/"+id;
		}
}



