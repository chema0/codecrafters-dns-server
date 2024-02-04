package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Header struct {
	ID uint16
	// QR      byte
	// OPCODE  byte
	// AA      byte
	// TC      byte
	// RD      byte
	// RA      byte
	// Z       byte
	// RCODE   byte
	FLAGS   uint16
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}

type DNSMessage struct {
	Header Header
}

func NewDNSMessage(data []byte) DNSMessage {
	dnsMessage := DNSMessage{}

	buf := bytes.NewReader(data)

	parseField(buf, &dnsMessage.Header.ID)
	parseField(buf, &dnsMessage.Header.FLAGS)
	parseField(buf, &dnsMessage.Header.QDCOUNT)
	parseField(buf, &dnsMessage.Header.ANCOUNT)
	parseField(buf, &dnsMessage.Header.NSCOUNT)
	parseField(buf, &dnsMessage.Header.ARCOUNT)

	return dnsMessage
}

func (m *DNSMessage) serialize() []byte {
	bytes := make([]byte, 12)

	binary.BigEndian.PutUint16(bytes[0:2], m.Header.ID)
	binary.BigEndian.PutUint16(bytes[2:4], m.Header.FLAGS)
	binary.BigEndian.PutUint16(bytes[4:6], m.Header.QDCOUNT)
	binary.BigEndian.PutUint16(bytes[6:8], m.Header.ANCOUNT)
	binary.BigEndian.PutUint16(bytes[8:10], m.Header.NSCOUNT)
	binary.BigEndian.PutUint16(bytes[10:12], m.Header.ARCOUNT)

	return bytes
}

func (m *DNSMessage) setFlags(qr, opcode, aa, tc, rd, ra, z, rcode uint16) {
	m.Header.FLAGS = qr<<15 | opcode<<11 | aa<<10 | tc<<9 | rd<<8 | ra<<7 | z<<4 | rcode
}

func parseField(buf *bytes.Reader, field any) {
	err := binary.Read(buf, binary.BigEndian, field)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
}
