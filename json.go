package main

type GobgpNeighbor struct {
	ApplyPolicy struct {
		ExportPolicy struct {
			Direction     int `json:"direction"`
			DefaultAction int `json:"default_action"`
		} `json:"export_policy"`
		ImportPolicy struct {
			Direction     int `json:"direction"`
			DefaultAction int `json:"default_action"`
		} `json:"import_policy"`
	} `json:"apply_policy"`
	Conf struct {
		LocalAs         int    `json:"local_as"`
		NeighborAddress string `json:"neighbor_address"`
		PeerAs          int    `json:"peer_as"`
		PeerType        int    `json:"peer_type"`
	} `json:"conf"`
	EbgpMultihop struct {
		Enabled     bool `json:"enabled"`
		MultihopTTL int  `json:"multihop_ttl"`
	} `json:"ebgp_multihop"`
	RouteReflector struct {
	} `json:"route_reflector"`
	State struct {
		Messages struct {
			Received struct {
				Update         int `json:"update"`
				Open           int `json:"open"`
				Keepalive      int `json:"keepalive"`
				Total          int `json:"total"`
				WithdrawUpdate int `json:"withdraw_update"`
				WithdrawPrefix int `json:"withdraw_prefix"`
			} `json:"received"`
			Sent struct {
				Open      int `json:"open"`
				Keepalive int `json:"keepalive"`
				Total     int `json:"total"`
			} `json:"sent"`
		} `json:"messages"`
		NeighborAddress string `json:"neighbor_address"`
		PeerAs          int    `json:"peer_as"`
		PeerType        int    `json:"peer_type"`
		Queues          struct {
		} `json:"queues"`
		SessionState int `json:"session_state"`
		RemoteCap    []struct {
			TypeURL string `json:"type_url"`
			Value   string `json:"value,omitempty"`
		} `json:"remote_cap"`
		LocalCap []struct {
			TypeURL string `json:"type_url"`
			Value   string `json:"value,omitempty"`
		} `json:"local_cap"`
		RouterID string `json:"router_id"`
	} `json:"state"`
	Timers struct {
		Config struct {
			ConnectRetry           int `json:"connect_retry"`
			HoldTime               int `json:"hold_time"`
			KeepaliveInterval      int `json:"keepalive_interval"`
			IdleHoldTimeAfterReset int `json:"idle_hold_time_after_reset"`
		} `json:"config"`
		State struct {
			KeepaliveInterval  int `json:"keepalive_interval"`
			NegotiatedHoldTime int `json:"negotiated_hold_time"`
			Uptime             struct {
				Seconds int `json:"seconds"`
			} `json:"uptime"`
			Downtime struct {
				Seconds int `json:"seconds"`
			} `json:"downtime"`
		} `json:"state"`
	} `json:"timers"`
	Transport struct {
		LocalAddress string `json:"local_address"`
	} `json:"transport"`
	RouteServer struct {
	} `json:"route_server"`
	GracefulRestart struct {
	} `json:"graceful_restart"`
	AfiSafis []struct {
		MpGracefulRestart struct {
			Config struct {
			} `json:"config"`
		} `json:"mp_graceful_restart"`
		Config struct {
			Family struct {
				Afi  int `json:"afi"`
				Safi int `json:"safi"`
			} `json:"family"`
			Enabled bool `json:"enabled"`
		} `json:"config"`
		State struct {
			Family struct {
				Afi  int `json:"afi"`
				Safi int `json:"safi"`
			} `json:"family"`
			Enabled  bool `json:"enabled"`
			Received int  `json:"received"`
			Accepted int  `json:"accepted"`
		} `json:"state"`
		ApplyPolicy struct {
			ExportPolicy struct {
				Direction     int `json:"direction"`
				DefaultAction int `json:"default_action"`
			} `json:"export_policy"`
			ImportPolicy struct {
				Direction     int `json:"direction"`
				DefaultAction int `json:"default_action"`
			} `json:"import_policy"`
		} `json:"apply_policy"`
		RouteSelectionOptions struct {
			Config struct {
			} `json:"config"`
		} `json:"route_selection_options"`
		UseMultiplePaths struct {
			Config struct {
			} `json:"config"`
			Ebgp struct {
				Config struct {
				} `json:"config"`
			} `json:"ebgp"`
			Ibgp struct {
				Config struct {
				} `json:"config"`
			} `json:"ibgp"`
		} `json:"use_multiple_paths"`
		RouteTargetMembership struct {
			Config struct {
			} `json:"config"`
		} `json:"route_target_membership"`
		LongLivedGracefulRestart struct {
			Config struct {
			} `json:"config"`
		} `json:"long_lived_graceful_restart"`
		AddPaths struct {
			Config struct {
			} `json:"config"`
		} `json:"add_paths"`
	} `json:"afi_safis"`
}
