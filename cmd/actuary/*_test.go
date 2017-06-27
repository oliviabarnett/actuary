// goal is for us to test all the functions that are being called
//probably mock responses from the Docker API to not have to test with a real daemon

package actuary


import (
	//"os"\
	"fmt"
	"encoding/json"
	//"log"
	//"io/ioutil"
	"testing"
	"strconv"
	//"io"
	"net/http" //Package http provides HTTP client and server implementations.
	"net/http/httptest" //Package httptest provides utilities for HTTP testing.
	"github.com/docker/engine-api/types"
	//"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api"
	//"github.com/gorilla/mux"
	//"github.com/diogomonica/actuary/actuary"


)

//variables for tests

type callPairing struct{
	call string
	obj []byte
}

type imageList struct{
	images []types.Image
}

func (list imageList) populateImageList(size int) (imageList){
	list.images = nil
	var img types.Image
	for i := 0; i < size; i++{
		img = types.Image{ID: strconv.Itoa(i)}
		list.images = append(list.images, img)
	}

	return list
}


var testTarget, err = NewTarget()

func testServer (t *testing.T, pairings ...callPairing) (server *httptest.Server) { //inject a different response based on the test?

	mux := http.NewServeMux()

	for _, pair := range pairings {
		mux.HandleFunc(
			fmt.Sprintf("/v1.26%s", pair.call),
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
				w.Header().Set("Content-Type", "application/json")
				w.Write(pair.obj)
				//log.Printf("Req: %s %s\n", r.Host, r.URL.Path)
		}))
	}
	server = httptest.NewServer(mux)

	//manipulate testTarget client to point to server
	testTarget.Client, err = client.NewClient(server.URL, api.DefaultVersion, nil, nil)


	if err != nil {
		t.Errorf("Could not manipulate test target client.")
		return 
	}
	return 
}

//1. host configuration

func TestCheckSeparatePartition(t *testing.T){
	
}

func TestCheckKernelVersion(t *testing.T) {
	//just checks info.KernelVersion of target. Fake info
	t.Log("Changing Kernel Version to 4.9.27-moby")

	testTarget.Info.KernelVersion = "4.9.27-moby"

	//t.Log("XXXXXXXX %i %i %i", testTarget.Info.ContainersRunning, testTarget.Info.ContainersPaused, testTarget.Info.ContainersStopped)

	res := CheckKernelVersion(testTarget)
	if  res.Status != "PASS" {
		t.Errorf("Kernel Version is correct, should have passed." )
	}
}

func TestCheckRunningServices(t *testing.T) {
	//not mac compatible
}

func TestCheckDockerVersion(t *testing.T) {
	
	var ver = types.Version{
		Version: "20000",
		Os: "linux",
		GoVersion: "go1.7.5",
		GitCommit: "deadbee",
	}

	for i := 0; i< 2; i++{
		vJSON, err := json.Marshal(ver)

		if err != nil {
			t.Errorf("Could not convert version to json.")
		}

		p := callPairing{"/version", vJSON}

		ts := testServer(t, p)
		
		res := CheckDockerVersion(testTarget)

		defer ts.Close()

		if i == 0 && res.Status != "PASS" {
			t.Errorf("Host using the correct Docker server, should pass" )
		}
		if i == 1 && res.Status == "PASS"{
			t.Errorf("Host not using the correct Docker server, should not pass" )
		}

		ver.Version = "0"
	}

 }

func TestCheckTrustedUsers(t *testing.T) {

}

func TestAuditDockerDaemon(t *testing.T) {
	
}

func TestAuditLibDocker(t *testing.T) {
	
}

func TestAuditEtcDocker(t *testing.T)  {
	
}

func TestAuditDockerService(t *testing.T)  {
	
}

func TestAuditDockerSocket(t *testing.T)  {
	
}

func TestAuditDockerDefault(t *testing.T)  {
	
}

func TestAuditDaemonJSON(t *testing.T)  {
	
}

func TestAuditContainerd(t *testing.T)  {

}

func TestAuditRunc(t *testing.T) {
	
}

//2. Docker daemon configuration

func TestRestrictNetTraffic(t *testing.T) {
	//calls .NetworkList(context.TODO(), netargs), fake a response network
	
	var network = []types.NetworkResource{{
			Name: "bridge",
			Options: map[string]string{"com.docker.network.bridge.enable_icc": "false"},
	}}



	for i := 0; i< 2; i++{
		nJSON, err := json.Marshal(network)

		if err != nil {
			t.Errorf("Could not convert network to json.")
		}

		p := callPairing{"/networks", nJSON}

		ts := testServer(t, p)
		
		res := RestrictNetTraffic(testTarget)

		defer ts.Close()

		if i == 0 && res.Status != "PASS" {
			t.Errorf("Net traffic restricted, should pass" )
		}
		if i == 1 && res.Status == "PASS"{
			t.Errorf("Net traffic not restricted, should not pass" )
		}

		//test fail case
		network[0].Options["com.docker.network.bridge.enable_icc"] = "true"
	}

 }

func TestCheckLoggingLevel(t *testing.T) {

}

func TestCheckIpTables(t *testing.T) {
	
}

func TestCheckInsecureRegistry(t *testing.T) {
	
}

func TestCheckAufsDriver(t *testing.T) {
	
}

func TestCheckTLSAuth(t *testing.T) {
	
}

func TestCheckUlimit(t *testing.T) {
	
}

func TestCheckUserNamespace(t *testing.T) {
	
}

func TestCheckDefaultCgroup(t *testing.T) {
	
}

func TestCheckBaseDevice(t *testing.T) {
	
}

func TestCheckAuthPlugin(t *testing.T) {
	
}

func TestCheckCentralLogging(t *testing.T) {
	
}

func TestCheckLegacyRegistry(t *testing.T) {
	
}


//3. Docker daemon configuration files

//4. Container Images and Build File

func TestCheckContainerUser(t *testing.T) {
	//all containers should have a not blank user?
	t.Log("Setting all container users")

	containers := testTarget.Containers

	//t.Log("XXXXXXX %i", len(testTarget.Containers))

	for _, container := range containers {
		container.Info.Config.User = "x"
	}

	res := CheckContainerUser(testTarget)

	if  res.Status != "PASS" {
		t.Errorf("All users checked, should have passed." )
	}
}

//5. Container Runtime

func TestCheckAppArmor(t *testing.T) {
	//all containers should have AppArmor profile
	for _, container := range testTarget.Containers {
		var cinfo = container.Info
		cinfo.AppArmorProfile = "AppArmor"
	}

	res := CheckAppArmor(testTarget)

	if  res.Status != "PASS" {
		t.Errorf("All containers have AppArmor, should have passed." )
	}
}

func TestCheckSELinux(t *testing.T) {
	filler := []string{"SELinux", "Array"}
	for _, container := range testTarget.Containers {
		var cinfo = container.Info
		cinfo.HostConfig.SecurityOpt = filler
		//t.Log("XXXXXXX", cinfo.HostConfig.SecurityOpt)
	}

	res := CheckSELinux(testTarget)

	if  res.Status != "PASS" {
		t.Errorf("All containers have SELinux options, should have passed." )
	}
}

func TestCheckKernelCapabilities(t *testing.T) {
	
	for _, container := range testTarget.Containers {
		var cinfo = container.Info
		cinfo.HostConfig.CapAdd =  nil
	}

	res := CheckKernelCapabilities(testTarget)

	if  res.Status != "PASS" {
		t.Errorf("No containers running with added capabilities, should have passed." )
	}
}

func TestCheckPrivContainers(t *testing.T) {
	for _, container := range testTarget.Containers {
		var cinfo = container.Info
		cinfo.HostConfig.Privileged = false 
	}

	res := CheckKernelCapabilities(testTarget)

	if  res.Status != "PASS" {
		t.Errorf("No containers are privileged, should have passed." )
	}	
}

func TestCheckSensitiveDirs(t *testing.T) {
	
}

func TestCheckSSHRunning(t *testing.T) {
// GET "/containers/{name:.*}/top"
	var processList = types.ContainerProcessList{
					Titles: []string{"UID", "PID","PPID","C","STIME","TTY","TIME","CMD"},
					Processes: [][]string{{"root","13642","882","0","17:03","pts/0","00:00:00","/bin/bash"}, 
										{"root", "13735","13642","0","17:06","pts/0","00:00:00","sleep 10"}},
					}

	temp := testTarget.Containers
	testTarget.Containers =  ContainerList{testTarget.Containers[0]}

	for i := 0; i< 2; i++{
		pJSON, err := json.Marshal(processList)

		if err != nil {
			t.Errorf("Could not convert process list to json.")
		}

		p := callPairing{"/containers/" + testTarget.Containers[0].ID +"/top", pJSON}

		ts := testServer(t, p)
	
		res := CheckSSHRunning(testTarget)

		defer ts.Close()

		if i == 0 && res.Status != "PASS" {
			t.Errorf("No containers running SSH service, should pass" )
		}
		if i == 1 && res.Status == "PASS"{
			t.Errorf("Container running SSH service, should not pass" )
		}
		
		//test fail case
		processList.Processes[0][3] = "ssh"
	}

		testTarget.Containers = temp
 }

func TestCheckPrivilegedPorts(t *testing.T) {
	
}

func TestCheckNeededPorts(t *testing.T) {
	
}

func TestCheckHostNetworkMode(t *testing.T) {
	
}

func TestCheckMemoryLimits(t *testing.T) {
	
}

func TestCheckCPUShares(t *testing.T) {
	
}

func TestCheckReadonlyRoot(t *testing.T) {
	
}

func TestCheckBindHostInterface(t *testing.T) {
	
}

func TestCheckRestartPolicy(t *testing.T) {
	
}

func TestCheckHostNamespace(t *testing.T) {
	
}

func TestCheckIPCNamespace(t *testing.T) {
	
}

func TestCheckHostDevices(t *testing.T) {
	
}

func TestCheckDefaultUlimit(t *testing.T) {
	
}

func TestCheckMountPropagation(t *testing.T) {
	
}

func TestCheckUTSnamespace(t *testing.T) {
	
}

func TestCheckSeccompProfile(t *testing.T) {
	
}

func TestCheckCgroupUsage(t *testing.T) {

}

func TestCheckAdditionalPrivs(t *testing.T) {
	
}

//6. Docker Security Operations

func TestCheckImageSprawl(t *testing.T) {
//GET /images/json
	var imgList imageList

	imgs := imgList.populateImageList(2).images

	container1 := types.Container{ImageID: "1"}
	container2 := types.Container{ImageID: "2"}

	containerLst := []types.Container{container1, container2}

	for i := 0; i< 2; i++{
		imagesJSON, err := json.Marshal(imgs)
		containerJSON, err := json.Marshal(containerLst)

		if err != nil {
			t.Errorf("Could not convert process list to json.")
		}

		p1 := callPairing{ "/containers/json", containerJSON}
		p2 := callPairing{ "/images/json", imagesJSON}

		ts := testServer(t, p1, p2)
		
		res := CheckImageSprawl(testTarget)

		defer ts.Close()

		if i == 0 && res.Status != "PASS" {
			t.Errorf("Correct amount of images, should pass." )
		}
		if i == 1 && res.Status == "PASS"{
			t.Errorf("Over 100 images, should not pass." )
		}

		//test fail case
		imgs = imgList.populateImageList(105).images
	}
}

func TestCheckContainerSprawl(t *testing.T) {

}








