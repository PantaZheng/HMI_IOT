function indexCtrl($scope, $http, $location, $window) {
    if ($window.sessionStorage.weChatId == null) {
        var weChatId = $location.search().weChatId;
        $window.sessionStorage.weChatId = weChatId;
    }
    $scope.c = {"content": '', "userName": ''};
    $http({
        url: "../teacher/conference/find",
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.cs = res;
    }).error(function (err) {
        $.alert(err);
    })
    $scope.submit = function () {
        if ($scope.c.content == '') {
            $.alert("请输入会议内容");
            return;
        }
        if ($scope.c.userName == '') {
            $.alert("请输入参会人员");
            return;
        }
        $http({
            url: "../teacher/conference",
            method: 'post',
            data: $scope.c,
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (res) {
            $.alert("提交成功", function () {
                $window.location.reload();
            });

        }).error(function (err) {
            $.alert(err);
        })
    }
}
function detailCtrl($scope, $http, $location, $window) {
    $http({
        url: "../teacher/conference/" + $location.search().id,
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.c = res;
    }).error(function (err) {
        $.alert(err);
    })
    $scope.m = function () {


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
                        $window.location.href = "update.html#?id=" + $location.search().id;
                    }
                },
                {
                    text: "删除",
                    className: 'color-danger',
                    onClick: function () {
                        $.confirm("确定要删除吗？", function () {
                            //点击确认后的回调函数
                            $http({
                                url: "../teacher/conference/" + $location.search().id,
                                method: 'delete',
                                headers: {token: $window.sessionStorage.weChatId}
                            }).success(function (data) {
                                $.alert("删除成功", function () {
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
    $http({
        url: "../teacher/conference/" + $location.search().id,
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.c = res;
    }).error(function (err) {
        $.alert(err);
    })
    $scope.submit = function () {
        if ($scope.c.content == '') {
            $.alert("请输入会议内容");
            return;
        }
        if ($scope.c.userName == '') {
            $.alert("请输入参会人员");
            return;
        }
        $http({
            url: "../teacher/conference",
            method: 'put',
            data: $scope.c,
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (res) {
            $.alert("提交成功", function () {
                window.location.href = "detail.html#?id=" + $scope.c.id;
            });

        }).error(function (err) {
            $.alert(err);
        })
    }

}