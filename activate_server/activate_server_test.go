package activate_server

import (
	"fmt"
	"testing"
)

func TestActivateServer(t *testing.T) {
	opts := NewOptions()
	fmt.Println("options:", opts)
	NewActivateServer(opts).Run()
}

func TestOptionsStore(t *testing.T) {
	opts := NewOptions()
	opts.Store()
	fmt.Println("options:", opts)
}
