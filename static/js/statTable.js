function CreateStatTableFromJSON(data) {

  // Header
  var col = ["Project", "Number of Member", "Number of Bug", "Number of Solution", "Begin Date", "Finish Date"];
  var colJSON = ["Project", "Member", "Bug", "Solution", "BeginDate", "FinishDate"]; // de dong bo voi du lieu JSON

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

  ProjectsStatistics(data);

}

function ProjectsStatistics(data) {

  // var NameProject [];
  // for (var i = 0 ; i <data.length;i++) {
  //   NameProject[i] = data[i]["Project"]
  // }
  var categories = [];
  var numberBug = [];
  var numberMember = [];
  var numberSolution = [];
  for (var i = 0;i<data.length;i++) {
    categories.push(data[i].Project);
    numberBug.push(data[i].Bug);
    numberMember.push(data[i].Member);
    numberSolution.push(data[i].Solution)
  }

  Highcharts.chart('container', {
    chart: {
        type: 'bar'
    },
    title: {
        text: 'Projects'
    },
    subtitle: {
        text: ''
    },
    xAxis: {
        categories: categories,
        title: {
            text: null
        }
    },
    yAxis: {
        min: 0,
        title: {
            text: 'Number',
            align: 'high'
        },
        labels: {
            overflow: 'justify'
        }
    },
    tooltip: {
        valueSuffix: ''
    },
    plotOptions: {
        bar: {
            dataLabels: {
                enabled: true
            }
        }
    },
    legend: {
        layout: 'vertical',
        align: 'right',
        verticalAlign: 'top',
        x: -40,
        y: 80,
        floating: true,
        borderWidth: 1,
        backgroundColor: ((Highcharts.theme && Highcharts.theme.legendBackgroundColor) || '#FFFFFF'),
        shadow: true
    },
    credits: {
        enabled: false
    },
    series: [{
        name: 'Members',
        data: numberMember
    }, {
        name: 'Bugs',
        data: numberBug
    }, {
        name: 'Solutions',
        data: numberSolution
    }]
  });
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
