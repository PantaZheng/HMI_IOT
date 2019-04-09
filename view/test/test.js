function Ctrl($scope, $http, $location, $window) {
    // $.showLoading();
    
	$scope.getId=function(){
		var url=$location.absUrl();
		alert(url);
	}
}