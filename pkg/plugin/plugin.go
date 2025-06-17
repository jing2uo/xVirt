package plugin

import (
	"fmt"

	"github.com/libvirt/libvirt-go"
)

func Start() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("Failed to connect to libvirt:", err)
		return
	}
	defer conn.Close()

	pools, err := conn.ListAllStoragePools(0)
	if err != nil {
		fmt.Println("Failed to list storage pools:", err)
		return
	}
	for _, pool := range pools {
		name, _ := pool.GetName()
		fmt.Println("Storage Pool:", name)
		pool.Free()
	}
}
