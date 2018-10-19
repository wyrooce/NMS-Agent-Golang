function loadSoftwareList() {

            $.post("http://localhost:9000/softwareList", function (data) {
                console.debug("salaam");
                console.debug(data);
                var table = document.getElementById("tblSoftwareList");

                var i;
                for (i = 0; i < data.length; i++) {
                    // newRow = "<tr><td>"+i+"</td><td>" + data[i].name + "</td><td>" + data[i].publisher + "</td><td>" + data[i].version + "</td><td>" + data[i].installDate + "</td><td>" + data[i].size + "</td><td>" + data[i].architecture + "</td></tr>";
                    // $("#tblSoftwareList").append(newRow)
                    var row = table.insertRow(i+1);
                    row.insertCell(0).innerHTML = i+1;
                    row.insertCell(1).innerHTML = data[i].name
                    row.insertCell(2).innerHTML = data[i].publisher
                    row.insertCell(3).innerHTML = data[i].version
                    row.insertCell(4).innerHTML = data[i].installDate
                    row.insertCell(4).innerHTML = data[i].size
                    row.insertCell(4).innerHTML = data[i].architecture

                }
            });
        }