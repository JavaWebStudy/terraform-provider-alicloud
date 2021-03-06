package ons

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// OnsMqttQueryMsgTransTrend invokes the ons.OnsMqttQueryMsgTransTrend API synchronously
// api document: https://help.aliyun.com/api/ons/onsmqttquerymsgtranstrend.html
func (client *Client) OnsMqttQueryMsgTransTrend(request *OnsMqttQueryMsgTransTrendRequest) (response *OnsMqttQueryMsgTransTrendResponse, err error) {
	response = CreateOnsMqttQueryMsgTransTrendResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMqttQueryMsgTransTrendWithChan invokes the ons.OnsMqttQueryMsgTransTrend API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmqttquerymsgtranstrend.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMqttQueryMsgTransTrendWithChan(request *OnsMqttQueryMsgTransTrendRequest) (<-chan *OnsMqttQueryMsgTransTrendResponse, <-chan error) {
	responseChan := make(chan *OnsMqttQueryMsgTransTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMqttQueryMsgTransTrend(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// OnsMqttQueryMsgTransTrendWithCallback invokes the ons.OnsMqttQueryMsgTransTrend API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmqttquerymsgtranstrend.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMqttQueryMsgTransTrendWithCallback(request *OnsMqttQueryMsgTransTrendRequest, callback func(response *OnsMqttQueryMsgTransTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMqttQueryMsgTransTrendResponse
		var err error
		defer close(result)
		response, err = client.OnsMqttQueryMsgTransTrend(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// OnsMqttQueryMsgTransTrendRequest is the request struct for api OnsMqttQueryMsgTransTrend
type OnsMqttQueryMsgTransTrendRequest struct {
	*requests.RpcRequest
	TransType   string           `position:"Query" name:"TransType"`
	EndTime     requests.Integer `position:"Query" name:"EndTime"`
	BeginTime   requests.Integer `position:"Query" name:"BeginTime"`
	TpsType     string           `position:"Query" name:"TpsType"`
	ParentTopic string           `position:"Query" name:"ParentTopic"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	Qos         requests.Integer `position:"Query" name:"Qos"`
	MsgType     string           `position:"Query" name:"MsgType"`
	SubTopic    string           `position:"Query" name:"SubTopic"`
}

// OnsMqttQueryMsgTransTrendResponse is the response struct for api OnsMqttQueryMsgTransTrend
type OnsMqttQueryMsgTransTrendResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateOnsMqttQueryMsgTransTrendRequest creates a request to invoke OnsMqttQueryMsgTransTrend API
func CreateOnsMqttQueryMsgTransTrendRequest() (request *OnsMqttQueryMsgTransTrendRequest) {
	request = &OnsMqttQueryMsgTransTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMqttQueryMsgTransTrend", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsMqttQueryMsgTransTrendResponse creates a response to parse from OnsMqttQueryMsgTransTrend response
func CreateOnsMqttQueryMsgTransTrendResponse() (response *OnsMqttQueryMsgTransTrendResponse) {
	response = &OnsMqttQueryMsgTransTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
