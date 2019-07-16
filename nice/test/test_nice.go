package main

import (
	"go-licode/nice"
	"time"
)
func main() {
	agent := nice.NewNiceAgent()
	_ = agent
	stream_id := agent.Nice_agent_add_stream(1)
	agent.Nice_agent_gather_candidates(stream_id)

	//listener1, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//listener2, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//_ = listener1
	//_ = listener2

	for {
		time.Sleep(1)
	}
}
