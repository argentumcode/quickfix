package quickfix

import "fmt"

type EventSeverity int

const (
	EventSeverityINFO EventSeverity = iota
	EventSeverityWARNING
	EventSeverityERROR
)

func (s EventSeverity) String() string {
	switch s {
	case EventSeverityINFO:
		return "INFO"
	case EventSeverityWARNING:
		return "WARNING"
	case EventSeverityERROR:
		return "ERROR"
	}
	return fmt.Sprintf("InvalidSeverity(%d)", s)
}

// Log is a generic interface for logging FIX messages and events.
type Log interface {
	//OnIncoming log incoming fix message
	OnIncoming([]byte)

	//OnOutgoing log outgoing fix message
	OnOutgoing([]byte)

	//OnEvent log fix event
	OnEvent(EventSeverity, string)

	//OnEventf log fix event according to format specifier
	OnEventf(EventSeverity, string, ...interface{})
}

// The LogFactory interface creates global and session specific Log instances
type LogFactory interface {
	//Create global log
	Create() (Log, error)

	//CreateSessionLog session specific log
	CreateSessionLog(sessionID SessionID) (Log, error)
}
