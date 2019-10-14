package emr

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

// RemoveClusterHosts invokes the emr.RemoveClusterHosts API synchronously
// api document: https://help.aliyun.com/api/emr/removeclusterhosts.html
func (client *Client) RemoveClusterHosts(request *RemoveClusterHostsRequest) (response *RemoveClusterHostsResponse, err error) {
	response = CreateRemoveClusterHostsResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveClusterHostsWithChan invokes the emr.RemoveClusterHosts API asynchronously
// api document: https://help.aliyun.com/api/emr/removeclusterhosts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveClusterHostsWithChan(request *RemoveClusterHostsRequest) (<-chan *RemoveClusterHostsResponse, <-chan error) {
	responseChan := make(chan *RemoveClusterHostsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveClusterHosts(request)
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

// RemoveClusterHostsWithCallback invokes the emr.RemoveClusterHosts API asynchronously
// api document: https://help.aliyun.com/api/emr/removeclusterhosts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveClusterHostsWithCallback(request *RemoveClusterHostsRequest, callback func(response *RemoveClusterHostsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveClusterHostsResponse
		var err error
		defer close(result)
		response, err = client.RemoveClusterHosts(request)
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

// RemoveClusterHostsRequest is the request struct for api RemoveClusterHosts
type RemoveClusterHostsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	HostIdList      *[]string        `position:"Query" name:"HostIdList"  type:"Repeated"`
}

// RemoveClusterHostsResponse is the response struct for api RemoveClusterHosts
type RemoveClusterHostsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ClusterId string `json:"ClusterId" xml:"ClusterId"`
}

// CreateRemoveClusterHostsRequest creates a request to invoke RemoveClusterHosts API
func CreateRemoveClusterHostsRequest() (request *RemoveClusterHostsRequest) {
	request = &RemoveClusterHostsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "RemoveClusterHosts", "emr", "openAPI")
	return
}

// CreateRemoveClusterHostsResponse creates a response to parse from RemoveClusterHosts response
func CreateRemoveClusterHostsResponse() (response *RemoveClusterHostsResponse) {
	response = &RemoveClusterHostsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}