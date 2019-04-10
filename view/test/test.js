var app=angular.module("testapp",[]);
app.config(['$locationProvider',function($locationProvider){
	$locationProvider.html5Mode({
		enabled:true,
		requireBase:false
	});
}])

function Ctrl($scope, $http, $location, $window) {
	var code=$location.search().code;
	$scope.code2=JSON.stringify($location.search().code);		
    $http({
        url: " https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx2203c68c9311ea43&secret=40c40547e174ed99d1281b2890f7eeb3&code="+$scope.code+"&grant_type=authorization_code",
        method: 'get',
        // headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        // console.log(res);
        $scope.infor = JSON.stringify(res);
		alert($scope.infor);
        // $window.sessionStorage.user = JSON.stringify(res);
        // $scope.getProject(res.role, res.level);
    }).error(function (err) {
        $.hideLoading();
        $.alert(err);
    })
    $http({
        url: " https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx2203c68c9311ea43&secret=40c40547e174ed99d1281b2890f7eeb3&code="+code+"&grant_type=authorization_code",
        method: 'get',
        // headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        // console.log(res);
        $scope.code = JSON.stringify(res);
		// alert($scope.infor);
    }).error(function (err) {
        $.hideLoading();
        $.alert(err);
    })		
}