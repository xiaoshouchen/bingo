package tests

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestConn(t *testing.T) {
	open, err := sql.Open("mysql", "hbb:password@tcp(139.224.216.133:3306)/sq_hbb")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	exec, err := open.Conn(context.TODO())
	exec.
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(m)
	err = open.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

//https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::HandshakeResponse
//auth payload content

//4              capability flags, CLIENT_PROTOCOL_41 always set
//4              max-packet size
//1              character set
//string[23]     reserved (all [0])
//string[NUL]    username
//  if capabilities & CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA {
//lenenc-int     length of auth-response
//string[n]      auth-response
//  } else if capabilities & CLIENT_SECURE_CONNECTION {
//1              length of auth-response
//string[n]      auth-response
//  } else {
//string[NUL]    auth-response
//  }
//  if capabilities & CLIENT_CONNECT_WITH_DB {
//string[NUL]    database
//  }
//  if capabilities & CLIENT_PLUGIN_AUTH {
//string[NUL]    auth plugin name
//  }
//  if capabilities & CLIENT_CONNECT_ATTRS {
//lenenc-int     length of all key-values
//lenenc-str     key
//lenenc-str     value
//   if-more data in 'length of all key-values', more keys and value pairs
//  }
