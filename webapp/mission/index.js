function updateCtrl($scope, $http, $location, $window) {
    $scope.p = JSON.parse($window.sessionStorage.m);
    $scope.teacher = JSON.parse($window.sessionStorage.teacher);
    $scope.principal = JSON.parse($window.sessionStorage.principal);
    $scope.projects = JSON.parse($window.sessionStorage.projects);
    $scope.types = ["横向项目", "纵向项目", "实验室项目"];
    $scope.update = function () {
        $http({
            data: $scope.p,
            url: "../teacher/mission",
            method: 'put',
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (data) {
            $.alert("修改成功", function () {
                $window.location.href = 'detail.html#?id=' + $scope.p.missionId + "&pid=" + $location.search().pid;
            });
        }).error(function (err) {
            $.alert(err);
        })
    }
}
function indexCtrl($scope, $http, $location, $window) {
    $.showLoading();
    if ($window.sessionStorage.weChatId == null) {
        var weChatId = $location.search().weChatId;
        $window.sessionStorage.weChatId = weChatId;
    }
    $http({
        url: "../user/user",
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.user = res;
        $window.sessionStorage.user = JSON.stringify(res);
        $scope.getMission();
        $.hideLoading();
    }).error(function (err) {
        $.hideLoading();
        $.alert(err);
    })
    $scope.getMission = function () {
        $http({
            url: "../user/mission",
            method: 'get',
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (res) {
            console.log(res);
            $scope.w = new Array();
            $scope.c = new Array();
            for (var i = 0; i < res.length; i++) {
                if (res[i].status == 0) {
                    $scope.c.push(res[i]);
                } else {
                    $scope.w.push(res[i]);
                }
            }
            $.hideLoading();
        }).error(function (err) {
            $.hideLoading();
            $.alert(err);
        })
    }
}
function detailCtrl($scope, $http, $location, $window) {
    $.showLoading();
    $scope.user = JSON.parse($window.sessionStorage.user);
    console.log($scope.user);
    var id = $location.search().id;
    $http({
        url: "../user/mission/detail/" + id,
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.mission = res;
        $.hideLoading();
    }).error(function (err) {
        $.hideLoading();
        $.alert(err);
    })
    $scope.back = function () {
        if ($location.search().pid != null) {
            $window.location.href = "../project/detail.html#?id=" + $location.search().pid;
        } else {
            $window.location.href = "index.html";
        }
    }
    $scope.m = function (status) {
        var tips = "";
        var p = {missionId: $location.search().id};
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
                        $window.sessionStorage.m = JSON.stringify($scope.mission);
                        var projectId = $location.search().pid;
                        $window.location.href = "../mission/update.html#?pid=" + projectId;
                    }
                },
                {
                    text: tips,
                    className: "color-primary",
                    onClick: function () {
                        $.confirm("确定要" + tips + "吗？", function () {
                            //点击确认后的回调函数
                            $http({
                                url: "../teacher/mission/status",
                                method: 'put',
                                data: p,
                                headers: {token: $window.sessionStorage.weChatId}
                            }).success(function (data) {
                                $.alert("修改成功", function () {
                                    var projectId = $location.search().pid;
                                    window.location.href = "../project/detail.html#?id=" +
                                        projectId;
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
                                url: "../teacher/mission/" + $location.search().id,
                                method: 'delete',
                                headers: {token: $window.sessionStorage.weChatId}
                            }).success(function (data) {
                                $.alert("删除成功", function () {
                                    var projectId = $location.search().pid;
                                    window.location.href = "../project/detail.html#?id=" +
                                        projectId;
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
function weeklyDetailCtrl($scope, $http, $location, $window) {

    $http({
        url: "../user/user",
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        var url = '';
        if (res.role == 'teacher') {
            url = '../teacher/weekly/';
        } else {
            url = '../user/weekly/';
        }

        $http({
            url: url + $location.search().weeklyId,
            method: 'get',
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (res) {
            console.log(res);
            $scope.weekly = res;
        }).error(function (err) {
            $.alert(err);
        })
    }).error(function (err) {
        $.alert(err);
    })
}

