
function CreateUserProjectTableFromJSON(data) {

  // Header
  var col = ["#","Project", "Description", "Number of Bug"];
  var colJSON = ["Id", "Project", "Description", "Number"]; // de dong bo voi du lieu JSON

  // Goi den bang can tim
  var table = document.getElementById("projectUserTable");

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
      if (true) { // Kiem tra xem sessionID cua user hay cua admin
        var tabCell = tr.insertCell(-1);
        tabCell.innerHTML = '<input type="submit" value="Add bug" onclick="CreateBug()"/>';
      }
      var tabCell1 = tr.insertCell(-1);
      tabCell1.innerHTML = '<input type="submit" value="Show all bugs" onclick="CreateBugTable()"/>';
  }

  // hide column 1
  for (var i=0; i<table.rows.length; i++){
        table.rows[i].cells[0].style.display = "none";
  }

  // Set onclick cho tung dong
  for (var i = 1; i < table.rows.length; i++) {
    table.rows[i].onmouseover = function(){

      document.getElementById("id").value = this.cells[0].innerHTML;

      this.style.color = "blue"
    };
    table.rows[i].onmouseout = function(){

      document.getElementById("id").value = "";

      this.style.color = "black"
    };
  }

}

//Tao them bang khi an vao show all bugs
function CreateBugProjectTableFromJSON(data) {

  // Header
  var col = ["#","Bug", "Description", "Solution"];
  var colJSON = ["Id", "BugName", "BugDescription", "SolutionDescription"]; // de dong bo voi du lieu JSON

  // Goi den bang can tim
  var table = document.getElementById("bugUserTable");

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
          tabCell.contentEditable =true;
      }
      var tabCell = tr.insertCell(-1);
      tabCell.innerHTML = '<input type="submit" value="Update" onclick="UpdateBug()"/>';

      var tabCell1 = tr.insertCell(-1);
      tabCell1.innerHTML = '<input type="submit" value="Delete" onclick="DeleteBug()"/>';
  }

  // hide column 1
  for (var i=0; i<table.rows.length; i++){
        table.rows[i].cells[0].style.display = "none";
  }

  // Set onclick cho tung dong
  for (var i = 1; i < table.rows.length; i++) {
    table.rows[i].onmouseover = function(){

      document.getElementById("idbug").value = this.cells[0].innerHTML;
      document.getElementById("tenbug").value = this.cells[1].innerHTML;
      document.getElementById("mtbug").value = this.cells[2].innerHTML;
      document.getElementById("mtsolution").value = this.cells[3].innerHTML;

      this.style.color = "blue";
    };
    table.rows[i].onmouseout = function(){

      document.getElementById("idbug").value = "";
      document.getElementById("tenbug").value = "";
      document.getElementById("mtbug").value = "";
      document.getElementById("mtsolution").value = "";

      this.style.color = "black";
    };
  }

}

function CreateBug() {
  if (document.getElementById('newBug').style.display === "none"){
    document.getElementById('newBug').style.display = "block";
    document.getElementById('id-add').value = document.getElementById('id').value;
  }else{
    document.getElementById('newBug').style.display = "none";
    document.getElementById('id-add').value = "";
  }
}

function CreateBugTable() {

  document.getElementById('bugUserTable').innerHTML = "";
  document.getElementById('id-add').value = document.getElementById('id').value;

  if (document.getElementById('id-add').value === document.getElementById('id').value){
    if (document.getElementById('bugUserTable').style.display === "none"){
      document.getElementById('bugUserTable').style.display = "block";
    }else{
      document.getElementById('bugUserTable').style.display = "none";
      document.getElementById('id-add').value = "";
    }
  }
  var data = JSON.stringify({"IdProject": parseInt(document.getElementById('id').value)});
  $.ajax({
      type:"POST",
      url: 'http://localhost:8080/userbugjson/',
      data:data,
      success: function (response){
        CreateBugProjectTableFromJSON(response);
      }
  });
}

function PostBug() {

  var idproject = parseInt(document.getElementById('id-add').value)
  var name = document.getElementById('new').value
  var des = document.getElementById('desc').value
  var solu = document.getElementById('solution').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/userprojectjson/";
  xhr.open("POST", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"BugName": name, "BugDescription": des, "SolutionDescription": solu, "IdProject": idproject});
  xhr.send(data);
  location.reload();
}

function UpdateBug() {

  var idbug = parseInt(document.getElementById('idbug').value)
  var name = document.getElementById('tenbug').value
  var des = document.getElementById('mtbug').value
  var solu = document.getElementById('mtsolution').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/userbugjson/";
  xhr.open("PUT", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Id": idbug, "BugName": name, "BugDescription": des, "SolutionDescription": solu});
  xhr.send(data);
  location.reload();
}

function DeleteBug() {
  var id = parseInt(document.getElementById('idbug').value)

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/userbugjson/";
  xhr.open("DELETE", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Id": id});

  xhr.send(data);
  location.reload();
}


$( document ).ready(function() {
  //Get JSON
  var url = 'http://localhost:8080/userprojectjson/'
  $.getJSON(url, function(data){
    console.log("It Worked!");
    CreateUserProjectTableFromJSON(data)
    console.log(data);
  })

});
