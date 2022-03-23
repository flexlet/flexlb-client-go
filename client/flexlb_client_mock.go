package client

import (
	"encoding/json"
	"flexlb/client/instance"
	"flexlb/client/service"
	"flexlb/models"
	"io/ioutil"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func NewTLSClient(host string, ca string, cert string, key string, insecure bool, formats strfmt.Registry) (*Flexlb, error) {
	tlsOpts := httptransport.TLSClientOptions{
		CA:                 ca,
		Key:                key,
		Certificate:        cert,
		InsecureSkipVerify: insecure,
	}
	client, err := httptransport.TLSClient(tlsOpts)
	if err != nil {
		return nil, err
	}
	transport := httptransport.NewWithClient(host, DefaultBasePath, DefaultSchemes, client)
	return New(transport, formats), nil
}

func GetReadyStatus(lb *Flexlb) (*models.ReadyStatus, error) {
	params := service.NewReadyzParams()
	if resp, err := lb.Service.Readyz(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}

func ListInstances(lb *Flexlb, name *string) ([]*models.Instance, error) {
	params := instance.NewListParams()
	params.Name = name
	if resp, err := lb.Instance.List(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}

func GetInstance(lb *Flexlb, name string) (*models.Instance, error) {
	params := instance.NewGetParams()
	params.Name = name
	if resp, err := lb.Instance.Get(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}

func StartInstance(lb *Flexlb, name string) (*models.Instance, error) {
	params := instance.NewStartParams()
	params.Name = name
	if resp, err := lb.Instance.Start(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}

func StopInstance(lb *Flexlb, name string) (*models.Instance, error) {
	params := instance.NewStopParams()
	params.Name = name
	if resp, err := lb.Instance.Stop(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}

func DeleteInstance(lb *Flexlb, name string) error {
	params := instance.NewDeleteParams()
	params.Name = name
	if _, err := lb.Instance.Delete(params); err != nil {
		return err
	} else {
		return nil
	}
}

func CreateInstance(lb *Flexlb, file string) (*models.Instance, error) {
	params := instance.NewCreateParams()
	if raw, err := ioutil.ReadFile(file); err != nil {
		return nil, err
	} else {
		var cfg models.InstanceConfig
		if err := json.Unmarshal(raw, &cfg); err != nil {
			return nil, err
		} else {
			params.Config = &cfg
		}

	}
	if resp, err := lb.Instance.Create(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}

func ModifyInstance(lb *Flexlb, name string, file string) (*models.Instance, error) {
	params := instance.NewModifyParams()
	params.Name = name
	if raw, err := ioutil.ReadFile(file); err != nil {
		return nil, err
	} else {
		var cfg models.InstanceConfig
		if err := json.Unmarshal(raw, &cfg); err != nil {
			return nil, err
		} else {
			params.Config = &cfg
		}
	}
	if resp, err := lb.Instance.Modify(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}
