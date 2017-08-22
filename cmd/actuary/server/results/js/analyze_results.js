// Determine the number of passing, failing, skipped, info only tests for each node

function analyzeResults(response, nodeID){
	var passed = 0
	var warned = 0
	var skipped = 0
	var info = 0
	var total = response.length
	var resultsID = "results-" + String(nodeID)
	var passedID = "passed-" + String(nodeID)
	var failedID = "failed-" + String(nodeID)
	var skippedID = "skipped-" + String(nodeID)
	var infoID = "info-" + String(nodeID)
	var results = $('<div />').addClass('tabContent scroll-box').attr("id", resultsID)
	var passDiv = $('<div />').addClass('tabContent scroll-box').attr("id", passedID)
	var failDiv = $('<div />').addClass('tabContent scroll-box').attr("id", failedID)
	var skipDiv = $('<div />').addClass('tabContent scroll-box').attr("id", skippedID)
	var infoDiv = $('<div />').addClass('tabContent scroll-box').attr("id", infoID)
	$("#data").append(results)
	$("#data").append(passDiv)
	$("#data").append(failDiv)
	$("#data").append(skipDiv)
	$("#data").append(infoDiv)
 	$.each(response, function(index, item){
		buildElement(item, resultsID)
		switch(item.Status){
			case "PASS":
				buildElement(item, passedID)
				passed++
				break;
			case "WARN":
				buildElement(item, failedID)
				warned++
				break;
			case "SKIP":
				buildElement(item, skippedID)
				skipped++
				break;
			case "INFO":
				buildElement(item, infoID)
				info++
				break;
		}
	});	
	total = total - info
	passed = Math.round(passed/total *100)
	warned = Math.round(warned/total *100)
	skipped = Math.round(skipped/total *100)
	printResults(passed, warned, skipped, info, nodeID)
}