// Copyright 2024 Commonwealth Scientific and Industrial Research Organisation (CSIRO) ABN 41 687 119 230
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	secret "github.com/ivcap-works/ivcap-core-api/gen/secret"
)

// BuildListPayload builds the payload for the secret list endpoint from CLI
// flags.
func BuildListPayload(secretListPage string, secretListFilter string, secretListOffset string, secretListLimit string, secretListJWT string) (*secret.ListPayload, error) {
	var err error
	var page *string
	{
		if secretListPage != "" {
			page = &secretListPage
		}
	}
	var filter *string
	{
		if secretListFilter != "" {
			filter = &secretListFilter
		}
	}
	var offset *string
	{
		if secretListOffset != "" {
			offset = &secretListOffset
		}
	}
	var limit *int
	{
		if secretListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(secretListLimit, 10, strconv.IntSize)
			val := int(v)
			limit = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
		}
	}
	var jwt string
	{
		jwt = secretListJWT
	}
	v := &secret.ListPayload{}
	v.Page = page
	v.Filter = filter
	v.Offset = offset
	v.Limit = limit
	v.JWT = jwt

	return v, nil
}

// BuildGetPayload builds the payload for the secret get endpoint from CLI
// flags.
func BuildGetPayload(secretGetSecretName string, secretGetSecretType string, secretGetJWT string) (*secret.GetPayload, error) {
	var secretName string
	{
		secretName = secretGetSecretName
	}
	var secretType *string
	{
		if secretGetSecretType != "" {
			secretType = &secretGetSecretType
		}
	}
	var jwt string
	{
		jwt = secretGetJWT
	}
	v := &secret.GetPayload{}
	v.SecretName = secretName
	v.SecretType = secretType
	v.JWT = jwt

	return v, nil
}

// BuildSetPayload builds the payload for the secret set endpoint from CLI
// flags.
func BuildSetPayload(secretSetBody string, secretSetJWT string) (*secret.SetPayload, error) {
	var err error
	var body SetRequestBody
	{
		err = json.Unmarshal([]byte(secretSetBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"expiry-time\": 5575456846245597618,\n      \"secret-name\": \"Ducimus nostrum.\",\n      \"secret-type\": \"Aut laborum qui unde rem assumenda.\",\n      \"secret-value\": \"Magnam accusamus enim omnis est.\"\n   }'")
		}
	}
	var jwt string
	{
		jwt = secretSetJWT
	}
	v := &secret.SetSecretRequestT{
		SecretName:  body.SecretName,
		SecretType:  body.SecretType,
		SecretValue: body.SecretValue,
		ExpiryTime:  body.ExpiryTime,
	}
	res := &secret.SetPayload{
		Secrets: v,
	}
	res.JWT = jwt

	return res, nil
}
