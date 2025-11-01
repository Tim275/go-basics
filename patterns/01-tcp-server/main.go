// TCP Server Pattern - Low-Level Network Programming
// Kapitel 2: Going from TCP to HTTP

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// 1️⃣ TCP Listener auf Port 8080 öffnen
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("TCP Server listening on :8080")

	// 2️⃣ Accept Loop - Wartet auf neue Connections
	for {
		// Accept() blockiert hier bis ein Client verbindet
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// 3️⃣ Neuer Goroutine pro Client = Concurrent Handling
		go handleConnection(conn)
	}
}

// handleConnection verarbeitet einen einzelnen Client
func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	log.Printf("New connection from %s", clientAddr)

	// 4️⃣ Lese Daten vom Client
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		log.Printf("Received from %s: %s", clientAddr, message)

		// 5️⃣ Echo Response zurück senden
		response := fmt.Sprintf("Echo: %s\n", strings.ToUpper(message))
		conn.Write([]byte(response))

		// Verbindung beenden bei "quit"
		if strings.ToLower(message) == "quit" {
			log.Printf("Client %s disconnected", clientAddr)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from %s: %v", clientAddr, err)
	}
}

/*
📚 TCP Konzepte:

1. Connection-oriented:
   - 3-Way Handshake (SYN, SYN-ACK, ACK)
   - Client und Server bauen dedizierte Verbindung auf

2. Reliable:
   - Pakete kommen garantiert an
   - Automatic Retransmits bei Verlust
   - Sequence Numbers für Ordnung

3. Bidirectional (Full-Duplex):
   - Beide Seiten können gleichzeitig senden/empfangen

4. Stream-based:
   - Kontinuierlicher Byte-Stream
   - Keine Message-Boundaries

🔧 Testen:
   # Terminal 1 - Server starten:
   go run main.go

   # Terminal 2 - Client verbinden:
   telnet localhost 8080
   > hello
   > quit

🌐 OSI Model - TCP ist Layer 4 (Transport):
   ┌──────────────────┐
   │ Layer 7          │ ← Application: HTTP, FTP, SMTP
   │ Application      │
   ├──────────────────┤
   │ Layer 6          │ ← Presentation: SSL/TLS, Encryption
   │ Presentation     │
   ├──────────────────┤
   │ Layer 5          │ ← Session: Connection Management
   │ Session          │
   ├──────────────────┤
   │ Layer 4  ⭐      │ ← Transport: TCP, UDP
   │ Transport        │   (WIR SIND HIER!)
   ├──────────────────┤
   │ Layer 3          │ ← Network: IP Routing
   │ Network          │
   ├──────────────────┤
   │ Layer 2          │ ← Data Link: Ethernet, WiFi
   │ Data Link        │
   ├──────────────────┤
   │ Layer 1          │ ← Physical: Cables, Signals
   │ Physical         │
   └──────────────────┘
*/
