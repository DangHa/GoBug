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

}

function ProjectsStatistics(data) {

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

  Highcharts.chart('project', {
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

function BugStatistics(data) {

  var categories = [];
  var numberBug = [];
  var numberProject = [];
  for (var i = 0;i<data.length;i++) {
    categories.push(data[i].Category);
    numberBug.push(data[i].Bug);
    numberProject.push(data[i].Project);
  }

  Highcharts.chart('bug', {
    chart: {
        type: 'bar'
    },
    title: {
        text: 'Category'
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
        name: 'Projects',
        data: numberProject
    }, {
        name: 'Bugs',
        data: numberBug
    }]
  });
}

function DevStatistics(data) {

  var developers = []; // 8 nguoi cao nhat solu/project ra thong ke
  var numberProject = [];
  var numberSolution = [];

  var max = 0;
  for (var i = 0;i<data.length;i++) {
    developers.push(data[i].UserName);
    numberProject.push(data[i].Project);
    numberSolution.push(data[i].Solution);
    if (max < numberSolution[i]/(numberProject[i]+1)) {
      max = numberSolution[i]/(numberProject[i]+1)
    }
  }



  Highcharts.chart('develop', {
    chart: {
        type: 'column'
    },
    title: {
        text: 'Developers'
    },
    subtitle: {
        text: develop[max]
    },
    xAxis: {
        categories: developers,
        crosshair: true
    },
    yAxis: {
        min: 0,
        title: {
            text: ""
        }
    },
    tooltip: {
        headerFormat: '<span style="font-size:17px">{point.key}</span><table>',
        pointFormat: '<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
            '<td style="padding:0"><b>{point.y:.1f} </b></td></tr>',
        footerFormat: '</table>',
        shared: true,
        useHTML: true
    },
    plotOptions: {
        column: {
            pointPadding: 0.2,
            borderWidth: 0
        }
    },
    series: [{
        name: 'Project',
        data: numberProject

    }, {
        name: 'Solution',
        data: numberSolution

    }]

  });

  console.log("Project: "+numberProject);
  console.log("Solution: "+numberSolution);

  // Highcharts.chart('dev', {
  //   chart: {
  //       type: 'bar'
  //   },
  //   title: {
  //       text: 'Developers'
  //   },
  //   subtitle: {
  //       text: ''
  //   },
  //   xAxis: {
  //       categories: dev,
  //       title: {
  //           text: null
  //       }
  //   },
  //   yAxis: {
  //       min: 0,
  //       title: {
  //           text: 'Number',
  //           align: 'high'
  //       },
  //       labels: {
  //           overflow: 'justify'
  //       }
  //   },
  //   tooltip: {
  //       valueSuffix: ''
  //   },
  //   plotOptions: {
  //       bar: {
  //           dataLabels: {
  //               enabled: true
  //           }
  //       }
  //   },
  //   legend: {
  //       layout: 'vertical',
  //       align: 'right',
  //       verticalAlign: 'top',
  //       x: -40,
  //       y: 80,
  //       floating: true,
  //       borderWidth: 1,
  //       backgroundColor: ((Highcharts.theme && Highcharts.theme.legendBackgroundColor) || '#FFFFFF'),
  //       shadow: true
  //   },
  //   credits: {
  //       enabled: false
  //   },
  //   series: [{
  //       name: 'Project',
  //       data: numberProject
  //   }, {
  //       name: 'Solution',
  //       data: numberSolution
  //   }]
  // });
}

$( document ).ready(function() {
  //Get JSON
  var url = 'http://localhost:8080/adminstatjson/'
  $.getJSON(url, function(data){
    console.log("It1 Worked!");
    CreateStatTableFromJSON(data)
    ProjectsStatistics(data);
  })

  url = 'http://localhost:8080/admin/bugstat/json/'
  $.getJSON(url, function(data){
    console.log("It2 Worked!");
    BugStatistics(data);
  })

  url = 'http://localhost:8080/admin/devstat/json/'
  $.getJSON(url, function(data){
    console.log("It3 Worked!");
    console.log(data);
    DevStatistics(data);
  })
});
