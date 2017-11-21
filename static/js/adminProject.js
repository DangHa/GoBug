function CreateProjectTableFromJSON(data) {

  // Header
  var col = ["#", "Project", "Description"];
  var colJSON = ["Id", "TenProject", "MieutaProject"]; // de dong bo voi du lieu JSON

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

  // hide column 1
  for (var i=0; i<table.rows.length; i++){
        table.rows[i].cells[0].style.display = "none";
  }


  // Set onclick cho tung dong
  for (var i = 1; i < table.rows.length; i++) {
    table.rows[i].onclick = function(){

      document.getElementById("id").value = this.cells[0].innerHTML;
      document.getElementById("project").value = this.cells[1].innerHTML;

      for (var j = 1; j < table.rows.length; j++){
        if (j === this.rowIndex) {
          this.style.color = "blue"
        }else{
          table.rows[j].style.color = "black"
        }
      }

    };

  }

}

function PostProject() {
  var project = document.getElementById('new').value
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
  var id = document.getElementById('id').value
  var project = document.getElementById('new').value
  var mieuta = document.getElementById('desc').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminprojectjson/";
  xhr.open("PUT", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Id": id, "Project": project, "Mieuta": mieuta});
  xhr.send(data);
  location.reload();
}

function DeleteProject() {
  var id = document.getElementById('id').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminprojectjson/";
  xhr.open("DELETE", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Id": id});

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
