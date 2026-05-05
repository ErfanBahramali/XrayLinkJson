# DESCRIPTION
This program takes input from the command line and checks whether the input is a JSON or a link.
If the input is a JSON, it converts it to Share links, and if the input is a link, it converts it to JSON.
This is done using functions from the libXray library, and the result is encoded in Base64.

# INSTALL
```bash
go install .
```

If `go install` succeeds but the command is not found, add Go's bin directory to your `PATH`:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

To keep it after restart:

**For Bash:**

```bash
echo 'export PATH="$(go env GOPATH)/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

**For Zsh:**

```bash
echo 'export PATH="$(go env GOPATH)/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

Run directly without installing:

```bash
go run . "vless://123456789@example.com:443?security=tls&sni=sni.example.com&type=ws&host=host.example.com&path=%2F#sample-vless"
```

# USAGE
```bash
xrayLinkJson "vless://123456789@example.com:443?security=tls&sni=sni.example.com&type=ws&host=host.example.com&path=%2F#sample-vless"
```

Output:
```
{"success":true,"data":{"transport":null,"log":null,"routing":null,"dns":null,"inbounds":null,"outbounds":[{"protocol":"vless","sendThrough":"sample-vless","tag":"","settings":{"vnext":[{"address":"example.com","port":443,"users":[{"id":"123456789","encryption":"none"}]}]},"streamSettings":{"address":null,"port":0,"network":"ws","security":"tls","tlsSettings":{"allowInsecure":false,"certificates":null,"serverName":"sni.example.com","alpn":null,"enableSessionResumption":false,"disableSystemRoot":false,"minVersion":"","maxVersion":"","cipherSuites":"","fingerprint":"","rejectUnknownSni":false,"pinnedPeerCertificateChainSha256":null,"pinnedPeerCertificatePublicKeySha256":null,"curvePreferences":null,"masterKeyLog":"","serverNameToVerify":""},"realitySettings":null,"rawSettings":null,"tcpSettings":null,"xhttpSettings":null,"splithttpSettings":null,"kcpSettings":null,"grpcSettings":null,"wsSettings":{"host":"host.example.com","path":"/","headers":null,"acceptProxyProtocol":false,"heartbeatPeriod":0},"httpupgradeSettings":null,"sockopt":null},"proxySettings":null,"mux":null}],"policy":null,"api":null,"metrics":null,"stats":null,"reverse":null,"fakeDns":null,"observatory":null,"burstObservatory":null}}
```

```bash
xrayLinkJson "dm1lc3M6Ly9leGFtcGxlLmNvbTo0NDM="
```

Output:
```
{"success":true,"data":{"transport":null,"log":null,"routing":null,"dns":null,"inbounds":null,"outbounds":[{"protocol":"vmess","sendThrough":"","tag":"","settings":{"vnext":[{"address":"example.com","port":443,"users":[{"id":"","security":"","experiments":""}]}]},"streamSettings":null,"proxySettings":null,"mux":null}],"policy":null,"api":null,"metrics":null,"stats":null,"reverse":null,"fakeDns":null,"observatory":null,"burstObservatory":null}}
```

```bash
xrayLinkJson "{\"outbounds\":[{\"protocol\":\"vless\",\"settings\":{\"vnext\":[{\"address\":\"example.com\",\"port\":443}]}}]}"
```

Output:
```
{"success":true,"data":"vless://example.com:443"}
```

# License
MIT © 2026 [ErfanBahramali](https://github.com/ErfanBahramali)
