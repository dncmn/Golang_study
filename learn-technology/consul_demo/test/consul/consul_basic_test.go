package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
	"testing"
)

var (
	ServiceName    = "game_server2"
	ServicePort    = 8300
	ServiceAddress = "http://127.0.0.1"
	ServiceIdJava  = "27c550e2-cecb-41d4-9d5c-fde0bae56716"
	client         = getConn()
)

func Test_health_check(T *testing.T) {
	id, _ := uuid.NewV4()

	check := &api.AgentServiceCheck{
		CheckID: id.String(),
		HTTP:    "http://127.0.0.1:8080/user/health_check",
		//HTTP:     "http://www.baidu.com",
		Name:     "baidu",
		Timeout:  "5s",
		Interval: "10s",
	}

	ser := &api.AgentServiceRegistration{
		ID:                id.String(),
		Name:              ServiceName,
		Tags:              []string{"for test"},
		Port:              ServicePort,
		Address:           ServiceAddress,
		EnableTagOverride: false,
		Check:             check,
	}

	err := client.Agent().ServiceRegister(ser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success ")
	fmt.Println(ser.ID)
}

func Test_ListServices(T *testing.T) {
	servies, err := client.Agent().Services()
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, ser := range servies {
		fmt.Printf("service:id=%v,port=%v\n", key, ser.Port)
	}
}

func Test_UnRegisterService(T *testing.T) {
	err := client.Agent().ServiceDeregister("f6c18bd1-626d-4a99-bbcc-db0288b00277")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Test_RegisterService(T *testing.T) {
	// 注册服务
	id, _ := uuid.NewV4()
	ser := &api.AgentServiceRegistration{
		ID:                id.String(),
		Name:              ServiceName,
		Tags:              []string{"for test"},
		Port:              ServicePort,
		Address:           ServiceAddress,
		EnableTagOverride: false,
		Check:             &api.AgentServiceCheck{},
	}

	err := client.Agent().ServiceRegister(ser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success ")
	fmt.Println(ser.ID)
}

func Test_GetLeader(T *testing.T) {
	client := getConn()

	// 获取leader信息
	leader, err := client.Status().Leader()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("leader is ", leader)
}

func Test_ConsulKV(T *testing.T) {

	client := getConn()
	valName := "maxConnections"
	// kv --put key:value
	p := &api.KVPair{Key: valName, Value: []byte("10000")}

	_, err := client.KV().Put(p, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// kv --get key:value
	val, _, err := client.KV().Get(valName, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(val.Value))
}

func getConn() (client *api.Client) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("%v", err)
		return
	}
	return
}
