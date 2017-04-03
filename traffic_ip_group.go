package main

// TrafficIPGroup The data representation structure of a Traffic IP Group within stingray.
type TrafficIPGroup struct {
	Properties struct {
		Basic struct {
			Enabled        bool `json:"enabled"`
			HashSourcePort bool `json:"hash_source_port"`
			IPMapping      []struct {
				IP             string `json:"ip"`
				TrafficManager bool   `json:"traffic_manager"`
			} `json:"ip_mapping"`
			Ipaddresses                  []string `json:"ipaddresses"`
			Keeptogether                 bool     `json:"keeptogether"`
			Location                     int      `json:"location"`
			Machines                     []string `json:"machines"`
			Mode                         string   `json:"mode"`
			Multicast                    string   `json:"multicast"`
			Note                         string   `json:"note"`
			RhiBgpMetricBase             int      `json:"rhi_bgp_metric_base"`
			RhiBgpPassiveMetricOffset    int      `json:"rhi_bgp_passive_metric_offset"`
			RhiOspfv2MetricBase          int      `json:"rhi_ospfv2_metric_base"`
			RhiOspfv2PassiveMetricOffset int      `json:"rhi_ospfv2_passive_metric_offset"`
			RhiProtocols                 string   `json:"rhi_protocols"`
			Slaves                       []string `json:"slaves"`
		} `json:"basic"`
	} `json:"properties"`
}
