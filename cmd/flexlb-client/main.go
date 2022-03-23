package main

import (
	"flag"
	"flexlb/client"
	"fmt"

	"gopkg.in/yaml.v2"
)

func main() {
	var (
		host     string
		ca       string
		cert     string
		key      string
		insecure bool
		status   bool
		create   bool
		modify   string
		config   string
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

	flag.BoolVar(&create, "create", false, "Create instance -config <config_file>")
	flag.BoolVar(&create, "c", false, "Create instance -config <config_file>")

	flag.StringVar(&modify, "modify", "", "Modify instance <name>")
	flag.StringVar(&modify, "m", "", "Modify instance <name>")

	flag.StringVar(&config, "config", "", "Instance config file (create, modify)")
	flag.StringVar(&config, "f", "", "Instance config file (create, modify)")

	flag.Parse()

	lb, err := client.NewTLSClient(host, ca, cert, key, insecure, nil)

	if err != nil {
		panic(err)
	}

	if status {
		if resp, err := client.GetReadyStatus(lb); err != nil {
			panic(err)
		} else {
			printAsYaml(resp)
		}
	}

	if list {
		if resp, err := client.ListInstances(lb, &name); err != nil {
			panic(err)
		} else {
			printAsYaml(resp)
		}
	}

	if get != "" {
		if resp, err := client.GetInstance(lb, get); err != nil {
			panic(err)
		} else {
			printAsYaml(resp)
		}
	}

	if resume != "" {
		if resp, err := client.StartInstance(lb, resume); err != nil {
			panic(err)
		} else {
			printAsYaml(resp)
		}
	}

	if pause != "" {
		if resp, err := client.StopInstance(lb, pause); err != nil {
			panic(err)
		} else {
			printAsYaml(resp)
		}
	}

	if delete != "" {
		if err := client.DeleteInstance(lb, delete); err != nil {
			panic(err)
		} else {
			fmt.Println("ok")
		}
	}

	if create {
		if resp, err := client.CreateInstance(lb, config); err != nil {
			panic(err)
		} else {
			printAsYaml(resp)
		}
	}

	if modify != "" {
		if resp, err := client.ModifyInstance(lb, modify, config); err != nil {
			panic(err)
		} else {
			printAsYaml(resp)
		}
	}

}

func printAsYaml(obj interface{}) {
	if raw, err := yaml.Marshal(obj); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(raw))
	}
}
