function CreateStatTableFromJSON(data) {

  // Header
  var col = ["Project", "Number of Member", "Number of Bug", "Number of Solution"];
  var colJSON = ["Project", "Member", "Bug", "Solution"]; // de dong bo voi du lieu JSON

  // Goi den bang can tim
  var table = document.getElementById("statTable");

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

$( document ).ready(function() {
  //Get JSON
  var url = 'http://localhost:8080/adminstatjson/'
  $.getJSON(url, function(data){
    console.log("It Worked!");
    CreateStatTableFromJSON(data)
    console.log(data);
  })

});
