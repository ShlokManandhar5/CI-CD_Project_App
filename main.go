package main

import (
        "fmt"
        "net"
        "net/http"
        "strings"
)

func getClientIP(r *http.Request) string {
        // Check X-Forwarded-For header
        ip := r.Header.Get("X-Forwarded-For")
        if ip != "" {
                ips := strings.Split(ip, ",")
                return strings.TrimSpace(ips[0])
        }

        // Check X-Real-IP header
        ip = r.Header.Get("X-Real-IP")
        if ip != "" {
                return ip
        }

        // Fallback to RemoteAddr
        host, _, err := net.SplitHostPort(r.RemoteAddr)
        if err != nil {
                return r.RemoteAddr
        }

        return host
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
        clientIP := getClientIP(r)

        fmt.Fprintf(w, "Hello, World! I am Shlok Mndr\n")
        fmt.Fprintf(w, "Your IP Address: %s\n", clientIP)
}

func main() {
        http.HandleFunc("/", helloHandler)

        fmt.Println("Server running on port 8080")
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                fmt.Println("Error:", err)
        }
}
