package ev3

import (
	"fmt"
	"log"
	"time"

	"github.com/dop251/goja"
	"github.com/ev3go/ev3dev"
)

var ev3Ports = map[int64]string{
	0: "ev3-ports:outA",
	1: "ev3-ports:outB",
	2: "ev3-ports:outC",
	3: "ev3-ports:outD",
}

var EV3Object = map[string]interface{}{
	"MoveLargeMotor": MoveLargeMotor,
}

func MoveLargeMotor(call goja.FunctionCall) goja.Value {
	port := call.Argument(0)
	speed := call.Argument(1)
	seconds := call.Argument(2)
	goPort := port.ToInteger()
	speedGo := speed.ToInteger()
	secondsGo := seconds.ToInteger()

	goRealPort := ev3Ports[goPort]

	motor, err := ev3dev.TachoMotorFor(goRealPort, "lego-ev3-l-motor")
	if err != nil {
		panic(fmt.Sprintf("Failed to get motor: %+v", err))
	}
	err = motor.SetStopAction("brake").Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to set stop action: %+v", err))
	}

	motor.SetSpeedSetpoint(int(speedGo) * motor.MaxSpeed() / 100).Command("run-forever")
	time.Sleep(time.Duration(secondsGo) * time.Second)
	motor.Command("stop")
	checkErrors(motor)

	return goja.Undefined()
}

func checkErrors(devs ...ev3dev.Device) {
	for _, d := range devs {
		err := d.(*ev3dev.TachoMotor).Err()
		if err != nil {
			drv, dErr := ev3dev.DriverFor(d)
			if dErr != nil {
				drv = fmt.Sprintf("(missing driver name: %v)", dErr)
			}
			addr, aErr := ev3dev.AddressOf(d)
			if aErr != nil {
				drv = fmt.Sprintf("(missing port address: %v)", aErr)
			}
			log.Fatalf("motor error for %s:%s on port %s: %v", d, drv, addr, err)
		}
	}
}
