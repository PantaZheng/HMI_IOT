val = $("#status").attr("value");
var t, action;
if (val == "working") {
	t = "完成";
	action = 0;
} else {
	t = "恢复";
	action = 1;
}


$(document).on("click", "#show-action", function() {
	$.actions({
		actions : [ {
			text : t,
			onClick : function() {
				$.confirm("确定要" + t + "吗?", function() {
					// 点击确认后的回调函数

					$.post("../lab/changeStatusAction", {
						project_id : $("#project_id").attr("value"),
						action : action,
						user_id : $("#user_id").attr("value"),
					}, function(data, status) {
						if (status == "success") {
							$.alert("修改成功", function() {
								 pid = $("#pid").attr("value");
								 uid=$("#uid").attr("value");
								
								// 点击确认后的回调函数
								location.href="../lab/projectDetailAction?user_id="+uid+"&project_id="+pid;
							});

						}
					});
				}, function() {
					// 点击取消后的回调函数
				});
			}
		}, {
			text : "修改",
			onClick : function() {
				$.confirm("确定要修改吗?", function() {
					// 点击确认后的回调函数
					window.location.href = "../lab/projectModifyAction?user_id="
						+ $("#user_id").attr("value") + "&project_id="
						+ $("#project_id").attr("value");
				}, function() {
					// 点击取消后的回调函数
				});
			}
		}, {
			text : "删除",
			onClick : function() {
				$.confirm("确定要删除吗?", function() {
					// 点击确认后的回调函数

					$.post("../lab/changeStatusAction", {
						project_id : $("#project_id").attr("value"),
						action : -1,
						user_id : $("#user_id").attr("value"),
					}, function(data, status) {
						if (status == "success") {
							$.alert("删除成功", function() {
								// 点击确认后的回调函数
								
								uid=$("#uid").attr("value");
								
								
								location.href="../lab/projectEntranceAction?user_id="+uid;

							});

						}
					});
				}, function() {
					// 点击取消后的回调函数
				});
			}
		} ]
	});
});
