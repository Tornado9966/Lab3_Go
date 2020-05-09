package main

import (
	"testing"
	. "gopkg.in/check.v1"
)
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s * TestSuite) TestBalancer(c *C){

	serversPool = make(map[int]string)
	serversPool[0] = "server1:8080"
	serversPool[1] = "server2:8080"
	serversPool[2] = "server3:8080"
	
	//Tests for the health function
	for i := 0; i < len(serversPool); i++ {
		res := health(serversPool[i])
		c.Assert(res, Equals, false)
	}
	
	//Tests for the find function
	sl := []int{1, 2, 3}
		
	for i := 0; i < len(sl); i++ {
		num, boolValue := find(sl, i+1)
		c.Assert(num, Equals, i)
		c.Assert(boolValue, Equals, true)
	}
	
	//Tests for the chooseServer function
	healthyServersPool = []int{0, 1, 2}
	
	addr := []string{"172.20.0.3:45234", 
                         "172.20.0.3:45234",
                         "172.20.0.4:34563",
                         "172.20.0.2:34563",
                         "172.20.0.4:34563"}
	
	for i := 0; i < len(addr); i++ {
		res := chooseServer(addr[i]) 
		if i == 0 || i == 1 {
			c.Assert(res, Equals, "server2:8080")
		}
		if i == 2 || i == 4 {
			c.Assert(res, Equals, "server3:8080")
		}
		if i == 3 {
			c.Assert(res, Equals, "server1:8080") 
		}
	}
	
	//Tests for the remove function
	arr := []int{1, 2, 3}
	
	for i := 0; i < len(arr); i++ {
		arr = remove(arr, i+1)
		if i == 0 {
			c.Assert(arr[i], Equals, 2)
			c.Assert(arr[i+1], Equals, 3)
			c.Assert(len(arr), Equals, 2)
		}
		if i == 1 {
			c.Assert(arr[i-1], Equals, 3)
			c.Assert(len(arr), Equals, 1)
		}
		if i == 2 {
			c.Assert(len(arr), Equals, 0)
		}
	}
}
