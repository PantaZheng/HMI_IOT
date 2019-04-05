function indexCtrl($scope, $http, $location, $window) {
    if ($window.sessionStorage.weChatId == null) {
        var weChatId = $location.search().weChatId;
        $window.sessionStorage.weChatId = weChatId;
    }
    $http({
        url: "../user/notice",
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.notices = res;
    }).error(function (err) {
        $.alert(err);
    })
}
function editCtrl($scope, $http, $location, $window) {
    $scope.n = {'content': ""};
    $scope.submit = function () {
        if ($scope.n.content == "") {
            $.alert("请输入内容");
            return;
        }
        $http({
            url: "../user/notice",
            method: 'post',
            headers: {token: $window.sessionStorage.weChatId},
            data: $scope.n
        }).success(function (res) {
            location.href = document.referrer;
        }).error(function (err) {
            $.alert(err);
        })
    }
}
