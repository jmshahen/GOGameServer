package main
 
import (
        "bufio";
        "fmt";
        "net";
        "io";
        "os";
)
 
 
var nl byte = 10;
 
func echo(conn *net.TCPConn) {
        addr := conn.RemoteAddr();
        rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn));
        for {
                s, err := rw.ReadString(nl);
                if len(s) > 0 {
                        fmt.Printf("conn %s said %d %s\n", addr, len(s), s);
                        rw.WriteString(s);
                        rw.Flush();
                } else if err == io.EOF {
                        fmt.Printf("conn %s eof\n", addr);
                        conn.Close();
                        return;
                } else {
                        fmt.Printf("error reading: %s\n", err);
                        conn.Close();
                        return;
                }
                if s == "quit\r\n" {
                        conn.Close();
                        return;
                }
        }
}
 
func main() {
        l, err := net.ListenTCP("tcp4", &net.TCPAddr{net.IPv4zero, 9998});
        if l == nil {
                fmt.Printf("cannot listen: %s\n", err);
                os.Exit(1);
        }
        fmt.Printf("listening at %s\n", l.Addr());
        for {
                conn, err := l.AcceptTCP();
                if conn == nil {
                        fmt.Printf("accept error: %s\n", err);
                        l.Close();
                        os.Exit(1);
                }
                fmt.Printf("connection from %s\n", conn.RemoteAddr());
                go echo(conn);
        }
}