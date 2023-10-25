package main

import (
	"fmt"
	"sync"
)

type RemoteControl struct {
	powerOn  bool
	channel  int
	volume   int
	mute     bool
	commands []Command
}

type Command interface {
	Execute()
}

type PowerOnCommand struct {
	remote *RemoteControl
}

func (cmd *PowerOnCommand) Execute() {
	cmd.remote.powerOn = true
	fmt.Println("Turning on")
}

type ChannelChangeCommand struct {
	remote     *RemoteControl
	newChannel int
}

func (cmd *ChannelChangeCommand) Execute() {
	cmd.remote.channel = cmd.newChannel
	fmt.Printf("Changing channel to %d\n", cmd.newChannel)
}

type VolumeAdjustCommand struct {
	remote    *RemoteControl
	newVolume int
}

func (cmd *VolumeAdjustCommand) Execute() {
	cmd.remote.volume = cmd.newVolume
	fmt.Printf("Adjusting volume to %d\n", cmd.newVolume)
}

var instance *RemoteControl
var once sync.Once

func GetRemoteControlInstance() *RemoteControl {
	once.Do(func() {
		instance = &RemoteControl{}
	})
	return instance
}

func main() {
	remote := GetRemoteControlInstance()

	powerOnCmd := &PowerOnCommand{remote}
	channelChangeCmd := &ChannelChangeCommand{remote, 5}
	volumeAdjustCmd := &VolumeAdjustCommand{remote, 25}

	powerOnCmd.Execute()
	channelChangeCmd.Execute()
	volumeAdjustCmd.Execute()

	fmt.Printf("TV is on: %v\n", remote.powerOn)
	fmt.Printf("Current channel: %d\n", remote.channel)
	fmt.Printf("Current volume: %d\n", remote.volume)
}
