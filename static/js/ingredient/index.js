function ingredient_add() {
		window.location.href="/ingredient/add";
}

function ingredient_edit(id) {
		window.location.href="/ingredient/edit/"+id;
}

function ingredient_delete(id){
		var r = confirm("Delete Dish ?");
		if (r == true) {
				window.location.href="/ingredient/delete/"+id;
		}
}



