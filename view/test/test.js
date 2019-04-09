function Ctrl($scope, $http, $location, $window) {
    $.showLoading();
    if ($window.sessionStorage.weChatId == null) {
        var weChatId = $location.search().weChatId;
        $window.sessionStorage.weChatId = weChatId;
    }
	$scope.getId=function(){
	console.log($window.sessionStorage.weChatId);	
	}
	
	
}