
function CreateMemberTableFromJSON(data) {

  // Header
  var col = ["Domain"];
  var colJSON = ["Domain"]; // de dong bo voi du lieu JSON

  // Goi den bang can tim
  var table = document.getElementById("memberTable");

  // lam dau bang
  var tr = table.insertRow(-1);                   // TABLE ROW.

  for (var i = 0; i < col.length; i++) {
      var th = document.createElement("th");      // TABLE HEADER.
      th.innerHTML = col[i];
      tr.appendChild(th);
  }

  // add du lieu vao cac dong
  for (var i = 0; i < data["Domain"].length; i++) {

      tr = table.insertRow(-1);

      for (var j = 0; j < col.length; j++) {
          var tabCell = tr.insertCell(-1);
          tabCell.innerHTML = data["Domain"][i];
      }
  }

}

$( document ).ready(function() {
  //Get JSON
  var url = 'http://localhost:8080/mastergetjson/'
  $.getJSON(url, function(data){
    console.log("It Worked!");
    CreateMemberTableFromJSON(data)
    console.log(data);
  })

});
