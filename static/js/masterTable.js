
function CreateMemberTableFromJSON(data) {

  // Header
  var col = ["Domain","Email"];
  var colJSON = ["Domain", "Email"]; // de dong bo voi du lieu JSON

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
  for (var i = 0; i < data.length; i++) {

      tr = table.insertRow(-1);

      for (var j = 0; j < col.length; j++) {
          var tabCell = tr.insertCell(-1);
          tabCell.innerHTML = data[i][colJSON[j]];
      }
  }

}

function Put() {
  var email = document.getElementById('email').value
  var domain = document.getElementById('domain').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/master/";
  xhr.open("PUT", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Email": email, "Domain": domain});
  xhr.send(data);
  location.reload();
}

function Delete() {
  var email = document.getElementById('email').value
  var domain = document.getElementById('domain').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/master/";
  xhr.open("DELETE", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Email": email, "Domain": domain});

  xhr.send(data);
  location.reload();
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
