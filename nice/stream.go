package nice

type NiceStream struct {
	name								string
	id 									uint
	n_components						uint
	initial_binding_request_received	bool
	components							[]*NiceComponent
	conncheck_list						[]*CandidateCheckPair
	gathering							bool
	gathering_started					bool
	peer_gathering_done					bool
	local_ufrag							string
	local_password						string
}

func NewNiceStream(stream_id uint, n_components uint, agent *NiceAgent) *NiceStream {
	s := &NiceStream{}
	s.id = stream_id
	s.n_components = n_components
	s.components = make([]*NiceComponent, 0)
	var i uint = 0
	for i = 0; i < n_components; i++ {
		s.components = append(s.components, NewNiceComponent(agent, s, i + 1))
	}
	s.peer_gathering_done = !agent.use_ice_trickle
	return s
}

func (this *NiceStream) find_component_by_id(component_id uint) *NiceComponent {
	for i := 0; i < len(this.components); i++ {
		if this.components[i].id == component_id {
			return this.components[i]
		}
	}
	return nil
}

func (this *NiceStream) nice_stream_initialize_credentials(rng *NiceRNG) {
	u := rng.nice_rng_generate_bytes_print(NICE_STREAM_DEF_UFRAG - 1)
	this.local_ufrag = string(u)
	p := rng.nice_rng_generate_bytes_print(NICE_STREAM_DEF_PWD - 1)
	this.local_password = string(p)
}