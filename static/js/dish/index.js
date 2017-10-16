function dish_add() {
		window.location.href="/dish/add";
}

function dish_edit(id) {
		window.location.href="/dish/edit/"+id;
}

function dish_delete(id){
		var r = confirm("Delete Dish ?");
		if (r == true) {
				window.location.href="/dish/delete/"+id;
		}
}



