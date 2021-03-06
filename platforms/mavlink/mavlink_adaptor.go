package mavlink

import (
	"io"

	"github.com/hybridgroup/gobot"
	"github.com/tarm/goserial"
)

type MavlinkAdaptor struct {
	gobot.Adaptor
	sp      io.ReadWriteCloser
	connect func(*MavlinkAdaptor)
}

// NewMavLinkAdaptor creates a new mavlink adaptor with specified name and port
func NewMavlinkAdaptor(name string, port string) *MavlinkAdaptor {
	return &MavlinkAdaptor{
		Adaptor: *gobot.NewAdaptor(
			name,
			"mavlink.MavlinkAdaptor",
			port,
		),
		connect: func(m *MavlinkAdaptor) {
			s, err := serial.OpenPort(&serial.Config{Name: m.Port(), Baud: 57600})
			if err != nil {
				panic(err)
			}
			m.sp = s
		},
	}
}

// Connect returns true if connection to device is successful
func (m *MavlinkAdaptor) Connect() bool {
	m.connect(m)
	return true
}

// Finalize returns true if connection to devices is closed successfully
func (m *MavlinkAdaptor) Finalize() bool {
	m.sp.Close()
	return true
}
