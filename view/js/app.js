var app=angular.module("app",[]);
app.config(['$locationProvider',function($locationProvider){
	$locationProvider.html5Mode({
		enabled:true,
		requireBase:false
	});
}])
