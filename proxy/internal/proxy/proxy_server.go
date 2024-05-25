package proxy

// var services = map[string]string{"Host: localhost:8080": "localhost:8081", "Host: test": "localhost:8081"}
// var services = map[string]string{"Host: localhost:8080": "google.com:80", "Host: test": "localhost:8081"}
//var services = map[string]string{"Host: localhost": "localhost:8081", "Host: test": "localhost:8081"}

//var services = map[string]string{"Host: localhost:8080": "google.com:80", "Host: test": "localhost:8081"}

//var proxyTargets = target.GetMockedProxyTargets()

//func Start() {
//	proxy_port := utils.GetProxyServerPort()
//	log.Info().Msg(fmt.Sprintf("Starting proxy proxy on port: %s", proxy_port))
//	listenOnTcpSocket(proxy_port)
//}
//
//func listenOnTcpSocket(port string) {
//	cer, err := tls.LoadX509KeyPair("proxy.crt", "proxy.key")
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	config := &tls.Config{Certificates: []tls.Certificate{cer}}
//	ln, err := tls.Listen("tcp", fmt.Sprintf(":%s", port), config)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer ln.Close()
//
//	//ln, err := net.Listen("tcp", ":"+port)
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//
//	for {
//		conn, err := ln.Accept()
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//
//		go handleConnection(conn)
//	}
//}
//
//func handleConnection(conn net.Conn) {
//	defer conn.Close()
//	//conn.SetReadDeadline(time.Now().Add(2 * time.Second))
//
//	request, err := readContentFromBuffer(conn)
//	if err != nil {
//		log.Error().Msg("Error reading request content")
//		return
//	}
//
//	req, proxyAddress := parseHttpRequest(request)
//	//req, proxyAddress := request, ""
//
//	response, err := proxy(req, proxyAddress)
//	if err != nil {
//		log.Error().Msg(fmt.Sprintf("Failed proxying to destination %s", proxyAddress))
//		return
//	}
//
//	bytesWritten, err := writeContentToBuffer(response, proxyAddress, conn)
//	if err != nil {
//		log.Error().Msg(fmt.Sprintf("Error writing response from destination %s", proxyAddress))
//		return
//	}
//
//	log.Info().Msg(fmt.Sprintf("Successfuly wrote %d bytes from response", bytesWritten))
//	return
//}
//
//func proxy(content string, proxyAddress string) (string, error) {
//	conn, err := net.Dial("tcp", proxyAddress)
//
//	if err != nil {
//		log.Error().Msg(fmt.Sprintf("Error establishing TCP connection with %s", proxyAddress))
//		return "", err
//	}
//
//	defer conn.Close()
//
//	_, err = writeContentToBuffer(content, proxyAddress, conn)
//	if err != nil {
//		log.Error().Msg(fmt.Sprintf("Proxying to destination %s failed, cannot write request content to buffer", proxyAddress))
//		return "", err
//	}
//
//	response, err := readContentFromBuffer(conn)
//	if err != nil {
//		log.Error().Msg(fmt.Sprintf("Error reading response from destination %s, reading from buffer failed", proxyAddress))
//		return "", err
//	}
//
//	return response, nil
//}

//func writeContentToBuffer(content string, proxyAddress string, conn net.Conn) (int, error) {
//	bytesWritten, err := conn.Write([]byte(content))
//	if err != nil {
//		log.Error().Msg(fmt.Sprintf("Error writing content to the buffer for destination %s", proxyAddress))
//		return 0, err
//	}
//
//	log.Info().Msg(fmt.Sprintf("Wrote %d bytes to the buffer, for destination %s", bytesWritten, proxyAddress))
//	return bytesWritten, nil
//}
//
//func readContentFromBuffer(conn net.Conn) (string, error) {
//	request := ""
//
//	bytesRead := 1024
//	for bytesRead == 1024 {
//		fmt.Println("BYTES READ: ", bytesRead)
//		content, read, err := readChunkFromBuffer(conn)
//		if err != nil {
//			log.Error().Msg("Cannot read content from buffer due to chunk reading error")
//			return "", err
//		}
//		request += content
//		bytesRead = read
//	}
//
//	return request, nil
//}
//
//func readChunkFromBuffer(conn net.Conn) (string, int, error) {
//	buf := make([]byte, 1024)
//	bytesRead, err := conn.Read(buf)
//	if err != nil {
//		log.Error().Msg("Error reading content chunk from buffer")
//		return "", 0, err
//	}
//	return string(buf), bytesRead, nil
//}

//func parseHttpRequest(request string) (string, string) {
//	res := appendHeaders(strings.Split(request, "\r\n"))
//	go request2.ParseRequest(request)
//
//	for i := 0; i < len(res); i++ {
//		line := res[i]
//		if strings.HasPrefix(line, "Host: ") {
//			hostSplit := strings.Split(line, ":")
//			host := strings.Join(hostSplit[:2], ":")
//			res[i] = fmt.Sprintf("Host: %s", services[host])
//			return strings.Join(res, "\r\n") + "\r\n", services[host]
//		}
//	}
//
//	return "", ""
//}
//
//func appendHeaders(request []string) []string {
//	if len(request) > 1 {
//		res := append(request[:2], request[1:]...)
//		res[1] = "X-TEST: ABCD"
//		return res
//	}
//
//	return append(request, "X-TEST: ABCD")
//}
