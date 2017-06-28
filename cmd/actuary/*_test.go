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
	//"github.com/docker/go-connections/nat"	
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

type typeContainerList struct{
	typeContainers []types.Container
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

func (list typeContainerList) populateContainerList(size int) (typeContainerList){
	list.typeContainers = nil
	var c types.Container
	for i := 0; i < size; i++{
		c = types.Container{ID: strconv.Itoa(i)}
		list.typeContainers = append(list.typeContainers, c)
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

	res := CheckKernelVersion(testTarget)
	if  res.Status != "PASS" {
		t.Errorf("Kernel Version is correct, should have passed." )
	}

	testTarget.Info.KernelVersion = "1.9.27-moby"

	res = CheckKernelVersion(testTarget)
	if  res.Status == "PASS" {
		t.Errorf("Kernel Version is incorrect, should not have passed." )
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
	//for all container runtime tests, simplify to one container

func containerTestsHelper(t *testing.T, orig func(t Target) Result, f func(c Container, i int) Container, err1 string, err2 string){

	temp := testTarget.Containers
	testTarget.Containers = ContainerList{testTarget.Containers[0]}

	//log.Printf("%v", testTarget)

	for i := 0; i< 2; i++{	

		testTarget.Containers[0] = f(testTarget.Containers[0], i)

		res := orig(testTarget)

		if i == 0 && res.Status != "PASS" {
			t.Errorf(err1)
		}
		if i == 1 && res.Status == "PASS"{
			t.Errorf(err2)
		}	
	}

	//restore
	testTarget.Containers = temp 
}


func TestCheckAppArmor(t *testing.T) {
	//all containers should have AppArmor profile

	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.AppArmorProfile = "app armor"
		}else {
			c.Info.AppArmorProfile = ""
		}

		return c
	}	

	containerTestsHelper(t, CheckAppArmor, f, "All containers have app armor profile, should pass.", "Container without app armor profile, should not pass.")
}

func TestCheckSELinux(t *testing.T) {

	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.SecurityOpt = []string{"SELinux", "Array"}
		}else {
			c.Info.HostConfig.SecurityOpt = nil
		}

		return c
	}	

	containerTestsHelper(t, CheckSELinux, f, "All containers have SELinux options, should have passed.", "No containers have SELinux options, should not have passed.")
}


func TestCheckKernelCapabilities(t *testing.T) {
	
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.CapAdd =  nil
		}else {
			c.Info.HostConfig.CapAdd = []string{"added", "capabilities"}
		}

		return c
	}	

	containerTestsHelper(t, CheckKernelCapabilities, f, "No containers running with added capabilities, should have passed.", "Containers running with added capabilities, should not have passed.")
}


func TestCheckPrivContainers(t *testing.T) {

	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.Privileged  =  false
		}else {
			c.Info.HostConfig.Privileged  = true
		}

		return c
	}	
	containerTestsHelper(t, CheckPrivContainers, f, "No containers are privileged, should have passed.", "Containers are privileged, should not have passed.")
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

	// f := func(c Container, i int) (Container) { 
	// 	if i == 0 {
	// 		c.Info.NetworkSettings.Ports = nat.PortMap{"80/tcp": []PortBinding{PortBinding{"hostIp", "2000"}, PortBinding{"hostIp", "2000"}}}
	// 	}else {
	// 		c.Info.NetworkSettings.Ports = nat.PortMap{"80/tcp": []PortBinding{PortBinding{"hostIp", "1000"}, PortBinding{"hostIp", "1000"}}}

	// 	}

	// 	return c
	// }	
	// containerTestsHelper(t, CheckPrivilegedPorts, f, "No ports are privileged, should have passed.", "Ports are privileged, should not have passed.")
}	


func TestCheckNeededPorts(t *testing.T) {
	
}

func TestCheckHostNetworkMode(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.NetworkMode  = ""
		}else {
			c.Info.HostConfig.NetworkMode = "host"
		}

		return c
	}	
	containerTestsHelper(t, CheckHostNetworkMode, f, "No containers are privileged, should have passed.", "Containers are privileged, should not have passed.")
}

func TestCheckMemoryLimits(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.Memory  = 10.0
		}else {
			c.Info.HostConfig.Memory = 0
		}

		return c
	}	
	containerTestsHelper(t, CheckMemoryLimits, f, "No containers have unlimited memory, should have passed.", "Container has unlimited memory, should not have passed.")
}

func TestCheckCPUShares(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.CPUShares = 100
		}else {
			c.Info.HostConfig.CPUShares = 0
		}

		return c
	}	
	containerTestsHelper(t, CheckCPUShares, f, "No containers with CPU sharing disable, should have passed.", "Containers with CPU sharing disabled, should not have passed.")
}

func TestCheckReadonlyRoot(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.ReadonlyRootfs = true
		}else {
			c.Info.HostConfig.ReadonlyRootfs = false
		}

		return c
	}	
	containerTestsHelper(t, CheckReadonlyRoot, f, "Containers all have read only root filesystem, should have passed.", "Containers' root FS is not mounted as read-only, should not have passed.")
}

func TestCheckBindHostInterface(t *testing.T) {
	
}

func TestCheckRestartPolicy(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.RestartPolicy.Name = "on-failure"
			c.Info.HostConfig.RestartPolicy.MaximumRetryCount = 5
		}else {
			c.Info.HostConfig.RestartPolicy.Name = ""
			c.Info.HostConfig.RestartPolicy.MaximumRetryCount = 0
		}

		return c
	}	
	containerTestsHelper(t, CheckRestartPolicy, f, "Containers all have restart policy set to 5, should have passed.", "Containers with no restart policy, should not have passed.")
}

func TestCheckHostNamespace(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.PidMode = ""
		}else {
			c.Info.HostConfig.PidMode = "host"		
		}

		return c
	}	
	containerTestsHelper(t, CheckHostNamespace, f, "Containers do not share the host's process namespace, should have passed.", "Containers sharing host's process namespace, should not have passed.")
}

func TestCheckIPCNamespace(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.IpcMode = ""
		}else {
			c.Info.HostConfig.IpcMode = "host"		
		}

		return c
	}	
	containerTestsHelper(t, CheckIPCNamespace, f, "Containers do not share the host's IPC namespace, should have passed.", "Containers sharing host's IPC namespace, should not have passed.")
}

func TestCheckHostDevices(t *testing.T) {
// 	f := func(c Container, i int) (Container) { 

// 		if i == 0 {
// 			c.Info.HostConfig.Devices = DeviceMapping{"1", "2", "3"}
// 		}else {
// 			c.Info.HostConfig.Devices = nil		
// 		}

// 		return c
// 	}	
// 	containerTestsHelper(t, CheckHostDevices, f, "Host devices not exposed to containers, should have passed.", "Host devices directly exposed to containers, should not have passed.")
 }

func TestCheckDefaultUlimit(t *testing.T) {
	// f := func(c Container, i int) (Container) { 

	// 	if i == 0 {
	// 		c.Info.HostConfig.Ulimits = 
	// 	}else {
	// 		c.Info.HostConfig.Ulimits = nil		
	// 	}

	// 	return c
	// }	
	// containerTestsHelper(t, CheckIPCNamespace, f, "Containers do not override default ulimit, should have passed.", "Containers overriding default ulimit, should not have passed.")
}

func TestCheckMountPropagation(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			mounts := c.Info.Mounts 
			for _, mount := range mounts {
				mount.Mode = "" 
			}	
		}else {
			mounts := c.Info.Mounts 
			mounts[0].Mode = "shared"
		}

		return c
	}	
	containerTestsHelper(t, CheckMountPropagation, f, "Mount propagation mode not set to shared, should have passed.", "Containers have mount propagation set to shared, should not have passed.")
}

func TestCheckUTSnamespace(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.UTSMode = ""
		}else {
			c.Info.HostConfig.UTSMode = "host"
		}

		return c
	}	
	containerTestsHelper(t, CheckUTSnamespace, f, "Containers do not share host's UTS namespace, should have passed.", "Containers share host's UTS namespace, should not have passed.")
}

func TestCheckSeccompProfile(t *testing.T) {
	f := func(c Container, i int) (Container) { 

		if i == 0 {
			c.Info.HostConfig.SecurityOpt = []string{"seccomp", "not disabled"}
		}else {
			c.Info.HostConfig.SecurityOpt = []string{"seccomp:unconfined"}
		}

		return c
	}	
	containerTestsHelper(t, CheckSeccompProfile, f, "Seccomp not disabled, should have passed.", "Containers running with seccomp disabled, should not have passed.")
}

func TestCheckCgroupUsage(t *testing.T) {
	f := func(c Container, i int) (Container){

		if i == 0 {
			c.Info.HostConfig.CgroupParent = ""
		}else {
			c.Info.HostConfig.CgroupParent = "cgroup"}

		return c
	}
		
	containerTestsHelper(t, CheckCgroupUsage, f, "Containers all using default cgroup, should have passed.", "Container not using default cgroup, should not have passed.")
}

func TestCheckAdditionalPrivs(t *testing.T) {
	f := func(c Container, i int) (Container){

		if i == 0 {
			c.Info.HostConfig.SecurityOpt = []string{"no-new-privileges"}
		}else {
			c.Info.HostConfig.SecurityOpt = []string{""}
		}

		return c
	}
		
	containerTestsHelper(t, CheckAdditionalPrivs, f, "Containers restricted from aquiring additional privileges, should have passed.", "Containers unrestricted from acquiring additional privileges, should not have passed.")
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
	var containerList1 typeContainerList
	//var containerList2 typeContainerList

	list1 := containerList1.populateContainerList(10).typeContainers
	//containerList2 = containerList2.populateContainerList(10)	

	for i := 0; i< 2; i++{

		containerJSON1, err := json.Marshal(list1)
		//containerJSON2, err := json.Marshal(containerList2)

		if err != nil {
			t.Errorf("Could not convert process list to json.")
		}

		p1 := callPairing{ "/containers/json", containerJSON1}
		//p2 := callPairing{ "/containers/json?all=true", containerJSON2}

		ts := testServer(t, p1)
		
		res := CheckContainerSprawl(testTarget)

		defer ts.Close()

		if i == 0 && res.Status != "PASS" {
			t.Errorf("Sprawl less than 25, should pass.")
		}
		// if i == 1 && res.Status == "PASS"{
		// 	t.Errorf("More than 25 containers not running, should not pass." )
		// }

		//test fail case
		//containerList2 = containerList2.populateContainerList(50)	
	}
}







