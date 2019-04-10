function Ctrl($scope, $http, $location, $window) {
    // $.showLoading();
//     var url=$location.url();
// 	alert(url);	
	// $scope.getId=function(){
// 		if ($window.sessionStorage.weChatId == null) {
// 		    var weChatId = $location.search().weChatId;
// 		    $window.sessionStorage.weChatId = weChatId;
// 		}
		var url=$location.absUrl();
		var url2=$location.search().code;
		var c=$location.search()['code'];
		
		$scope.infor=JSON.stringify(url);
		$scope.code=JSON.stringify(url2);
		$scope.code1=JSON.stringify(c);		
		// $scope.w=JSON.stringify("code=dsdsd&appid=fdfdfd");
		alert(c);	
		// alert("jjj")
		// alert($window.sessionStorage.weChatId);	
	// }
	
	
}