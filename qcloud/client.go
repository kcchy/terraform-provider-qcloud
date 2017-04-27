package qcloud

import (
	"encoding/json"
	"fmt"
	"github.com/QcloudApi/qcloud_sign_golang"
	"log"
	"os"
)

type CommonResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

// create tencent cloud load balancer

type Client struct {
	Region    string
	SecretId  string
	SecretKey string
}

type CreateLoadBalancerArgs struct {
	LoadBalancerType int
	Forward          int
	LoadBalancerName string
	DomainPrefix     string
	VpcId            string
	SubnetId         string
	ProjectId        int
}

type CreateLoadBalancerResponse struct {
	CommonResponse
	DealIds           []string          `json:"dealIds"`
	UnLoadBalancerIds UnLoadBalancerIds `json:"unLoadBalancerIds"`
}

type UnLoadBalancerIds map[string][]string

func (client *Client) CreateLoadBalancer(args *CreateLoadBalancerArgs) (response *CreateLoadBalancerResponse, err error) {
	response = &CreateLoadBalancerResponse{}

	config := map[string]interface{}{
		"secretId":  client.SecretId,
		"secretKey": client.SecretKey,
		"debug":     true,
	}

	params := map[string]interface{}{
		"Action":           "CreateLoadBalancer",
		"Region":           client.Region,
		"loadBalancerType": args.LoadBalancerType,
		"forward":          args.Forward,
		"loadBalancerName": args.LoadBalancerName,
		"domainPrefix":     args.DomainPrefix,
		"vpcId":            args.VpcId,
		"subnetId":         args.SubnetId,
		"projectId":        args.ProjectId,
	}

	clb, err := QcloudApi.SendRequest("lb", params, config)

	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(clb), response)
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Println(clb, client.SecretId, client.SecretKey, client.Region)

	return response, nil
}

// query tencent cloud load balancer by load balancer name

type DescribeLoadBalancerResponse struct {
	CommonResponse
	TotalCount      int                `json:"totalCount"`
	LoadBalancerSet []*LoadBalancerSet `json:"loadBalancerSet"`
}

type LoadBalancerSet struct {
	LoadBalancerId   string   `json:"loadBalancerId"`
	UnLoadBalancerId string   `json:"unLoadBalancerId"`
	LoadBalancerName string   `json:"loadBalancerName"`
	LoadBalancerType int      `json:"loadBalancerType"`
	Domain           string   `json:"domain"`
	LoadBalancerVips []string `json:"loadBalancerVips"`
	Status           int      `json:"status"`
	CreateTime       string   `json:"createTime"`
	StatusTime       string   `json:"statusTime"`
	ProjectId        int      `json:"projectId"`
	VpcId            int      `json:"vpcId"`
	SubnetId         int      `json:"subnetId"`
}

func (client *Client) DescribeLoadBalancer(loadBalancerId string) (response *DescribeLoadBalancerResponse, err error) {
	response = &DescribeLoadBalancerResponse{}

	config := map[string]interface{}{
		"secretId":  client.SecretId,
		"secretKey": client.SecretKey,
		"debug":     false,
	}

	params := map[string]interface{}{
		"Action":            "DescribeLoadBalancers",
		"Region":            client.Region,
		"loadBalancerIds.1": loadBalancerId,
	}

	clb, err := QcloudApi.SendRequest("lb", params, config)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(clb), response)
	return response, nil
}

// delete tencent cloud load balancer by load balancer id

type DeleteLoadBalancerResponse struct {
	CommonResponse
	RequestId int `json:"requestId"`
}

func (client *Client) DeleteLoadBalancer(loadBalancerId string) (response *DeleteLoadBalancerResponse, err error) {
	response = &DeleteLoadBalancerResponse{}

	config := map[string]interface{}{
		"secretId":  client.SecretId,
		"secretKey": client.SecretKey,
		"debug":     false,
	}

	params := map[string]interface{}{
		"Action":            "DeleteLoadBalancers",
		"Region":            client.Region,
		"loadBalancerIds.1": loadBalancerId,
	}

	clb, err := QcloudApi.SendRequest("lb", params, config)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(clb), response)
	return response, nil
}

// update tencent cloud load balancer, the api only can update load balanacer name and domian prefix

type ModifyLoadBalancerAttributesArgs struct {
	LoadBalancerId   string
	LoadBalancerName string
	DomainPrefix     string
}

type ModifyLoadBalancerAttributesResponse struct {
	CommonResponse
	RequestId int `json:"requestId"`
}

func (client *Client) ModifyLoadBalancerAttributes(args *ModifyLoadBalancerAttributesArgs) (response *ModifyLoadBalancerAttributesResponse, err error) {
	response = &ModifyLoadBalancerAttributesResponse{}

	config := map[string]interface{}{
		"secretId":  client.SecretId,
		"secretKey": client.SecretKey,
		"debug":     false,
	}

	params := map[string]interface{}{
		"Action":           "ModifyLoadBalancerAttributes",
		"Region":           client.Region,
		"loadBalancerId":   args.LoadBalancerId,
		"loadBalancerName": args.LoadBalancerName,
		"domainPrefix":     args.DomainPrefix,
	}

	clb, err := QcloudApi.SendRequest("lb", params, config)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(clb), response)
	return response, nil
}
