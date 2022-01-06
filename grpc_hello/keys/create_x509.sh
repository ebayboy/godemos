
# 1. 创建私钥
openssl ecparam -genkey -name secp384r1 -out server.key

# 2. 生成自签名证书(Common Name是grpc client链接需要用的)
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

#Country Name (2 letter code) [AU]:CN
#State or Province Name (full name) [Some-State]:beijing
#Locality Name (eg, city) []:beijing
#Organization Name (eg, company) [Internet Widgits Pty Ltd]:JD
#Organizational Unit Name (eg, section) []:dev
#Common Name (e.g. server FQDN or YOUR name) []:go-grpc-example
#Email Address []:ebayboy@163.com


# 3. 打因证书信息
openssl x509 -in cert.pem -noout -text


