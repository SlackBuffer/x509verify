
```bash
# 生成 root ca 私钥
openssl genrsa -out rootCA.key 4096
# root ca 自签名得到 root ca 根证书
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024 -out rootCA.crt

# 生成服务端私钥
openssl genrsa -out mydomain.com.key 2048
# 生成 csr
openssl req -new -sha256 -key mydomain.com.key -subj "/C=US/ST=CA/O=MyOrg, Inc./CN=mydomain.com" -out mydomain.com.csr
## 校验 csr 内容
openssl req -in mydomain.com.csr -noout -text
# 是 root ca 私钥签名 csr 生成服务端证书
openssl x509 -req -in mydomain.com.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out mydomain.com.crt -days 1024 -sha256
## 校验证书内容
openssl x509 -in mydomain.com.crt -text -noout
```

https://gist.github.com/fntlnz/cf14feb5a46b2eda428e000157447309

https://pkg.go.dev/crypto/x509#example-Certificate.Verify

证书存成字符串需顶格，不能有 tab

---

opts.Roots 不为 nil 就用 opts.Roots 用作根证书来校验

fabric 
	// we only support a single validation chain;
	// if there's more than one then there might
	// be unclarity about who owns the identity