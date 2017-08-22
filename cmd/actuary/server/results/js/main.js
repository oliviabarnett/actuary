// Keep track of nodeSelected 
var nodeSelected = ""
var dataSelected = ""

window.onload=function(){
	var domain = window.location.href.split("/")[0]
	
	// On page load, get official list of nodes in swarm for comparison
	getNodeList(domain).then(function(response){
		var nodeList = response.split(" ")
		if ((nodeList.length-1) == 1){
			var swarmAll = $('<h3>/>').attr("id", "swarm-all").addClass("pointer").text((nodeList.length-1) + " node")
		}else{
			var swarmAll = $('<h3>/>').attr("id", "swarm-all").addClass("pointer").text((nodeList.length-1) + " nodes")
		}
		$('#swarm-data').append(
			$('<div/>').attr("id", "swarm-stats").append(
				swarmAll,
				$('<ul></ul>').append(
					$('<li></li>').attr("id", "swarm-passing").addClass("pointer").text("0 passed"),
					$('<li></li>').attr("id", "swarm-failing").addClass("pointer").text("0 failed"),
					$('<li></li>').attr("id", "swarm-undetermined").addClass("pointer").text("0 undetermined")
				)
			)
		)
		// Clicking functionality for filtering nodes by passing, failing, undetermined
		$("#swarm-all").click(function(){
			if (dataSelected != "") {
				$("#"+ dataSelected).css({"font-weight": "500", "text-transform": "lowercase"})
			}
			$(".passing").show()
			$(".failing").show()
			$(".undetermined").show()
			$(".node").show()
			$("#swarm-all").css({"font-weight": "bold", "text-transform": "uppercase"})
			$("#nodes-header").text("All Nodes by ID:")
			dataSelected = this.id
		});
		$("#swarm-passing").click(function(){
			if (dataSelected != "") {
				$("#"+ dataSelected).css({"font-weight": "500", "text-transform": "lowercase"})
				$(".failing").hide()
				$(".undetermined").hide()
				$(".node").hide()
			}
			$(".passing").show()
			$("#swarm-passing").css({"font-weight": "bold", "text-transform": "uppercase"})
			$("#nodes-header").text("All Passing Nodes by ID:")
			dataSelected = this.id
		});
		$("#swarm-failing").click(function(){
			if (dataSelected != "") {
				$("#"+ dataSelected).css({"font-weight": "500", "text-transform": "lowercase"})
				$(".passing").hide()
				$(".undetermined").hide()
				$(".node").hide()
			}
			$(".failing").show()
			$("#swarm-failing").css({"font-weight": "bold", "text-transform": "uppercase"})
			$("#nodes-header").text("All Failing Nodes by ID:")
			dataSelected = this.id
		});
		$("#swarm-undetermined").click(function(){
			if (dataSelected != "") {
				$("#"+ dataSelected).css({"font-weight": "500", "text-transform": "lowercase"})
				$(".passing").hide()
				$(".failing").hide()
				$(".node").hide()
			}
			$(".undetermined").show()
			$("#swarm-undetermined").css({"font-weight": "bold", "text-transform": "uppercase"})
			$("#nodes-header").text("All Undetermined Nodes by ID:")
			dataSelected = this.id
		});
		// Initially, add all nodes with status "loading" before test information has been received
		for (i = 0; i < nodeList.length-1; i++){
				var nodeBox = $('<div/>').addClass('row node').attr("id", nodeList[i])
				var nodeHeader = $('<h4>/>').addClass("pointer").text(nodeList[i]).attr({"id": "header-results-" + nodeList[i], "style": "cursor: pointer;"})
				var nodeStats = $('<div/>').attr('id', "stats-" + String(nodeList[i])).text("Loading...")
				$(nodeBox).append(nodeHeader, nodeStats)
				$("#nodes-all").append(nodeBox)
		}
		// Check each node -- see if data has been recieved yet
		for (count = 0; count < nodeList.length-1; count++){
			pollList(nodeList[count], domain)
		}
	}), function(error) {
		console.log("failed getNodeList")
	}
}

// Get official list of nodes from server API call
function getNodeList(domain){
	url = domain + "/getNodeList"
	return new Promise((resolve, reject) => {
		var x = new XMLHttpRequest();
		x.open("Get", url);
		x.onload = () => resolve(x.responseText);
		x.onerror = () => reject(x.statusText);
		x.send();
	});
};

// If node data has been received, display, else wait and then try again
function pollList(node, domain){
	checkNode(domain, node).then(function(response){
		if (response[0] == "true"){
			getResults(domain, response[1])			
		}else if (response[0] == "false") {
			sleep(2000).then(() => {
				pollList(response[1], response[2])
			})
		}
	}), function(error, nodeID) {
		console.log("failed to get node " + response[1]+ " because " + error)
	}	
}	

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

// Call to server checks if data has been received on server side
function checkNode(domain, nodeID){
	url = domain + "/checkNode"
 	return new Promise((resolve, reject) => {
 		var x = new XMLHttpRequest();
 		x.open("POST", url);
 		x.setRequestHeader('Content-type', 'text/html')
 		x.onload = () => resolve([String(x.responseText), nodeID, domain]);
 		x.onerror = () => reject([String(x.statusText), nodeID, domain]);
 		x.send(nodeID);
 	});
};

// Get the output of the specified node from the server
function getResults(domain, nodeID){
	var urlParams = new URLSearchParams(window.location.search)
	var token = getCookie('token')
	if (token != "") {
		domain = domain + "/result"
		var x = new XMLHttpRequest()
		x.open("Get", domain + "?nodeID=" + nodeID) 
		x.setRequestHeader('Authorization', 'Bearer ' + token)
		x.onreadystatechange = function(){
			if (x.readyState == 4 && x.status == 200){
				var data = JSON.parse(x.responseText)
				analyzeResults(data, nodeID)
			}
		}
		x.send()
	}else{
		console.log("Could not retrieve token")
	}
}

function getCookie(cname){
	var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for(var i = 0; i <ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}
