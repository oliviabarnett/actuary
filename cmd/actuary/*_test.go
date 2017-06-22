// goal is for us to test all the functions that are being called
//probably mock responses from the Docker API to not have to test with a real daemon

package actuary


import (
	//"os"
	"testing"
	//"net/http" //Package http provides HTTP client and server implementations.
	//"net/http/httptest" //Package httptest provides utilities for HTTP testing.

	//"github.com/docker/docker/api/types"
	//"github.com/docker/docker/client"
	//"github.com/diogomonica/actuary/actuary"

)

//from https://gist.github.com/cyli/f565a5777183f664d78d7b4a2f3bb7be
// type TestingClient struct {
// 	cli    *client.Client
// 	name   string
// 	labels []string
// 	uuid   string
// }

// func GetClient() (*client.Client, error) {
// 	cli, err := client.NewClient(testServer().URL) // <-- fix. Used to be NewEnvClient
// 	if err != nil {
// 		return nil, err
// 	}
// 	return cli, nil
// }

// func NewTestingClient(name string, labels ...string) (*TestingClient, error) {
// 	client, err := GetClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &TestingClient{
// 		cli:    client,
// 		name:   name,
// 		labels: labels,
// 		uuid:   UUID(),
// 	}, nil
// }

// func testServer (t *testing.T, func()) { //inject a different response based on the test?
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		
		
// 	})

// 	server := httptest.NewServer(mux)
// }

var testTarget, err = NewTarget()




//1. host configuration

func TestCheckSeparatePartition(t *testing.T){
	//so I need a fake fstab?
	// t.Log("Creating dummy fstab")
	// fstab, _ := CreateFstab("/etc")
	// data := `/var/lib/docker`
	// fstab.Update(data)
	
	// if CheckSeparatePartition(actuary.NewTarget()) != "1.1 Create a separate partition for containers" {
	// 	t.Errorf("Should be correctly partitioned")
	// }

	// dummy.Destroy()
	
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

//func TestCheckDockerVersion(t *testing.T)  {

//}

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