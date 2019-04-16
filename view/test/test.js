var app=angular.module("app",[]);
app.config(['$locationProvider',function($locationProvider){
	$locationProvider.html5Mode({
		enabled:true,
		requireBase:false
	});
}])

function Ctrl($scope, $http, $location, $window) {
	// var code=$location.search().code;
	$window.sessionStorage.code = $location.search().code;

    $http.get("http://bci.renjiwulian.com/teacher/list"
    ).then(function(results){
    	$scope.teacherList=results; 
		console.log($scope.teacherList);
    }); 

    $scope.submit = function () {
		$scope.stuInfor=JSON.stringify({
			openid:"",
			code:$location.search().code,
			name:$scope.name,
			sex:$scope.sex,
			telephone:$scope.tel,
			school:$scope.school,
			supervisor:$scope.teacher
		})
        $http({
            url:"http://bci.renjiwulian.com/student/enroll",
            method: 'post',
            data: $scope.stuInfor,
            // headers: {token: $window.sessionStorage.weChatId}
        }).success(function (data) {
            alert("提交成功！");
            $window.location.reload(true); 
        }).error(function (err) {
            alert("552");
        })
    }
			
}
