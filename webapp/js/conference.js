
$(document).on("click", "#show-action", function() {
	$.actions({
		actions: [{
			text: "修改",
			onClick: function() {
				$.confirm("确定要修改吗?", function() {
					// 点击确认后的回调函数
					window.location.href = "../lab/updateViewAction?user_id=" + $("#user_id").attr("value") + "&id=" + $("#conferenceId").attr("value");
				}, function() {
					// 点击取消后的回调函数
				});
			}
		}, {
			text: "删除",
			onClick: function() {
				$.confirm("确定要删除吗?", function() {
					// 点击确认后的回调函数

					$.post("../lab/deleteMeetingAction", {
						id: $("#conferenceId").attr("value"),
						
						user_id: $("#user_id").attr("value"),
					}, function(data, status) {
						if (status == "success") {
							$.alert("删除成功", function() {
								// 点击确认后的回调函数

								uid = $("#user_id").attr("value");

								location.href = "../lab/meetingViewAction?user_id=" + uid;

							});

						}
					});
				}, function() {
					// 点击取消后的回调函数
				});
			}
		}]
	});
});

