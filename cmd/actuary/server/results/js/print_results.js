// print_results.js handles the printing of all elements to the page

// Build the output of each test as a new row
function buildElement(item, divID){
	var row = document.createElement('div')
	row.className = "data-element"
	var name = $('<h5/>').addClass('name').text(item.Name + " - ")
	var status = $('<span/>').addClass('status ' + item.Status).text(item.Status)
	var output = $('<p> <em> </em> </p>').addClass('output').text(item.Output)
	$(name).append(status)
	$(row).append(name, output)
	document.getElementById(divID).appendChild(row)
}

// After tests have been analyzed, classify each node as either 1. passing 2. failing or 3. undetermined
function classifyNode(classification, nodeID){
	var old = $("#swarm-" + classification).text().split(" ")
	var count = String(parseInt(old[0]) + 1)
	$("#swarm-"+ classification).text(count + " " + old[1])
	$("#"+nodeID).addClass(classification).removeClass("node")
}

// Print the given node's output to page
// Classify each node based on test analysis
// Make each header clickable to filter tests
function printResults(passed, warned, skipped, info, nodeID){
	document.getElementById("stats-" + String(nodeID)).innerHTML = ""
	$("#" + "stats-" + String(nodeID)).append( 
		$('<h4>/>').addClass('stats').attr("id", "header-passed-" + nodeID).text(String(passed) + "% passed"),
		$('<h4>/>').addClass('stats').attr("id", "header-failed-" + nodeID).text(String(warned) + "% failed"),
		$('<h4>/>').addClass('stats').attr("id", "header-skipped-" + nodeID).text(String(skipped) + "% skipped"),
		$('<h4>/>').addClass('stats').attr("id", "header-info-" + nodeID).text(String(info) + " info only tests"))
	
	// ARBITRARY NUMBERS CURRENTLY FOR CLASSIFICATION
	if (skipped > 50){ 
		$("#" + "header-results-" + String(nodeID)).css('color', '#E1CA2C')
		classifyNode("undetermined", nodeID)
		
	} else if (passed < warned) {
		$("#" + "header-results-" + String(nodeID)).css('color', 'red')
		classifyNode("failing", nodeID)
	} else {
		$("#" + "header-results-" + String(nodeID)).css('color', 'green')
		classifyNode("passing", nodeID)
	}
	$("#header-results-" + nodeID).click(function(){
		if (nodeSelected != "") {
			id = nodeSelected.split("-")
			id = id[1] + "-" + id[2]
			document.getElementById(id).style.display = 'none';
			$("#" + nodeSelected).css({"font-weight": "500", "text-transform": "lowercase"})
		}
		document.getElementById("results-" + nodeID).style.display = 'block';
		$("#header-results-" +nodeID).css({"font-weight": "bold", "text-transform": "uppercase"})
		nodeSelected = this.id
		$("#test-details").text("Test Details for " + nodeSelected.split("-")[2] + ":")
	});
	$("#header-passed-" + nodeID).click(function(){
		if (nodeSelected != "") {
			id = nodeSelected.split("-")
			id = id[1] + "-" + id[2]
			document.getElementById(id).style.display = 'none';
			$("#" + nodeSelected).css({"font-weight": "500", "text-transform": "lowercase"})
		}
		document.getElementById("passed-" + nodeID).style.display = 'block';
		$("#header-passed-" +nodeID).css({"font-weight": "bold", "text-transform": "uppercase"})
		nodeSelected = this.id
		$("#test-details").text("Test Details for " + nodeSelected.split("-")[2] + ":")
	});
	$("#header-failed-" + nodeID).click(function(){
		if (nodeSelected != "") {
			id = nodeSelected.split("-")
			id = id[1] + "-" + id[2]
			document.getElementById(id).style.display = 'none';
			$("#" + nodeSelected).css({"font-weight": "500", "text-transform": "lowercase"})
		}
		document.getElementById("failed-" + nodeID).style.display = 'block';
		$("#header-failed-" +nodeID).css({"font-weight": "bold", "text-transform": "uppercase"})
		nodeSelected = this.id
		$("#test-details").text("Test Details for " + nodeSelected.split("-")[2] + ":")

	});
	$("#header-skipped-" + nodeID).click(function(){
		if (nodeSelected != "") {
			id = nodeSelected.split("-")
			id = id[1] + "-" + id[2]
			document.getElementById(id).style.display = 'none';
			$("#" + nodeSelected).css({"font-weight": "500", "text-transform": "lowercase"})
		}
		document.getElementById("skipped-" + nodeID).style.display = 'block';
		$("#header-skipped-" +nodeID).css({"font-weight": "bold", "text-transform": "uppercase"})
		nodeSelected = this.id
		$("#test-details").text("Test Details for " + nodeSelected.split("-")[2] + ":")

	});
	$("#header-info-" + nodeID).click(function(){
		if (nodeSelected != "") {
			id = nodeSelected.split("-")
			id = id[1] + "-" + id[2]
			document.getElementById(id).style.display = 'none';
			$("#" + nodeSelected).css({"font-weight": "500", "text-transform": "lowercase"})
		}
		document.getElementById("info-" + nodeID).style.display = 'block';
		$("#header-info-" +nodeID).css({"font-weight": "bold", "text-transform": "uppercase"})
		nodeSelected = this.id
		$("#test-details").text("Test Details for " + nodeSelected.split("-")[2] + ":")
	});
}