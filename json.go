package main

type GobgpNeighbor struct {
	Add_paths struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"add-paths"`
	Afi_safis []struct {
		Add_paths struct {
			Config struct{} `json:"config"`
			State  struct{} `json:"state"`
		} `json:"add-paths"`
		Apply_policy struct {
			Config struct{} `json:"config"`
			State  struct{} `json:"state"`
		} `json:"apply-policy"`
		Config struct {
			Afi_safi_name string `json:"afi-safi-name"`
			Enabled       bool   `json:"enabled"`
		} `json:"config"`
		Ipv4_labelled_unicast struct {
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
		} `json:"ipv4-labelled-unicast"`
		Ipv4_unicast struct {
			Config       struct{} `json:"config"`
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
			State struct{} `json:"state"`
		} `json:"ipv4-unicast"`
		Ipv6_labelled_unicast struct {
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
		} `json:"ipv6-labelled-unicast"`
		Ipv6_unicast struct {
			Config       struct{} `json:"config"`
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
			State struct{} `json:"state"`
		} `json:"ipv6-unicast"`
		L2vpn_evpn struct {
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
		} `json:"l2vpn-evpn"`
		L2vpn_vpls struct {
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
		} `json:"l2vpn-vpls"`
		L3vpn_ipv4_multicast struct {
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
		} `json:"l3vpn-ipv4-multicast"`
		L3vpn_ipv4_unicast struct {
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
		} `json:"l3vpn-ipv4-unicast"`
		L3vpn_ipv6_multicast struct {
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
		} `json:"l3vpn-ipv6-multicast"`
		L3vpn_ipv6_unicast struct {
			Prefix_limit struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"prefix-limit"`
		} `json:"l3vpn-ipv6-unicast"`
		Long_lived_graceful_restart struct {
			Config struct{} `json:"config"`
			State  struct{} `json:"state"`
		} `json:"long-lived-graceful-restart"`
		Mp_graceful_restart struct {
			Config struct{} `json:"config"`
			State  struct{} `json:"state"`
		} `json:"mp-graceful-restart"`
		Prefix_limit struct {
			Config struct{} `json:"config"`
			State  struct{} `json:"state"`
		} `json:"prefix-limit"`
		Route_selection_options struct {
			Config struct{} `json:"config"`
			State  struct{} `json:"state"`
		} `json:"route-selection-options"`
		Route_target_membership struct {
			Config struct{} `json:"config"`
			State  struct{} `json:"state"`
		} `json:"route-target-membership"`
		State struct {
			Family int64 `json:"family"`
		} `json:"state"`
		Use_multiple_paths struct {
			Config struct{} `json:"config"`
			Ebgp   struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"ebgp"`
			Ibgp struct {
				Config struct{} `json:"config"`
				State  struct{} `json:"state"`
			} `json:"ibgp"`
			State struct{} `json:"state"`
		} `json:"use-multiple-paths"`
	} `json:"afi-safis"`
	Apply_policy struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"apply-policy"`
	As_path_options struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"as-path-options"`
	Config struct {
		Local_as         int64  `json:"local-as"`
		Neighbor_address string `json:"neighbor-address"`
		Peer_as          int64  `json:"peer-as"`
	} `json:"config"`
	Ebgp_multihop struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"ebgp-multihop"`
	Error_handling struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"error-handling"`
	Graceful_restart struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"graceful-restart"`
	Logging_options struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"logging-options"`
	Route_reflector struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"route-reflector"`
	Route_server struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"route-server"`
	State struct {
		Adj_table struct {
			Accepted int64 `json:"accepted"`
			Received int64 `json:"received"`
		} `json:"adj-table"`
		Admin_state           string `json:"admin-state"`
		Local_capability_list []struct {
			Code  int64 `json:"code"`
			Value int64 `json:"value"`
		} `json:"local-capability-list"`
		Messages struct {
			Received struct {
				Keepalive int64 `json:"keepalive"`
				Open      int64 `json:"open"`
				Total     int64 `json:"total"`
				Update    int64 `json:"update"`
			} `json:"received"`
			Sent struct {
				Keepalive int64 `json:"keepalive"`
				Open      int64 `json:"open"`
				Total     int64 `json:"total"`
			} `json:"sent"`
		} `json:"messages"`
		Neighbor_address       string   `json:"neighbor-address"`
		Peer_as                int64    `json:"peer-as"`
		Peer_type              string   `json:"peer-type"`
		Queues                 struct{} `json:"queues"`
		Remote_capability_list []struct {
			Code  int64 `json:"code"`
			Value int64 `json:"value"`
		} `json:"remote-capability-list"`
		Remote_router_id string `json:"remote-router-id"`
		Session_state    string `json:"session-state"`
	} `json:"state"`
	Timers struct {
		Config struct {
			Connect_retry      int64 `json:"connect-retry"`
			Hold_time          int64 `json:"hold-time"`
			Keepalive_interval int64 `json:"keepalive-interval"`
		} `json:"config"`
		State struct {
			Downtime             int64 `json:"downtime"`
			Keepalive_interval   int64 `json:"keepalive-interval"`
			Negotiated_hold_time int64 `json:"negotiated-hold-time"`
			Uptime               int64 `json:"uptime"`
		} `json:"state"`
	} `json:"timers"`
	Transport struct {
		Config struct {
			Local_address string `json:"local-address"`
		} `json:"config"`
		State struct{} `json:"state"`
	} `json:"transport"`
	TTL_security struct {
		Config struct{} `json:"config"`
		State  struct{} `json:"state"`
	} `json:"ttl-security"`
	Use_multiple_paths struct {
		Config struct{} `json:"config"`
		Ebgp   struct {
			Config struct{} `json:"config"`
			State  struct{} `json:"state"`
		} `json:"ebgp"`
		Ibgp struct {
			Config struct{} `json:"config"`
			State  struct{} `json:"state"`
		} `json:"ibgp"`
		State struct{} `json:"state"`
	} `json:"use-multiple-paths"`
}
