package nice

type NiceStream struct {
	name								string
	id 									uint
	n_components						uint
	initial_binding_request_received	bool
	components							[]*NiceComponent
	conncheck_list						[]*CandidateCheckPair
	
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
	stream.
}