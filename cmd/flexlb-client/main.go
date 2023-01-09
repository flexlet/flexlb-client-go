package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/flexlet/flexlb-client-go/client"
)

func main() {
	var (
		host     string
		ca       string
		cert     string
		key      string
		insecure bool
		status   bool
		create   string
		modify   string
		list     bool
		name     string
		get      string
		resume   string
		pause    string
		delete   string
	)
	flag.StringVar(&host, "host", "localhost:8443", "FlexLB API server")
	flag.StringVar(&ca, "ca", "/etc/flexlb/certs/ca.crt", "TLS CA certificate")
	flag.StringVar(&cert, "cert", "/etc/flexlb/certs/client.crt", "TLS client certificate")
	flag.StringVar(&key, "key", "/etc/flexlb/certs/client.key", "TLS client key")
	flag.BoolVar(&insecure, "insecure", true, "TLS insecure skip verify")

	flag.BoolVar(&status, "status", false, "Show ready status")
	flag.BoolVar(&status, "s", false, "Show ready status")

	flag.BoolVar(&list, "list", false, "List instance")
	flag.BoolVar(&list, "l", false, "List instance")

	flag.StringVar(&name, "name", "", "Instance name fuzzy filter")
	flag.StringVar(&name, "n", "", "Instance name fuzzy filter")

	flag.StringVar(&get, "get", "", "Get instance <name>")
	flag.StringVar(&get, "g", "", "Get instance <name>")

	flag.StringVar(&resume, "resume", "", "Resume instance <name>")
	flag.StringVar(&resume, "r", "", "Resume instance <name>")

	flag.StringVar(&pause, "pause", "", "Pause instance <name>")
	flag.StringVar(&pause, "p", "", "Pause instance <name>")

	flag.StringVar(&delete, "delete", "", "Stop instance <name>")
	flag.StringVar(&delete, "d", "", "Stop instance <name>")

	flag.StringVar(&create, "create", "", "Create instance <config_file>")
	flag.StringVar(&create, "c", "", "Create instance <config_file>")

	flag.StringVar(&modify, "modify", "", "Modify instance <config_file>")
	flag.StringVar(&modify, "m", "", "Modify instance <config_file>")

	flag.Parse()

	lb, err := client.NewTLSClient(host, ca, cert, key, insecure, nil)

	if err != nil {
		panic(err)
	}

	if status {
		if resp, err := lb.GetReadyStatus(); err != nil {
			panic(err)
		} else {
			printAsJson(resp)
		}
	}

	if list {
		if resp, err := lb.ListInstances(&name); err != nil {
			panic(err)
		} else {
			printAsJson(resp)
		}
	}

	if get != "" {
		if resp, err := lb.GetInstance(get); err != nil {
			panic(err)
		} else {
			printAsJson(resp)
		}
	}

	if resume != "" {
		if resp, err := lb.StartInstance(resume); err != nil {
			panic(err)
		} else {
			printAsJson(resp)
		}
	}

	if pause != "" {
		if resp, err := lb.StopInstance(pause); err != nil {
			panic(err)
		} else {
			printAsJson(resp)
		}
	}

	if delete != "" {
		if err := lb.DeleteInstance(delete); err != nil {
			panic(err)
		} else {
			fmt.Println("ok")
		}
	}

	if create != "" {
		if cfg, err1 := client.LoadConfig(create); err1 != nil {
			panic(err1)
		} else {
			if resp, err2 := lb.CreateInstance(cfg); err2 != nil {
				panic(err2)
			} else {
				printAsJson(resp)
			}
		}
	}

	if modify != "" {
		if cfg, err1 := client.LoadConfig(modify); err1 != nil {
			panic(err1)
		} else {
			if resp, err2 := lb.ModifyInstance(cfg); err2 != nil {
				panic(err2)
			} else {
				printAsJson(resp)
			}
		}
	}

}

func printAsJson(obj interface{}) {
	if raw, err := json.MarshalIndent(obj, "", "    "); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(raw))
	}
}
