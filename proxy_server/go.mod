module github.com/dklassen/ready_racer/proxy_server

go 1.17

// +heroku goVersion 1.17

require (
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)

replace github.com/dklassen/racer_ready/proxy_server/pkg => ../pkg