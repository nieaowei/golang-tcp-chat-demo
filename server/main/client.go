/*******************************************************
 *  File        :   client.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/15 12:59 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"golang-tcp-chat-demo/proto"
	"net"
	"strconv"
)

type Client struct {
	conn net.Conn   // Client-side connection.
	buf  [8192]byte // Client-side data caching.
}

// Read package from connection and read data.
func (p *Client) readPackage() (msg proto.Message, err error) {
	// Read the first four bytes of the packet.
	lenth, err := p.conn.Read(p.buf[0:4])
	if lenth != 4 {
		err = errors.New("read header failed.")
		return
	}

	lenth, err = strconv.Atoi(string(p.buf[0:4]))
	if err != nil {
		err = errors.New("read package lenth failed.")
		return
	}
	// Read entire data according to the lenth.
	lenth1, err := p.conn.Read(p.buf[0:lenth])
	if lenth1 != lenth || err != nil {
		err = errors.New("read body failed")
		return
	}

	err = json.Unmarshal(p.buf[0:lenth], &msg)
	if err != nil {
		fmt.Println("data unmarshal failed.")
	}

	return
}

func (p *Client) writePackage(data []byte) (err error) {
	lenth := uint32(len(data))                    //Get data lenth.
	binary.BigEndian.PutUint32(p.buf[0:4], lenth) // convert 'lenth' to bigendian.
	lenthW, err := p.conn.Write(p.buf[0:4])
	if err != nil {
		fmt.Println("write header failed.")
		return
	}
	lenthW, err = p.conn.Write(data)
	if err != nil {
		fmt.Println("write data failed.")
		return
	}
	if lenthW != int(lenth) {
		fmt.Println("write data not finish.")
		err = errors.New("write data finish.")
		return
	}
	return
}

func (p *Client) Process() (err error) {
	for {
		var msg proto.Message
		msg, err = p.readPackage()
		if err != nil {
			return
		}
		err = p.processMsg(msg)
		if err != nil {
			return
		}
	}
}

func (p *Client) processMsg(msg proto.Message) (err error) {
	switch msg.Cmd {
	case proto.UserLogin:
		err = p.login(msg)
	case proto.UserRegister:
		err = p.register(msg)
	default:
		err = errors.New("unsupport message.")
	}
	return
}

// Process the results after logging in.
func (p *Client) loginResp(err error) {
	// Customize the results of the protocol.
	respMsg := proto.Message{
		Cmd: proto.UserLoginRes,
	}
	// Results after login.
	loginRes := proto.LoginCmdRes{
		Code: 200,
	}
	if err != nil {
		loginRes.Code = 500
		loginRes.Error = fmt.Sprintf("%v", err)
	}
	// Marshal the results aftering login.
	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("marshal failed.")
		return
	}
	// Convert from results aftering login to the results of the protocol.
	respMsg.Data = string(data)
	data, err = json.Marshal(respMsg)
	if err != nil {
		fmt.Println("marshal failed.")
		return
	}
	// Send data.
	err = p.writePackage(data)
	if err != nil {
		fmt.Println("send failed.", err)
		return
	}
}

func (p *Client) login(msg proto.Message) (err error) {
	defer func() {
		p.loginResp(err)
	}()

	fmt.Printf("recv user login request, data:%v", msg)
	var cmd proto.LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	_, err = mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return
	}
	return
}

func (p *Client) register(msg proto.Message) (err error) {
	var cmd proto.RegisterCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}

	err = mgr.Register(&cmd.User)
	if err != nil {
		return
	}
	return
}
