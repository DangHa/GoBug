function CreateProjectTableFromJSON(data) {

  // Header
  var col = ["#", "Email", "Position", "Number of project"];
  var colJSON = ["Id", "Email", "Vaitro", "Number"]; // de dong bo voi du lieu JSON

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
      var tabCell1 = tr.insertCell(-1);
      tabCell1.innerHTML = '<input type="submit" value="Delete" onclick="DeleteMember()"/>'
  }

  // hide column 1
  for (var i=0; i<table.rows.length; i++){
        table.rows[i].cells[0].style.display = "none";
  }

  // Set onclick cho tung dong
  for (var i = 1; i < table.rows.length; i++) {
    table.rows[i].onmouseover = function(){

      document.getElementById("id").value = this.cells[0].innerHTML;
      document.getElementById("email").value = this.cells[1].innerHTML;
      document.getElementById('position').value = this.cells[2].innerHTML;

      this.style.color = "blue";
    };
    table.rows[i].onmouseout = function(){
      document.getElementById("id").value = "";
      document.getElementById("email").value = "";
      document.getElementById('position').value = "";

      this.style.color = "black";
    };

  }
}

function Positi(obj) {
  document.getElementById('posit').value = obj.value
}

function PostMember() {
  var email = document.getElementById('new').value
  var vaitro = document.getElementById('posit').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminmemberjson/";
  xhr.open("POST", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Email": email, "Vaitro": vaitro});

  xhr.send(data);
  location.reload();
}

function DeleteMember() {
  var email = document.getElementById('email').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminmemberjson/";
  xhr.open("DELETE", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Email": email});

  xhr.send(data);
  location.reload();
}


$( document ).ready(function() {
  //Get JSON
  var url = 'http://localhost:8080/adminmemberjson/'
  $.getJSON(url, function(data){
    console.log("It Worked!");
    CreateProjectTableFromJSON(data)
    console.log(data);
  })

});
