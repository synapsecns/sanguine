package types

const (
	agentStatusOffsetFlag   = 0
	agentStatusOffsetDomain = 8
	agentStatusOffsetIndex  = 40
	agentStatusSize         = 48
)

// AgentStatus is the agent status interface.
type AgentStatus interface {
	// Flag is the current status flag of the agent.
	Flag() uint8
	// Domain that agent is assigned to.
	Domain() uint32
	// Index of the agent in list of agents.
	Index() uint32
}

type agentStatus struct {
	flag   uint8
	domain uint32
	index  uint32
}

// NewAgentStatus creates a new agent status.
func NewAgentStatus(flag uint8, domain, index uint32) AgentStatus {
	return &agentStatus{
		flag:   flag,
		domain: domain,
		index:  index,
	}
}

func (s agentStatus) Flag() uint8 {
	return s.flag
}

func (s agentStatus) Domain() uint32 {
	return s.domain
}

func (s agentStatus) Index() uint32 {
	return s.index
}

var _ AgentStatus = agentStatus{}
