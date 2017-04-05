package stingray

// Pool The data representation structure of a Pool within stingray.
type Pool struct {
	Properties struct {
		Basic struct {
			BandwidthClass                string   `json:"bandwidth_class"`
			FailurePool                   string   `json:"failure_pool"`
			MaxConnectionAttempts         int      `json:"max_connection_attempts"`
			MaxIdleConnectionsPernode     int      `json:"max_idle_connections_pernode"`
			MaxTimedOutConnectionAttempts int      `json:"max_timed_out_connection_attempts"`
			Monitors                      []string `json:"monitors"`
			NodeCloseWithRst              bool     `json:"node_close_with_rst"`
			NodeConnectionAttempts        int      `json:"node_connection_attempts"`
			NodeDeleteBehavior            string   `json:"node_delete_behavior"`
			NodeDrainToDeleteTimeout      int      `json:"node_drain_to_delete_timeout"`
			NodesTable                    []struct {
				Node     string `json:"node"`
				Priority int    `json:"priority"`
				State    string `json:"state"`
				Weight   int    `json:"weight"`
				SourceIP string `json:"source_ip"`
			} `json:"nodes_table"`
			Note              string `json:"note"`
			PassiveMonitoring bool   `json:"passive_monitoring"`
			PersistenceClass  string `json:"persistence_class"`
			Transparent       bool   `json:"transparent"`
		} `json:"basic"`
		AutoScaling struct {
			AddnodeDelaytime int      `json:"addnode_delaytime"`
			CloudCredentials string   `json:"cloud_credentials"`
			Cluster          string   `json:"cluster"`
			DataCenter       string   `json:"data_center"`
			DataStore        string   `json:"data_store"`
			Enabled          bool     `json:"enabled"`
			External         bool     `json:"external"`
			Hysteresis       int      `json:"hysteresis"`
			Imageid          string   `json:"imageid"`
			IpsToUse         string   `json:"ips_to_use"`
			LastNodeIdleTime int      `json:"last_node_idle_time"`
			MaxNodes         int      `json:"max_nodes"`
			MinNodes         int      `json:"min_nodes"`
			Name             string   `json:"name"`
			Port             int      `json:"port"`
			Refractory       int      `json:"refractory"`
			ResponseTime     int      `json:"response_time"`
			ScaleDownLevel   int      `json:"scale_down_level"`
			ScaleUpLevel     int      `json:"scale_up_level"`
			Securitygroupids []string `json:"securitygroupids"`
			SizeID           string   `json:"size_id"`
			Subnetids        []string `json:"subnetids"`
		} `json:"auto_scaling"`
		Connection struct {
			MaxConnectTime        int `json:"max_connect_time"`
			MaxConnectionsPerNode int `json:"max_connections_per_node"`
			MaxQueueSize          int `json:"max_queue_size"`
			MaxReplyTime          int `json:"max_reply_time"`
			QueueTimeout          int `json:"queue_timeout"`
		} `json:"connection"`
		DNSAutoscale struct {
			Enabled   bool     `json:"enabled"`
			Hostnames []string `json:"hostnames"`
			Port      int      `json:"port"`
		} `json:"dns_autoscale"`
		Ftp struct {
			SupportRfc2428 bool `json:"support_rfc_2428"`
		} `json:"ftp"`
		HTTP struct {
			Keepalive              bool `json:"keepalive"`
			KeepaliveNonIdempotent bool `json:"keepalive_non_idempotent"`
		} `json:"http"`
		KerberosProtocolTransition struct {
			Principal string `json:"principal"`
			Target    string `json:"target"`
		} `json:"kerberos_protocol_transition"`
		LoadBalancing struct {
			Algorithm       string `json:"algorithm"`
			PriorityEnabled bool   `json:"priority_enabled"`
			PriorityNodes   int    `json:"priority_nodes"`
		} `json:"load_balancing"`
		Node struct {
			CloseOnDeath  bool `json:"close_on_death"`
			RetryFailTime int  `json:"retry_fail_time"`
		} `json:"node"`
		SMTP struct {
			SendStarttls bool `json:"send_starttls"`
		} `json:"smtp"`
		Ssl struct {
			ClientAuth          bool     `json:"client_auth"`
			CommonNameMatch     []string `json:"common_name_match"`
			EllipticCurves      []string `json:"elliptic_curves"`
			Enable              bool     `json:"enable"`
			Enhance             bool     `json:"enhance"`
			SendCloseAlerts     bool     `json:"send_close_alerts"`
			ServerName          bool     `json:"server_name"`
			SignatureAlgorithms string   `json:"signature_algorithms"`
			SslCiphers          string   `json:"ssl_ciphers"`
			SslSupportSsl2      string   `json:"ssl_support_ssl2"`
			SslSupportSsl3      string   `json:"ssl_support_ssl3"`
			SslSupportTLS1      string   `json:"ssl_support_tls1"`
			SslSupportTLS11     string   `json:"ssl_support_tls1_1"`
			SslSupportTLS12     string   `json:"ssl_support_tls1_2"`
			StrictVerify        bool     `json:"strict_verify"`
		} `json:"ssl"`
		TCP struct {
			Nagle bool `json:"nagle"`
		} `json:"tcp"`
		UDP struct {
			AcceptFrom     string `json:"accept_from"`
			AcceptFromMask string `json:"accept_from_mask"`
		} `json:"udp"`
	} `json:"properties"`
}
