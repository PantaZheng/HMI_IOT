function detailCtrl($scope, $http, $location, $window) {
    $http({
        url: "../user/weekly/" + $location.search().weeklyId,
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.weekly = res;
        $.hideLoading();
    }).error(function (err) {
        $.hideLoading();
        $.alert(err);
    })

    $scope.delete = function (id) {
        $.confirm("确定要删除吗？", function () {
            //点击确认后的回调函数
            $http({
                url: "../user/weekly/" + id,
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

function editCtrl($scope, $http, $location, $window) {
    $http({
        url: "../user/mission",
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.wm = new Array();
        for (var i = 0; i < res.length; i++) {
            if (res[i].status == 1) {
                $scope.wm.push(res[i]);
            }
        }
    }).error(function (err) {
        $.alert(err);
    })
    $scope.w = {'content': '', 'others': '', 'mission': {'missionId': 0}};
    $scope.submit = function () {
        if ($scope.w.mission
                .missionId == 0) {
            alert("请选择所属任务");
            return;
        }
        if ($scope.w.content == '') {
            alert("请输入主要内容");
            return;
        }
        $http({
            url: "../user/weekly",
            method: 'post',
            data: $scope.w,
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (res) {
            $.alert("提交成功", function () {
                location.href = document.referrer;
            });
        }).error(function (err) {
            $.alert(err);
        })
    }

}
function studentIndexCtrl($scope, $http, $location, $window) {
    $.showLoading();
    if ($window.sessionStorage.weChatId == null) {
        var weChatId = $location.search().weChatId;
        $window.sessionStorage.weChatId = weChatId;
    }
    var studentId = $location.search().studentId;
    if (studentId == null) {
        getWeekly('../user/weekly/find');
        $scope.showMenu = true;
    } else {
        $scope.showMenu = false;
        getWeekly('../teacher/weekly/find?studentId=' + studentId);
    }

    function getWeekly(url) {

        $http({
            url: url,
            method: 'get',
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (res) {
            console.log(res);
            $scope.weeklies = res;
            $.hideLoading();
        }).error(function (err) {
            $.hideLoading();
            $.alert(err);
        })
    }

}

function teacherIndexCtrl($scope, $http, $location, $window) {
    $.showLoading();
    if ($window.sessionStorage.weChatId == null) {
        var weChatId = $location.search().weChatId;
        $window.sessionStorage.weChatId = weChatId;
    }
    $http({
        url: "../teacher/user/activeStudent",
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.students = res;
        $.hideLoading();
    }).error(function (err) {
        $.hideLoading();
        $.alert(err);
    })


}