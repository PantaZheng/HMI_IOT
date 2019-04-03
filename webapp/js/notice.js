var xmlHttp;
if (window.XMLHttpRequest) {// code for IE7+, Firefox, Chrome, Opera, Safari
	xmlHttp = new XMLHttpRequest();
} else {// code for IE6, IE5
	xmlHttp = new ActiveXObject("Microsoft.XMLHTTP");
}

function get() {
	xmlHttp.open("GET", "/weixin/notice?t=" + Math.random(), true);
	xmlHttp
			.setRequestHeader('Content-Type',
					'application/x-www-form-urlencoed');

	xmlHttp.onreadystatechange = function() {
		if (xmlHttp.readyState == 4) {
			var obj = xmlHttp.responseText;
			var obj = eval('(' + obj + ')');
			// alert(obj.data[0].title);

			var target = document.getElementById('content');
			var h3 = target.getElementsByTagName('h3')[0];

			var en = target.getElementsByTagName('em')[0];

			var p = target.getElementsByTagName('p')[0];

			var output = Mustache.render(
					'{{#data}}h3.innerHTML={{title}}{{time}}{{content}}', obj);
		}

	}
	xmlHttp.send();

}

function submit() {
	var id = document.getElementById("id");
	var user_id = document.getElementById("user_id");
	var pri = document.getElementById("principal");
	var title = document.getElementById("title");
	var content = document.getElementById("content");
	if (id == null) {
		xmlHttp.open("post", "/weixin/notice", true);
		xmlHttp.onreadystatechange = function() {
			if (xmlHttp.readyState == 4) {
				var obj = xmlHttp.responseText;
				var obj = eval('(' + obj + ')');
				alert(obj);
			}
		}
		xmlHttp.send("user_id=" + user_id + "title=" + title + "content="
				+ content);
	} else {
		xmlHttp.open("post", "/weixin/notice", true);
		xmlHttp.onreadystatechange = function() {
			if (xmlHttp.readyState == 4) {
				var obj = xmlHttp.responseText;
				var obj = eval('(' + obj + ')');
				alert(obj);
			}
		}
		xmlHttp.send("principal=" + pri + "title=" + title + "content="
				+ content);
	}

}
