$(function(){
	$('.aside h4').click(function(){
//		$(this).toggleClass('active');
		$(this).siblings('ul').slideToggle();
	})
})

function confirmDelete(url) {
	if (confirm("确定删除吗")) {
		location.href = url
	}
}

window.onload = function () {
	const iframe = document.getElementById("rightMain")
	if (iframe) {
		iframe.height = window.innerHeight - 80 + 'px'
	}
}
