 function test(){
   $.ajax({
            //提交数据的类型 POST GET
            type:"POST",
            //提交的网址
            url:"http://localhost:8080/weixin/test",
            //提交的数据
            data:{name:"sanmao"},
            //返回数据的格式
           // datatype: "html",//"xml", "html", "script", "json", "jsonp", "text".
            //在请求之前调用的函数
           // beforeSend:function(){$("#msg").html("logining");},
            //成功返回之后调用的函数             
            success:function(data){
               alert(data);            
            }   ,
            //调用执行后调用的函数
            complete: function(XMLHttpRequest, textStatus){
               alert(XMLHttpRequest.responseText);
               alert(textStatus);
                //HideLoading();
            },
            //调用出错执行的函数
            error: function(){
                //请求出错处理
            }         
         });

  }
