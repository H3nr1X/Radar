package main

import (
	"log"
	"github.com/H3nr1X/ReadWriteMemory"
)

const (
	dwEntityList 	= 0x4D4F1FC
	m_bDormant 		= 0xED
	dwLocalPlayer 	= 0xD3AC5C
	m_bSpotted = 0x93D
)
func radar() {
	process, err := ReadWriteMemory.ProcessByName("csgo")
	if err != nil {
		log.Panicf("csgo.exe not found. Error: %s", err.Error())
	}
	client := process.Modules["client.dll"].ModBaseAddr
	for {
		for i := 0; i < 64; i++ {
			entitySize := 0x10
			Player, _ := process.ReadIntPtr(client + dwEntityList + (uintptr(i) * uintptr(entitySize)))
			isDormant, _ := process.ReadIntPtr(Player + m_bDormant)
			if isDormant == 1 {continue}
			process.WriteInt(Player + m_bSpotted, 1)
	
		}
	}
}
func main() {
	radar()
}