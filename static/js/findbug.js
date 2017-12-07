//Tao them bang khi an vao show all bugs
function CreateFindBugTableFromJSON(data) {

  // Header
  var col = ["#","Bug", "Project","Description", "Solution", "Tester","Found Date", "Update Date"];
  var colJSON = ["Id", "BugName", "Project", "BugDescription", "SolutionDescription", "User","FoundDate", "UpdateDate"]; // de dong bo voi du lieu JSON

  // Goi den bang can tim
  var table = document.getElementById("findbug");

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

}


function CreateBugTable() {

  document.getElementById('findbug').innerHTML = "";

  var data = JSON.stringify({"BugName": document.getElementById('bugname').value});
  $.ajax({
      type:"POST",
      url: 'http://localhost:8080/user/findbug/',
      data:data,
      success: function (response){
        CreateFindBugTableFromJSON(response);
      }
  });
}
