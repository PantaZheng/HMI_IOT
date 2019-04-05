function indexCtrl($scope, $http, $location, $window) {
            if ($window.sessionStorage.weChatId == null) {
                var weChatId = $location.search().weChatId;
                $window.sessionStorage.weChatId = weChatId;
            }
            $http({
                url: "../teacher/user/activeStudent",
                method: 'get',
                headers: {token: $window.sessionStorage.weChatId}
                }).success(function (res) {
                console.log(res);
                $scope.students = res;
            }).error(function (err) {
                alert(err);
        })

    $scope.assess=function(userid){{
        if($window.sessionStorage.year&&$window.sessionStorage.month)
        {
            $window.location.href = 'ass.html#?studentId=' + userid+ "&year=" + $window.sessionStorage.year + "&month=" + $window.sessionStorage.month;
        }
        else
        {
            alert("未选择年月");
        }
    }}

    $scope.next=function(){
        //     $http({
        //     url: "../teacher/assessment/excel?month=" + $window.sessionStorage.month + "&year=" + $window.sessionStorage.year,
        //     method: 'get',
        //     headers: {token: $window.sessionStorage.weChatId},
        //     responseType: 'arraybuffer'
        // }).success(function (data) {
        //     console.log(data);
        //     var type = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet";
        //     var blob = new Blob([data], {type: type});
        //     saveAs(blob, "考核.xlsx");
        // }).error(function (err) {
        //     alert(err);
        // })
        $http({
            url:"../teacher/assessment/excelFile?month=" + $window.sessionStorage.month + "&year=" + $window.sessionStorage.year,
            method:'get',
            headers:{token:$window.sessionStorage.weChatId}
        }).success(function(res){
            $window.location.href = '../excel/'+res;
        }).error(function (err) {
            alert(err);
        })
    }
}
function monthCtrl($scope, $http, $location, $window) {
    $scope.year = '';
    $scope.month = '';
    var months = new Array();
    for (var i = 1; i <= 12; i++) {
        months.push(i);
    }
    var today = new Date();
    var year = today.getFullYear();
    $("#year").picker({
        toolbarTemplate: '<header class="bar bar-nav">\
  <button class="button button-link pull-right close-picker">确定</button>\
  <h1 class="title">年份</h1>\
  </header>',
        cols: [
            {
                textAlign: 'center',
                values: [year, year - 1]
            }
        ]
    });
    $("#month").picker({
        toolbarTemplate: '<header class="bar bar-nav">\
  <button class="button button-link pull-right close-picker">确定</button>\
  <h1 class="title">月份</h1>\
  </header>',
        cols: [
            {
                textAlign: 'center',
                values: months
            }
        ]
    });
$scope.tm=function() {
        if ($scope.year == '') {
            alert("请选择年份");
            return;
        }
        if ($scope.month == '') {
            alert("请选择月份");
            return;
        }
        else{
            $window.sessionStorage.year = $scope.year;
            $window.sessionStorage.month = $scope.month;
            console.log($window.sessionStorage.month);
            window.location.href ='index.html';
        }

        //$window.location.href = 'ass.html#?studentId=' + $location.search().studentId + "&year=" + $scope.year + "&month=" + $scope.month;
    }
}

function assCtrl($scope, $http, $location, $window) {
    $scope.ranks = ["A", "B", "C", "D", "E"];
    var studentId = $location.search().studentId;
    var month = $window.sessionStorage.month;
    var year = $window.sessionStorage.year;
    $scope.month = month;
    $scope.year = year;
    $scope.studentId = studentId;
    $http({
        url: "../teacher/weekly/smy2?studentId=" + studentId + "&month=" + $window.sessionStorage.month + "&year=" +$window.sessionStorage.year,
        method: 'get',
        headers: {token: $window.sessionStorage.weChatId}
    }).success(function (res) {
        console.log(res);
        $scope.weeklies = res;
    }).error(function (err) {
        alert(err);
    })
    loadData();

    $scope.save = function () {
        $scope.a.commit = 0;
        console.log($scope.a);
        submit();
    }
    $scope.commit = function () {
        $scope.a.commit = 1;
        console.log($scope.a);
        submit();
    }
    function submit() {
        $scope.a.year = $scope.year;
        $scope.a.month = $scope.month;
        $scope.a.student = {"userId": $scope.studentId};
        $http({
            url: "../teacher/assessment",
            method: 'post',
            data: $scope.a,
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (res) {
            $.alert('操作成功', function () {
                loadData();
                $window.location.herf ='index.html';
            });

        }).error(function (err) {
            alert(err);
        })
    }

    function loadData() {
        $http({
            url: "../teacher/assessment/smy?studentId=" + studentId + "&month=" + $window.sessionStorage.month+ "&year=" + $window.sessionStorage.year,
            method: 'get',
            headers: {token: $window.sessionStorage.weChatId}
        }).success(function (res) {
            console.log(res);
            $scope.a = res;
        }).error(function (err) {
            alert(err);
        })
    }

}

function exportCtrl($scope, $http, $location, $window) {
  //   var months = new Array();
  //   for (var i = 1; i <= 12; i++) {
  //       months.push(i);
  //   }
  //   var today = new Date();
  //   var year = today.getFullYear();
  //   $("#year").picker({
  //       toolbarTemplate: '<header class="bar bar-nav">\
  // <button class="button button-link pull-right close-picker">确定</button>\
  // <h1 class="title">年份</h1>\
  // </header>',
  //       cols: [
  //           {
  //               textAlign: 'center',
  //               values: [year, year - 1]
  //           }
  //       ]
  //   });
  //   $("#month").picker({
  //       toolbarTemplate: '<header class="bar bar-nav">\
  // <button class="button button-link pull-right close-picker">确定</button>\
  // <h1 class="title">月份</h1>\
  // </header>',
  //       cols: [
  //           {
  //               textAlign: 'center',
  //               values: months
  //           }
  //       ]
  //   });
  //   $scope.year = '';
  //   $scope.month = '';
  //   $scope.next = function () {
  //       if ($scope.year == '') {
  //           $.alert('请选择年份');
  //
  //       }
  //       if ($scope.month == '') {
  //           $.alert('请选择月份');
  //           return;
  //       }
        $http({
            url: "../teacher/assessment/excel?month=" + $window.sessionStorage.month + "&year=" + $window.sessionStorage.year,
            method: 'get',
            headers: {token: $window.sessionStorage.weChatId},
            responseType: 'arraybuffer'
        }).success(function (data) {
            console.log(data);
            var type = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet";
            var blob = new Blob([data], {type: type});
            saveAs(blob, "考核.xlsx");
        }).error(function (err) {
            alert(err);
        })
    //}
}