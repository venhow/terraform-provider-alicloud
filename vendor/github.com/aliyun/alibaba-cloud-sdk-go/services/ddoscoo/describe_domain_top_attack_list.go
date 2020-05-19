package ddoscoo

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

// DescribeDomainTopAttackList invokes the ddoscoo.DescribeDomainTopAttackList API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describedomaintopattacklist.html
func (client *Client) DescribeDomainTopAttackList(request *DescribeDomainTopAttackListRequest) (response *DescribeDomainTopAttackListResponse, err error) {
	response = CreateDescribeDomainTopAttackListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainTopAttackListWithChan invokes the ddoscoo.DescribeDomainTopAttackList API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describedomaintopattacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainTopAttackListWithChan(request *DescribeDomainTopAttackListRequest) (<-chan *DescribeDomainTopAttackListResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainTopAttackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainTopAttackList(request)
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

// DescribeDomainTopAttackListWithCallback invokes the ddoscoo.DescribeDomainTopAttackList API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describedomaintopattacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainTopAttackListWithCallback(request *DescribeDomainTopAttackListRequest, callback func(response *DescribeDomainTopAttackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainTopAttackListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainTopAttackList(request)
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

// DescribeDomainTopAttackListRequest is the request struct for api DescribeDomainTopAttackList
type DescribeDomainTopAttackListRequest struct {
	*requests.RpcRequest
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
}

// DescribeDomainTopAttackListResponse is the response struct for api DescribeDomainTopAttackList
type DescribeDomainTopAttackListResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	AttackList []Data `json:"AttackList" xml:"AttackList"`
}

// CreateDescribeDomainTopAttackListRequest creates a request to invoke DescribeDomainTopAttackList API
func CreateDescribeDomainTopAttackListRequest() (request *DescribeDomainTopAttackListRequest) {
	request = &DescribeDomainTopAttackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeDomainTopAttackList", "ddoscoo", "openAPI")
	return
}

// CreateDescribeDomainTopAttackListResponse creates a response to parse from DescribeDomainTopAttackList response
func CreateDescribeDomainTopAttackListResponse() (response *DescribeDomainTopAttackListResponse) {
	response = &DescribeDomainTopAttackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}