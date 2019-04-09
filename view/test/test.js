function Ctrl($scope, $http, $location, $window) {
    // $.showLoading();
    
	$scope.getId=function(){
		if ($window.sessionStorage.weChatId == null) {
		    var weChatId = $location.search().weChatId;
		    $window.sessionStorage.weChatId = weChatId;
		}
		// alert("jjj")
		alert($window.sessionStorage.weChatId);	
	}
	
	
}