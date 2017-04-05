package stingray

// VirtualServer The data representation structure of a Virtual Server within stingray.
type VirtualServer struct {
	Properties struct {
		Basic struct {
			AddClusterIP             bool     `json:"add_cluster_ip"`
			AddXForwardedFor         bool     `json:"add_x_forwarded_for"`
			AddXForwardedProto       bool     `json:"add_x_forwarded_proto"`
			AutodetectUpgradeHeaders bool     `json:"autodetect_upgrade_headers"`
			BandwidthClass           string   `json:"bandwidth_class"`
			CloseWithRst             bool     `json:"close_with_rst"`
			Completionrules          []string `json:"completionrules"`
			ConnectTimeout           int      `json:"connect_timeout"`
			Enabled                  bool     `json:"enabled"`
			FtpForceServerSecure     bool     `json:"ftp_force_server_secure"`
			GlbServices              []string `json:"glb_services"`
			ListenOnAny              bool     `json:"listen_on_any"`
			ListenOnHosts            []string `json:"listen_on_hosts"`
			ListenOnTrafficIps       []string `json:"listen_on_traffic_ips"`
			Note                     string   `json:"note"`
			Pool                     string   `json:"pool"`
			Port                     int      `json:"port"`
			ProtectionClass          string   `json:"protection_class"`
			Protocol                 string   `json:"protocol"`
			RequestRules             []string `json:"request_rules"`
			ResponseRules            []string `json:"response_rules"`
			SlmClass                 string   `json:"slm_class"`
			SoNagle                  bool     `json:"so_nagle"`
			SslClientCertHeaders     string   `json:"ssl_client_cert_headers"`
			SslDecrypt               bool     `json:"ssl_decrypt"`
			SslHonorFallbackScsv     string   `json:"ssl_honor_fallback_scsv"`
			Transparent              bool     `json:"transparent"`
		} `json:"basic"`
		Aptimizer struct {
			Enabled bool          `json:"enabled"`
			Profile []interface{} `json:"profile"`
		} `json:"aptimizer"`
		Connection struct {
			Keepalive              bool   `json:"keepalive"`
			KeepaliveTimeout       int    `json:"keepalive_timeout"`
			MaxClientBuffer        int    `json:"max_client_buffer"`
			MaxServerBuffer        int    `json:"max_server_buffer"`
			MaxTransactionDuration int    `json:"max_transaction_duration"`
			ServerFirstBanner      string `json:"server_first_banner"`
			Timeout                int    `json:"timeout"`
		} `json:"connection"`
		ConnectionErrors struct {
			ErrorFile string `json:"error_file"`
		} `json:"connection_errors"`
		Cookie struct {
			Domain      string `json:"domain"`
			NewDomain   string `json:"new_domain"`
			PathRegex   string `json:"path_regex"`
			PathReplace string `json:"path_replace"`
			Secure      string `json:"secure"`
		} `json:"cookie"`
		DNS struct {
			EdnsUdpsize int      `json:"edns_udpsize"`
			MaxUdpsize  int      `json:"max_udpsize"`
			RrsetOrder  string   `json:"rrset_order"`
			Verbose     bool     `json:"verbose"`
			Zones       []string `json:"zones"`
		} `json:"dns"`
		Ftp struct {
			DataSourcePort    int  `json:"data_source_port"`
			ForceClientSecure bool `json:"force_client_secure"`
			PortRangeHigh     int  `json:"port_range_high"`
			PortRangeLow      int  `json:"port_range_low"`
			SslData           bool `json:"ssl_data"`
		} `json:"ftp"`
		Gzip struct {
			CompressLevel int      `json:"compress_level"`
			Enabled       bool     `json:"enabled"`
			EtagRewrite   string   `json:"etag_rewrite"`
			IncludeMime   []string `json:"include_mime"`
			MaxSize       int      `json:"max_size"`
			MinSize       int      `json:"min_size"`
			NoSize        bool     `json:"no_size"`
		} `json:"gzip"`
		HTTP struct {
			ChunkOverheadForwarding string `json:"chunk_overhead_forwarding"`
			LocationRegex           string `json:"location_regex"`
			LocationReplace         string `json:"location_replace"`
			LocationRewrite         string `json:"location_rewrite"`
			MimeDefault             string `json:"mime_default"`
			MimeDetect              bool   `json:"mime_detect"`
		} `json:"http"`
		KerberosProtocolTransition struct {
			Enabled   bool   `json:"enabled"`
			Principal string `json:"principal"`
			Target    string `json:"target"`
		} `json:"kerberos_protocol_transition"`
		Log struct {
			ClientConnectionFailures  bool   `json:"client_connection_failures"`
			Enabled                   bool   `json:"enabled"`
			Filename                  string `json:"filename"`
			Format                    string `json:"format"`
			SaveAll                   bool   `json:"save_all"`
			ServerConnectionFailures  bool   `json:"server_connection_failures"`
			SessionPersistenceVerbose bool   `json:"session_persistence_verbose"`
			SslFailures               bool   `json:"ssl_failures"`
		} `json:"log"`
		RecentConnections struct {
			Enabled bool `json:"enabled"`
			SaveAll bool `json:"save_all"`
		} `json:"recent_connections"`
		RequestTracing struct {
			Enabled bool `json:"enabled"`
			TraceIo bool `json:"trace_io"`
		} `json:"request_tracing"`
		Rtsp struct {
			StreamingPortRangeHigh int `json:"streaming_port_range_high"`
			StreamingPortRangeLow  int `json:"streaming_port_range_low"`
			StreamingTimeout       int `json:"streaming_timeout"`
		} `json:"rtsp"`
		Sip struct {
			DangerousRequests      string `json:"dangerous_requests"`
			FollowRoute            bool   `json:"follow_route"`
			MaxConnectionMem       int    `json:"max_connection_mem"`
			Mode                   string `json:"mode"`
			RewriteURI             bool   `json:"rewrite_uri"`
			StreamingPortRangeHigh int    `json:"streaming_port_range_high"`
			StreamingPortRangeLow  int    `json:"streaming_port_range_low"`
			StreamingTimeout       int    `json:"streaming_timeout"`
			TimeoutMessages        bool   `json:"timeout_messages"`
			TransactionTimeout     int    `json:"transaction_timeout"`
		} `json:"sip"`
		SMTP struct {
			ExpectStarttls bool `json:"expect_starttls"`
		} `json:"smtp"`
		Ssl struct {
			AddHTTPHeaders         bool     `json:"add_http_headers"`
			ClientCertCas          []string `json:"client_cert_cas"`
			EllipticCurves         []string `json:"elliptic_curves"`
			IssuedCertsNeverExpire []string `json:"issued_certs_never_expire"`
			OcspEnable             bool     `json:"ocsp_enable"`
			OcspIssuers            []struct {
				Issuer        string `json:"issuer"`
				Aia           bool   `json:"aia"`
				Nounce        string `json:"nounce"`
				Required      string `json:"required"`
				ResponderCert string `json:"responder_cert"`
				Signer        string `json:"signer"`
				URL           string `json:"url"`
			} `json:"ocsp_issuers"`
			OcspMaxResponseAge    int    `json:"ocsp_max_response_age"`
			OcspStapling          bool   `json:"ocsp_stapling"`
			OcspTimeTolerance     int    `json:"ocsp_time_tolerance"`
			OcspTimeout           int    `json:"ocsp_timeout"`
			PreferSslv3           bool   `json:"prefer_sslv3"`
			RequestClientCert     string `json:"request_client_cert"`
			SendCloseAlerts       bool   `json:"send_close_alerts"`
			ServerCertDefault     string `json:"server_cert_default"`
			ServerCertHostMapping []struct {
				Host        string `json:"host"`
				Certificate bool   `json:"certificate"`
			} `json:"server_cert_host_mapping"`
			SignatureAlgorithms string `json:"signature_algorithms"`
			SslCiphers          string `json:"ssl_ciphers"`
			SslSupportSsl2      string `json:"ssl_support_ssl2"`
			SslSupportSsl3      string `json:"ssl_support_ssl3"`
			SslSupportTLS1      string `json:"ssl_support_tls1"`
			SslSupportTLS11     string `json:"ssl_support_tls1_1"`
			SslSupportTLS12     string `json:"ssl_support_tls1_2"`
			TrustMagic          bool   `json:"trust_magic"`
		} `json:"ssl"`
		Syslog struct {
			Enabled     bool   `json:"enabled"`
			Format      string `json:"format"`
			IPEndPoint  string `json:"ip_end_point"`
			MsgLenLimit int    `json:"msg_len_limit"`
		} `json:"syslog"`
		TCP struct {
			ProxyClose bool `json:"proxy_close"`
		} `json:"tcp"`
		UDP struct {
			EndPointPersistence       bool `json:"end_point_persistence"`
			PortSmp                   bool `json:"port_smp"`
			ResponseDatagramsExpected int  `json:"response_datagrams_expected"`
			Timeout                   int  `json:"timeout"`
		} `json:"udp"`
		WebCache struct {
			ControlOut    string `json:"control_out"`
			Enabled       bool   `json:"enabled"`
			ErrorPageTime int    `json:"error_page_time"`
			MaxTime       int    `json:"max_time"`
			RefreshTime   int    `json:"refresh_time"`
		} `json:"web_cache"`
	} `json:"properties"`
}
