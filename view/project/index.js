app.controller('projectCtrl',function($scope, $http, $location, $window) {
    $scope.types = ["横向项目", "纵向项目", "实验室项目"];
    $http.get("../testData/project.json"
    ).then(function(results){
    	$scope.project=results.data.data; 
		console.log($scope.projcet)
    }); 

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
    $scope.getProjects = function () {
        var url = '';
        if ($scope.user.level == 1) {
            url = '../admin/project';
        } else {
            url = '../teacher/project';
        }
        if ($window.sessionStorage.projects == null || $window.sessionStorage.projects == '') {
            $http({
                url: url,
                method: 'get',
                headers: {token: $window.sessionStorage.weChatId}
            }).success(function (data) {
                $scope.projects = data;
                $window.sessionStorage.projects = JSON.stringify($scope.projects);
            }).error(function (err) {
                $.alert(err);
            })
        } else {
            $scope.projects = JSON.parse($window.sessionStorage.projects);
        }

    }

    $scope.getPrincipal = function () {
        if ($window.sessionStorage.principal == null) {
            $http({
                url: "../teacher/user/activePrincipal",
                method: 'get',
                headers: {token: $window.sessionStorage.weChatId}
            }).success(function (data) {
                $scope.principal = data;
                $window.sessionStorage.principal = JSON.stringify($scope.principal);
            }).error(function (err) {
                $.alert(err);
            })
        } else {
            $scope.principal = JSON.parse($window.sessionStorage.principal);
        }

    }
    $scope.getTeacher = function () {
        if ($window.sessionStorage.teacher == null) {
            $http({
                url: "../teacher/user/teacher",
                method: 'get',
                headers: {token: $window.sessionStorage.weChatId}
            }).success(function (data) {
                $scope.teacher = data;
                $window.sessionStorage.teacher = JSON.stringify($scope.teacher);
            }).error(function (err) {
                $.alert(err);
            })
        } else {
            $scope.teacher = JSON.parse($window.sessionStorage.teacher);
        }

    }
    $scope.getProject = function (role, level) {
        var url = "";
        if (level == 1) {
            url = "../admin/project"
        } else if (role == "teacher") {
            url = "../teacher/project"
        } else {
            url = "../user/project"
        }
        $http({
            url: url,
            method: 'get',
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (data) {
            $scope.hp0 = new Array();
            $scope.zp0 = new Array();
            $scope.sp0 = new Array();
            $scope.hp1 = new Array();
            $scope.zp1 = new Array();
            $scope.sp1 = new Array();
            for (var i = 0; i < data.length; i++) {
                if (data[i].type == '横向项目') {
                    if (data[i].status == 0) {
                        $scope.hp0.push(data[i]);
                    } else {
                        $scope.hp1.push(data[i]);
                    }

                } else if (data[i].type == '纵向项目') {
                    if (data[i].status == 0) {
                        $scope.zp0.push(data[i]);
                    } else {
                        $scope.zp1.push(data[i]);
                    }
                } else {
                    if (data[i].status == 0) {
                        $scope.sp0.push(data[i]);
                    } else {
                        $scope.sp1.push(data[i]);
                    }
                }
            }
        }).error(function (err) {
            $.alert(err);
        })
    }
})

function detailCtrl($scope, $http, $location, $window) {
    $.showLoading();
    var id = $location.search().id;
    $http({
        url: "../teacher/project/detail/" + id,
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.p = res;
        $scope.ps = new Array();
        for (var i = 0; i < res.missions.length; i++) {
            var name = res.missions[i].principal.name;
            var isContain = false;
            for (var j = 0; j < $scope.ps.length; j++) {
                if (name == $scope.ps[j]) {
                    isContain = true;
                    break;
                }
            }
            if (isContain == false) {
                $scope.ps.push(name);
            }
        }
        console.log($scope.ps);
        $.hideLoading();
    }).error(function (err) {
        $.hideLoading();
        $.alert(err);
    })

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
}

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

