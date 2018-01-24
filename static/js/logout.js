function LogOut() {

  var data = JSON.stringify({});

  $.ajax({
      type:"POST",
      url: 'http://localhost:8080/loginAdmin/',
      data:data,
      success: function (response){
        window.location="http://localhost:8080/";
      }
  });
}
