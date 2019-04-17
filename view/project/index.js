app.controller('projectCtrl',function($scope, $http, $location, $window) {
    $scope.types = ["横向项目", "纵向项目", "实验室项目"];
	
    $http.get("../testData/project.json"
    ).then(function(results){
    	$scope.project=results.data.data; 
		console.log($scope.project)
    }); 

	$scope.jump =function(data){
		$window.sessionStorage.pid = data;
		window.location ="../project/detail.html";
	}
    $scope.submit = function () {
        console.log($scope.project);
        if ($scope.project.principal.userId == null) {
            $.alert("请选择负责人");
            return;
        }
        var url = "";

        if ($scope.project.project == null) {
            url = "../teacher/project";
        } else {
            url = "../teacher/mission";
        }
        $http({
            url: url,
            method: 'post',
            data: $scope.project,
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (data) {
            $.alert("添加成功", function () {
                $window.sessionStorage.projects = '';
                $window.location.reload(true);
            });

        }).error(function (err) {
            $.alert(err);
        })
    }

})

app.controller('detailCtrl',function($scope, $http, $location, $window) {
    // $.showLoading();
    $http.get("../testData/project.json"
    ).then(function(results){
    	$scope.project=results.data.data; 
    }); 
	
	$scope.pid=$window.sessionStorage.pid;
	var pid=$window.sessionStorage.pid;
//     $http({
//         url: "http://bci.renjiwulian.com/project/details/" + pid,
//         method: 'get'
//     }).success(function (res) {
//         console.log(res);
//         $scope.p = res;
//     }).error(function (err) {
//         $.hideLoading();
//         alert(err);
//     })

    $scope.m = function (status) {
        var tips = "";
        var p = {projectId: $location.search().id};
        if (status == 0) {
            tips = "恢复";
        } else {
            tips = "完成";
        }
        $.actions({
            title: "选择操作",
            onClose: function () {
                console.log("close");
            },
            actions: [
                {
                    text: "修改",
                    className: "color-primary",
                    onClick: function () {
                        $window.sessionStorage.p = JSON.stringify($scope.p);
                        $window.location.href = "../project/update.html";
                    }
                },
                {
                    text: tips,
                    className: "color-primary",
                    onClick: function () {
                        $.confirm("确定要" + tips + "吗？", function () {
                            //点击确认后的回调函数
                            $http({
                                url: "../teacher/project/status",
                                method: 'put',
                                data: p,
                                headers: {token: $window.sessionStorage.weChatId}
                            }).success(function (data) {
                                $.alert("修改成功", function () {
                                    location.href = document.referrer;
                                });
                            }).error(function (err) {
                                $.alert(err);
                            })
                        }, function () {
                            //点击取消后的回调函数
                        });
                    }
                },
                {
                    text: "删除",
                    className: 'color-danger',
                    onClick: function () {
                        $.confirm("确定要删除吗？", function () {
                            //点击确认后的回调函数
                            $http({
                                url: "../teacher/project/" + $location.search().id,
                                method: 'delete',
                                headers: {token: $window.sessionStorage.weChatId}
                            }).success(function (data) {
                                $.alert("删除成功", function () {
                                    $window.sessionStorage.projects = '';
                                    location.href = document.referrer;
                                });
                            }).error(function (err) {
                                $.alert(err);
                            })
                        }, function () {
                            //点击取消后的回调函数
                        });

                    }
                }
            ]
        });
    }
})

function updateCtrl($scope, $http, $location, $window) {
    $scope.p = JSON.parse($window.sessionStorage.p);
    $scope.teacher = JSON.parse($window.sessionStorage.teacher);
    $scope.principal = JSON.parse($window.sessionStorage.principal);
    $scope.projects = JSON.parse($window.sessionStorage.projects);
    $scope.types = ["横向项目", "纵向项目", "实验室项目"];
    $scope.update = function () {
        $http({
            data: $scope.p,
            url: "../teacher/project",
            method: 'put',
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (data) {
            $.alert("修改成功", function () {
                $window.location.href = 'index.html';
            });
        }).error(function (err) {
            $.alert(err);
        })
    }
}

