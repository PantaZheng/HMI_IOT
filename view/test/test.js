function Ctrl($scope, $http, $location, $window) {
    // $.showLoading();
//     var url=$location.url();
// 	alert(url);	
	$scope.getId=function(){
// 		if ($window.sessionStorage.weChatId == null) {
// 		    var weChatId = $location.search().weChatId;
// 		    $window.sessionStorage.weChatId = weChatId;
// 		}
		var url=$location.absUrl();
		var url2=$location.search().code;
		$scope.infor=JSON.stringify(url);
		alert($scope.infor);	
		// alert("jjj")
		// alert($window.sessionStorage.weChatId);	
	}
	
	
}