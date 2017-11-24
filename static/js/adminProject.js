function CreateProjectTableFromJSON(data) {

  // Header
  var col = ["#", "Project", "Description"];
  var colJSON = ["Id", "ProjectName", "ProjectDescription"]; // de dong bo voi du lieu JSON

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
          tabCell.contentEditable =true;
      }

      var tabCell = tr.insertCell(-1);
      tabCell.innerHTML = '<input type="submit" value="Update" onclick="PutProject()"/>'

      var tabCell1 = tr.insertCell(-1);
      tabCell1.innerHTML = '<input type="submit" value="Delete" onclick="DeleteProject()"/>'

      var tabCell2 = tr.insertCell(-1);
      tabCell2.innerHTML = '<input type="submit" value="Add member" onclick="AddMember()"/>'

      var tabCell3 = tr.insertCell(-1);
      tabCell3.innerHTML = '<input type="submit" value="Show all members" onclick="CreateMemberTable()"/>'
  }

  // hide column 1
  for (var i=0; i<table.rows.length; i++){
        table.rows[i].cells[0].style.display = "none";
  }


  // Set onclick cho tung dong
  for (var i = 1; i < table.rows.length; i++) {
    table.rows[i].onmouseover = function(){

      document.getElementById("id").value = this.cells[0].innerHTML;
      document.getElementById("project").value = this.cells[1].innerHTML;
      document.getElementById('mieuta').value = this.cells[2].innerHTML;

      this.style.color = "blue";
    };
    table.rows[i].onmouseout = function(){
      document.getElementById("id").value = "";
      document.getElementById("project").value = "";
      document.getElementById('mieuta').value = "";
      this.style.color = "black";
    };

  }

}

//Tao them bang khi an vao show all bugs
function CreateMemberProjectTableFromJSON(data) {

  // Header
  var col = ["#","Member", "Position"];
  var colJSON = ["Id", "Member", "Position"]; // de dong bo voi du lieu JSON

  // Goi den bang can tim
  var table = document.getElementById("membersTable");

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
      var tabCell = tr.insertCell(-1);
      tabCell.innerHTML = '<input type="submit" value="Delete" onclick="DeleteMember()"/>';
  }

  // hide column 1
  for (var i=0; i<table.rows.length; i++){
        table.rows[i].cells[0].style.display = "none";
  }

  // Set onclick cho tung dong
  for (var i = 1; i < table.rows.length; i++) {
    table.rows[i].onmouseover = function(){

      document.getElementById("iduser").value = this.cells[0].innerHTML;

      this.style.color = "blue";
    };
    table.rows[i].onmouseout = function(){

      document.getElementById("iduser").value = "";

      this.style.color = "black";
    };
  }

}

function AddMember() {
  if (document.getElementById('ad').style.display === "none"){
    document.getElementById('ad').style.display = "block";
    document.getElementById('id-add').value = document.getElementById('id').value;
  }else{
    document.getElementById('ad').style.display = "none";
    document.getElementById('id-add').value = "";
  }
}

function CreateMemberTable() {

  document.getElementById('membersTable').innerHTML = "";
  document.getElementById('id-add').value = document.getElementById('id').value;

  if (document.getElementById('id-add').value === document.getElementById('id').value){
    if (document.getElementById('membersTable').style.display === "none"){
      document.getElementById('membersTable').style.display = "block";
    }else{
      document.getElementById('membersTable').style.display = "none";
      document.getElementById('id-add').value = "";
    }
  }
  var data = JSON.stringify({"IdProject": parseInt(document.getElementById('id').value)});
  $.ajax({
      type:"POST",
      url: 'http://localhost:8080/adminmemberprojectjson/',
      data:data,
      success: function (response){
        CreateMemberProjectTableFromJSON(response);
      }
  });
}

function AddMember2(){
  var email = document.getElementById('memberemail').value
  var idproject = parseInt(document.getElementById('id-add').value)

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminmember/";
  xhr.open("POST", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Idproject": idproject, "Email": email});

  xhr.send(data);
  location.reload();
}

//Chua sua
function DeleteMember() {
  var iduser = parseInt(document.getElementById('iduser').value)
  var idproject = parseInt(document.getElementById('id-add').value)

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminmemberprojectjson/";
  xhr.open("DELETE", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"IdUser": iduser, "IdProject": idproject});

  xhr.send(data);
  location.reload();
}

//Cho project table
function PostProject() {
  var project = document.getElementById('new').value
  var mieuta = document.getElementById('desc').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminprojectjson/";
  xhr.open("POST", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Project": project, "Description": mieuta});

  xhr.send(data);
  location.reload();
}

function PutProject() {
  var id = document.getElementById('id').value
  var project = document.getElementById('project').value
  var mieuta = document.getElementById('mieuta').value

  var xhr = new XMLHttpRequest();
  var url = "http://localhost:8080/adminprojectjson/";
  xhr.open("PUT", url, true);
  xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
  var data = JSON.stringify({"Id": id, "Project": project, "Description": mieuta});
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
    console.log(data);
    CreateProjectTableFromJSON(data)

  })

});
