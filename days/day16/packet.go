package day16

import (
	"fmt"
	"strconv"
)

type Packet struct {
	version    int
	pType      int
	value      int
	subPackets []Packet
}

func ParseBits(bits []uint64, verbose bool) (*Packet, int) {
	p := new(Packet)

	var (
		cursor int
		length int
	)

	cursor = 0

	// version
	length = 3
	version, _ := strconv.ParseInt(ToBitString(bits[cursor:cursor+length]), 2, 64)
	p.version = int(version)
	cursor += length
	if verbose {
		fmt.Println("version: ", version)
	}

	// type
	length = 3
	pType, _ := strconv.ParseInt(ToBitString(bits[cursor:cursor+length]), 2, 64)
	p.pType = int(pType)
	cursor += length

	switch p.pType {
	case 4:
		length = 5
		valueBits := make([]uint64, 0)
		for {
			if verbose {
				fmt.Println(bits)
			}

			newBits := bits[cursor : cursor+length]
			valueBits = append(valueBits, newBits[1:]...)
			cursor += length
			if newBits[0] == 0 {
				break
			}
		}
		value, _ := strconv.ParseInt(ToBitString(valueBits), 2, 64)
		p.value = int(value)

		if verbose {
			fmt.Println("type 4: ", p.value)
		}

		return p, cursor
		break
	default:
		if verbose {
			fmt.Println("type ", p.pType)
		}

		p.value = -1
		lengthTypeId := bits[cursor : cursor+1]
		cursor += 1

		switch lengthTypeId[0] {
		case 1:
			length = 11
			subPacketsCount, _ := strconv.ParseInt(ToBitString(bits[cursor:cursor+length]), 2, 64)
			cursor += length

			var seen int64
			seen = 0

			var (
				newPacket *Packet
				l         int
			)

			if verbose {
				fmt.Println("should see count of ", subPacketsCount)
			}

			for seen < subPacketsCount {
				if verbose {
					fmt.Println("Subpacket:")
				}

				newPacket, l = ParseBits(bits[cursor:], verbose)
				p.subPackets = append(p.subPackets, *newPacket)

				cursor += l
				seen += 1
			}
		case 0:
			length = 15
			subPacketsLength, _ := strconv.ParseInt(ToBitString(bits[cursor:cursor+length]), 2, 64)
			cursor += length

			var seen int64
			seen = 0

			var (
				newPacket *Packet
				l         int
			)

			if verbose {
				fmt.Println("should see length of ", subPacketsLength)
			}

			for seen < subPacketsLength {
				if verbose {
					fmt.Println("Subpacket:")
				}

				newPacket, l = ParseBits(bits[cursor:], verbose)
				p.subPackets = append(p.subPackets, *newPacket)

				cursor += l
				seen += int64(l)
			}
		}
	}

	return p, cursor
}

func (p *Packet) Solve(verbose bool) int {
	switch p.pType {
	case 0:
		sum := 0

		for _, packet := range p.subPackets {
			sum += packet.Solve(verbose)
		}

		return sum
	case 1:
		product := 1

		for _, packet := range p.subPackets {
			product *= packet.Solve(verbose)
		}

		return product
	case 2:
		minimum := 1000000000000000000

		for _, packet := range p.subPackets {
			value := packet.Solve(verbose)
			if value < minimum {
				minimum = value
			}
		}

		return minimum
	case 3:
		maximum := 0

		for _, packet := range p.subPackets {
			value := packet.Solve(verbose)
			if value > maximum {
				maximum = value
			}
		}

		return maximum
	case 4:
		return p.value
	case 5:
		if p.subPackets[0].Solve(verbose) > p.subPackets[1].Solve(verbose) {
			return 1
		}
		return 0
	case 6:
		if p.subPackets[0].Solve(verbose) < p.subPackets[1].Solve(verbose) {
			return 1
		}
		return 0
	case 7:
		if p.subPackets[0].Solve(verbose) == p.subPackets[1].Solve(verbose) {
			return 1
		}
		return 0
	}

	return 0
}

func (p *Packet) SumVersions() int {
	sum := 0

	sum += p.version

	for _, subPacket := range p.subPackets {
		sum += subPacket.SumVersions()
	}

	return sum
}

func ToBitString(input []uint64) string {
	output := ""

	for i := 0; i < len(input); i++ {
		digit := strconv.Itoa(int(input[i]))
		output += digit
	}

	return output
}
