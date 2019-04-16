app.controller('indexCtrl',function($scope, $http, $location, $window) {
    // $.showLoading();
	$window.sessionStorage.code=JSON.stringify($location.search().code);
})

app.controller('studentCtrl',function($scope, $http, $location, $window) {
    $.showLoading();
	$window.sessionStorage.role="student";
    $http.get("http://bci.renjiwulian.com/anon/list/teacher"
    ).then(function(results){
    	$scope.teacherList=results.data; 
		console.log($scope.teacherList)
    }); 

    $scope.submit = function () {
        $scope.stuInfor=JSON.stringify({
			// role:"student",
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
            data: $scope.stuInfor
        }).success(function (results) {			
			alert("提交成功！");
			$scope.openid=results.data.openid;
			alert($scope.openid);
			$window.sessionStorage.openid=results.data.openid;
            $window.location.reload(true);          
        }).error(function (err) {
            alert(err);
        })
    }
})

app.controller('teacherCtrl',function($scope, $http, $location, $window) {
    $.showLoading();
	$window.sessionStorage.role="teacher";
    $scope.submit = function () {
        $scope.teaInfor=JSON.stringify({
			// role:"teacher",
        	code:$location.search().code,
        	name:$scope.name,
        	sex:$scope.sex,
        	telephone:$scope.tel,
        	school:$scope.school
        })
        $http({
            url:"http://bci.renjiwulian.com/teacher/enroll",
            method: 'post',
            data: $scope.teaInfor
        }).success(function (results) {			
			alert("提交成功！");
			$scope.openid=results.data.openid;
			alert($scope.openid);
			$window.sessionStorage.openid=results.data.openid;
            $window.location.reload(true);          
        }).error(function (err) {
            alert(err);
        })
    }
})