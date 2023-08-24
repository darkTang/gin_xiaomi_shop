$(function(){
	$('.aside h4').click(function(){
//		$(this).toggleClass('active');
		$(this).siblings('ul').slideToggle();
	})
})

function confirmDelete(id) {
	if (confirm("确定删除吗")) {
		location.href = `/admin/role/delete?id=${id}`
	}
}

window.onload = function () {
	document.getElementById("rightMain").height = window.innerHeight - 80 + 'px'
}
