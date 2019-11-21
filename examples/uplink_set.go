package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		ClientOV   *ov.OVClient
		new_uplink = "test_new"
		upd_uplink = "test_update"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1000,
		"")

	fmt.Println("#................... Get-all Uplink-Sets ...............#")
	sort := "name:desc"
	uplinkset_list, err := ovc.GetUplinkSets("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Uplink-Set List .................#")
		for i := 0; i < len(uplinkset_list.Members); i++ {
			fmt.Println(uplinkset_list.Members[i].Name)
		}
	}

	// Create Uplink Set
	networkUris := new([]utils.Nstring)
	*networkUris = append(*networkUris, utils.NewNstring("/rest/ethernet-networks/33fabe08-d91e-45ba-8b10-5a19d00e3b1d"))

	fcNetworkUris := make([]utils.Nstring, 0)
	fcoeNetworkUris := make([]utils.Nstring, 0)
	portConfigInfos := make([]ov.PortConfigInfos, 0)
	privateVlanDomains := make([]ov.PrivateVlanDomains, 0)

	uplinkSet := ov.UplinkSet{Name: new_uplink,
		LogicalInterconnectURI:         utils.NewNstring("/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d"),
		NetworkURIs:                    *networkUris,
		FcNetworkURIs:                  fcNetworkUris,
		FcoeNetworkURIs:                fcoeNetworkUris,
		PortConfigInfos:                portConfigInfos,
		ConnectionMode:                 "Auto",
		NetworkType:                    "Ethernet",
		EthernetNetworkType:            "Tagged",
		Type:                           "uplink-setV5",
		ManualLoginRedistributionState: "NotSupported",
		PrivateVlanDomains:             privateVlanDomains}

	err = ovc.CreateUplinkSet(uplinkSet)
	if err != nil {
		fmt.Println("............... UplinkSet Creation Failed:", err)
	} else {
		fmt.Println(".... Uplink Set Created Successfully")
	}

	fmt.Println("#................... Get Uplink-Set by Name & Update it ...............#")
	new_uplinkset, _ := ovc.GetUplinkSetByName(new_uplink)
	new_uplinkset.Name = upd_uplink

	err = ovc.UpdateUplinkSet(new_uplinkset)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#.................... Uplink-Set Update success ...........#")
	}

	err = ovc.DeleteUplinkSet(new_uplinkset.Name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#.................... Deleted Uplink Set Successfully .....#")
	}
}
