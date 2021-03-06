// +build integration

package integration

import (
	"context"
	"github.com/cloudscale-ch/cloudscale-go-sdk"
	"reflect"
	"testing"
)


func TestIntegrationSubnet_GetAndList(t *testing.T) {
	integrationTest(t)

	createNetworkRequest := &cloudscale.NetworkCreateRequest{
		Name: networkBaseName,
	}

	network, err := client.Networks.Create(context.TODO(), createNetworkRequest)
	if err != nil {
		t.Fatalf("Networks.Create returned error %s\n", err)
	}

	expectedNumberOfSubnets := 1
	if numberOfSubnets := len(network.Subnets); numberOfSubnets != expectedNumberOfSubnets {
		t.Errorf("Number of Subnets got=%#v\nwant=%#v", numberOfSubnets, expectedNumberOfSubnets)
	}

	subnet, err := client.Subnets.Get(context.Background(), network.Subnets[0].UUID)
	if err != nil {
		t.Fatalf("Subnets.Get returned error %s\n", err)
	}

	if uuid := subnet.UUID; uuid != network.Subnets[0].UUID {
		t.Errorf("Subnet.UUID got=%s\nwant=%s", uuid, network.Subnets[0].UUID)
	}

	if networkUUID := subnet.Network.UUID; networkUUID != network.UUID {
		t.Errorf("subnet.Network.UUID got=%s\nwant=%s", networkUUID, network.UUID)
	}

	if networkUUID := subnet.Network.UUID; networkUUID != network.UUID {
		t.Errorf("subnet.Network.UUID got=%s\nwant=%s", networkUUID, network.UUID)
	}

	subnets, err := client.Subnets.List(context.Background())
	if err != nil {
		t.Fatalf("Subnets.List returned error %s\n", err)
	}

	if numSubnets := len(subnets); numSubnets < 1 {
		t.Errorf("Subnets.List got=%d\nwant=%d\n", numSubnets, 1)
	}

	err = client.Networks.Delete(context.Background(), network.UUID)
	if err != nil {
		t.Fatalf("Networks.Delete returned error %s\n", err)
	}
}

func TestIntegrationSubnet_CRUD(t *testing.T) {
	integrationTest(t)

	autoCreateSubnet := false;
	createNetworkRequest := &cloudscale.NetworkCreateRequest{
		Name:                 networkBaseName,
		AutoCreateIPV4Subnet: &autoCreateSubnet,
	}
	network, err := client.Networks.Create(context.TODO(), createNetworkRequest)
	if err != nil {
		t.Fatalf("Networks.Create returned error %s\n", err)
	}

	createSubnetRequest := &cloudscale.SubnetCreateRequest{
		CIDR: "192.168.192.0/22",
		GatewayAddress: "192.168.192.2",
		DNSServers: []string{"77.109.128.2", "213.144.129.20"},
		Network: network.UUID,
	}
	expected, err := client.Subnets.Create(context.TODO(), createSubnetRequest)
	if err != nil {
		t.Fatalf("Subnets.Create returned error %s\n", err)
	}

	subnet, err := client.Subnets.Get(context.Background(), expected.UUID)
	if err != nil {
		t.Fatalf("Subnets.Get returned error %s\n", err)
	}

	if !reflect.DeepEqual(subnet, expected) {
		t.Errorf("Error = %#v, expected %#v", subnet, expected)
	}

	subnets, err := client.Subnets.List(context.Background())
	if err != nil {
		t.Fatalf("Subnets.List returned error %s\n", err)
	}

	if numSubnets := len(subnets); numSubnets < 1 {
		t.Errorf("Subnets.List \n got=%d\nwant=%d", numSubnets, 1)
	}

	err = client.Subnets.Delete(context.Background(), expected.UUID)
	if err != nil {
		t.Fatalf("Subnets.Delete returned error %s\n", err)
	}
	err = client.Networks.Delete(context.Background(), network.UUID)
	if err != nil {
		t.Fatalf("Networks.Delete returned error %s\n", err)
	}
}

func TestIntegrationSubnet_Update(t *testing.T) {
	integrationTest(t)

	autoCreateSubnet := false;
	createNetworkRequest := &cloudscale.NetworkCreateRequest{
		Name:                 networkBaseName,
		AutoCreateIPV4Subnet: &autoCreateSubnet,
	}
	network, err := client.Networks.Create(context.TODO(), createNetworkRequest)
	if err != nil {
		t.Fatalf("Networks.Create returned error %s\n", err)
	}

	createSubnetRequest := &cloudscale.SubnetCreateRequest{
		Network: network.UUID,
		CIDR:  "10.0.0.0/8",
	}

	subnet, err := client.Subnets.Create(context.TODO(), createSubnetRequest)
	if err != nil {
		t.Fatalf("Subnets.Create returned error %s\n", err)
	}

	// update gateway
	expectedGateway := "10.255.255.254"
	updateRequest := &cloudscale.SubnetUpdateRequest{
		GatewayAddress: expectedGateway,
	}

	err = client.Subnets.Update(context.Background(), subnet.UUID, updateRequest)
	if err != nil {
		t.Fatalf("Subnets.Update returned error %s\n", err)
	}

	updatedSubnet, err := client.Subnets.Get(context.Background(), subnet.UUID)
	if err != nil {
		t.Fatalf("Subnets.Get returned error %s\n", err)
	}

	if actualGateway := updatedSubnet.GatewayAddress; actualGateway != expectedGateway {
		t.Errorf("Subnet MTU\ngot=%#v\nwant=%#v", updatedSubnet.GatewayAddress, expectedGateway)
	}

	// update dNSServers
	expectedDNSServers := []string{"77.109.128.2", "213.144.129.20", "1.1.1.1"}
	updateRequest = &cloudscale.SubnetUpdateRequest{
		DNSServers: expectedDNSServers,
	}

	err = client.Subnets.Update(context.Background(), subnet.UUID, updateRequest)
	if err != nil {
		t.Fatalf("Subnets.Update returned error %s\n", err)
	}

	updatedSubnet, err = client.Subnets.Get(context.Background(), subnet.UUID)
	if err != nil {
		t.Fatalf("Subnets.Get returned error %s\n", err)
	}

	if actualDNSServers := updatedSubnet.DNSServers; !reflect.DeepEqual(actualDNSServers, expectedDNSServers) {
		t.Errorf("Subnet MTU\ngot=%#v\nwant=%#v", updatedSubnet.DNSServers, expectedDNSServers)
	}

	err = client.Subnets.Delete(context.Background(), subnet.UUID)
	if err != nil {
		t.Fatalf("Networks.Delete returned error %s\n", err)
	}
	err = client.Networks.Delete(context.Background(), network.UUID)
	if err != nil {
		t.Fatalf("Networks.Delete returned error %s\n", err)
	}
}
