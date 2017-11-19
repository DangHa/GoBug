function CreateProjectTableFromJSON(data) {

  // Header
  var col = ["Project","Developer", "Tester", "Bug"];
  var colJSON = ["Project", "Developer", "Tester", "Bug"]; // de dong bo voi du lieu JSON

  // Goi den bang can tim
  var table = document.getElementById("projectTable");

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

function PostProject() {
  var project = document.getElementById('domain').value
  var mieuta = document.getElementById('desc').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminprojectjson/";
  xhr.open("POST", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Project": project, "Mieuta": mieuta});

  xhr.send(data);
  location.reload();
}

function PutProject() {
  var project = document.getElementById('domain').value
  var mieuta = document.getElementById('desc').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminprojectjson/";
  xhr.open("PUT", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Project": project, "Mieuta": mieuta});
  xhr.send(data);
  location.reload();
}

function DeleteProject() {
  var project = document.getElementById('domain').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminprojectjson/";
  xhr.open("DELETE", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Project": project});

  xhr.send(data);
  location.reload();
}


$( document ).ready(function() {
  //Get JSON
  var url = 'http://localhost:8080/adminprojectjson/'
  $.getJSON(url, function(data){
    console.log("It Worked!");
    CreateProjectTableFromJSON(data)
    console.log(data);
  })

});
