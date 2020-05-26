/**
 * @Author: fanpengfei
 * @Description:
 * @File:  pubsub_timer
 * @Version: 1.0.0
 * @Date: 2020/5/26 8:28
 */

package main

import "time"

func (c *PubSub) initPing() {
	c.ping = make(chan struct{}, 1)
	go func() {
		timer := time.NewTimer(pingTimeout)
		timer.Stop()

		healthy := true
		for {
			timer.Reset(pingTimeout)
			select {
			case <-c.ping:
				healthy = true
				if !timer.Stop() {
					<-timer.C
				}
			case <-timer.C:
				pingErr := c.Ping()
				if healthy {
					healthy = false
				} else {
					if pingErr == nil {
						pingErr = errPingTimeout
					}
					c.mu.Lock()
					c.reconnect(pingErr)
					healthy = true
					c.mu.Unlock()
				}
			case <-c.exit:
				return
			}
		}
	}()
}

func main() {

}
