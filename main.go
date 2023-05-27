package main

type image struct{}

type container struct{}

type netns struct{}

type limits struct {
	CPU    int
	Memory int
	//...
}

type runParams struct {
	Limits limits
	Env    map[string]string
	Mounts interface{}
	Ports  []int
}

func Run(imageName string, params runParams) int {
	image := downloadImage(imageName)

	container := createContainer(params.Limits, params.Env)
	defer cleanup(container)

	setupFS(container, image, params.Mounts)

	ns := setupNetwork(container, params.Ports)

	result := runContainer(container, ns)

	return result
}

func downloadImage(imageName string) image {
	// check if image exists in docker cache and get the latest downloaded layer
	// updates = getUpdates(imageName, latestLayer)
	// if len(updates) > 0 { downloadUpdates(updates) }

	return image{}
}

func createContainer(Limits limits, Env map[string]string) container {
	// mkdir /sys/fs/cgroup/cpu/<CONTAINER NAME>
	// echo <CPU Limit> > /sys/fs/cgroup/cpu/<CONTAINER NAME>/cpu.cfs_quota_us

	// mkdir /sys/fs/cgroup/memory/<CONTAINER NAME>
	// echo <Memory Limit> > /sys/fs/cgroup/memory/<CONTAINER NAME>/memory.limit_in_bytes

	// TODO cgroup of network and io

	// initialize container's environment variables
	// for each k,v in Env:
	// 		export key=value

	return container{}
}

func setupFS(container container, image image, mounts interface{}) {
	// prepare directories
	// mkdir -p ~/minidocker/base_image ~/minidocker/base-image-overlay ~/minidocker/fs

	// mount docker image
	// mount -o loop <IMAGE> ~/minidocker/base_image

	// configure unionfs to simulate writable layer
	// mount -t unionfs -o dirs=~/minidocker/base-image-overlay=rw:~/minidocker/base-image=ro unionfs ~/minidocker/fs

	// count additional directories for each mountpoint
	// mount --bind <MOUNTPOINT HOST> ~/minidocker/fs/<MOUNTPOINT>

	// set ~/minidocker/fs as root directory
	// chroot ~/minidocker/fs

}

func setupNetwork(container container, ports []int) netns {
	// create a new network namespace
	// ip netns add <CONTAINER NAME>_network

	// add virtual interface to a new network namespace
	// ip link add link eth0 <CONTAINER NAME>_int netns <CONTAINER NAME>_network type ipvlan mode l2

	// put up the new interface and loopback interface
	// ip -n <CONTAINER NAME>_network link set lo up
	// ip -n <CONTAINER NAME>_network link set <CONTAINER NAME>_int up

	// assign IP address to the interface
	// ip -n <CONTAINER NAME>_network addr add 192.168.0.78/24 dev <CONTAINER NAME>_int

	// configure default route
	// ip -n <CONTAINER NAME>_network route add default via 192.168.0.1 dev <CONTAINER NAME>_int

	// configure DNS settings for the namespace (add a default DNS server)
	// echo "nameserver 8.8.8.8" > /etc/netns/<CONTAINER NAME>_network/resolv.conf
}

func runContainer(container container, ns netns) int {
	// create a new user and PID namespaces, fork a child process to execute the entrypoint
	// unshare --net=/var/run/netns/<CONTAINER NAME>_network --user --pid --map-root-user --mount-proc --fork <ENTRY POINT>

	// assign process to cgroups
	// echo <PID> > /sys/fs/cgroup/memory/<CONTAINER NAME>/cgroup.procs
	// echo <PID> > /sys/fs/cgroup/cpu/<CONTAINER NAME>/cgroup.procs

	return 0 /* return code */
}

func cleanup(container container) {
	// ip netns del <CONTAINER NAME>_network
}

func main() {
	Run("<image>", runParams{ /*...*/ })
}
