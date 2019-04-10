var app=angular.module("testapp",[]);
app.config(['$locationProvider',function($locationProvider){
	$locationProvider.html5Mode({
		enabled:true,
		requireBase:false
	});
}])

function Ctrl($scope, $http, $location, $window) {
	var url=$location.absUrl();
	var url2=$location.search().code;
	
	$scope.infor=JSON.stringify(url);
	$scope.code=JSON.stringify(url2);
	$scope.code2=JSON.stringify($location.search().code);		
}