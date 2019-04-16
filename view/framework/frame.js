app.controller('ctrl',function($scope, $http,$location, $window) {
    $http.get("../testData/member.json"
    ).then(function(results){
    	$scope.list=results.data.data; 
		console.log($scope.list)
    }); 

})
