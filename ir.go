package main

import (
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/chbmuc/lirc"
)

func makeHandler(c *MQTT.Client, cmd string) func(event lirc.Event) {
	return func(event lirc.Event) {
		c.Publish("/lights/set", 0, false, cmd)
	}
}

func main() {
	/*
	 *  MQTT setup
	 */

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("golirc")

	mc := MQTT.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	/*
	 *  LIRC setup
	 */

	ir, err := lirc.Init("/var/run/lirc/lircd")
	if err != nil {
		panic(err)
	}

	ir.Handle("", "BTN_0", makeHandler(mc, "0"))
	ir.Handle("", "BTN_1", makeHandler(mc, "1"))
	ir.Handle("", "BTN_2", makeHandler(mc, "2"))
	ir.Handle("", "BTN_3", makeHandler(mc, "3"))

	go ir.Run()

	select {}
}
