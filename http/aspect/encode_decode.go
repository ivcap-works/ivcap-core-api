// $ goa gen github.com/ivcap-works/ivcap-core-api/design

package client

import (
	"bytes"
	aspect "github.com/ivcap-works/ivcap-core-api/gen/aspect"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "aspect" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*aspect.ReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("aspect", "read", "*aspect.ReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReadAspectPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("aspect", "read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeReadRequest returns an encoder for requests sent to the aspect read
// server.
func EncodeReadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*aspect.ReadPayload)
		if !ok {
			return goahttp.ErrInvalidType("aspect", "read", "*aspect.ReadPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeReadResponse returns a decoder for responses returned by the aspect
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "bad-request" (type *aspect.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *aspect.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-scopes" (type *aspect.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *aspect.NotImplementedT): http.StatusNotImplemented
//   - "not-found" (type *aspect.ResourceNotFoundT): http.StatusNotFound
//   - "not-authorized" (type *aspect.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeReadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ReadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "read", err)
			}
			err = ValidateReadResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "read", err)
			}
			res := NewReadAspectRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body ReadBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("aspect", "read", err)
				}
				err = ValidateReadBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("aspect", "read", err)
				}
				return nil, NewReadBadRequest(&body)
			case "invalid-credential":
				return nil, NewReadInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("aspect", "read", resp.StatusCode, string(body))
			}
		case http.StatusForbidden:
			var (
				body ReadInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "read", err)
			}
			err = ValidateReadInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "read", err)
			}
			return nil, NewReadInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ReadNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "read", err)
			}
			err = ValidateReadNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "read", err)
			}
			return nil, NewReadNotImplemented(&body)
		case http.StatusNotFound:
			var (
				body ReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "read", err)
			}
			err = ValidateReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "read", err)
			}
			return nil, NewReadNotFound(&body)
		case http.StatusUnauthorized:
			return nil, NewReadNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("aspect", "read", resp.StatusCode, string(body))
		}
	}
}

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "aspect" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListAspectPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("aspect", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the aspect list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*aspect.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("aspect", "list", "*aspect.ListPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		values := req.URL.Query()
		if p.Entity != nil {
			values.Add("entity", *p.Entity)
		}
		if p.Schema != nil {
			values.Add("schema", *p.Schema)
		}
		if p.AtTime != nil {
			values.Add("at-time", *p.AtTime)
		}
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		values.Add("filter", p.Filter)
		values.Add("order-by", p.OrderBy)
		if p.OrderDesc != nil {
			values.Add("order-desc", fmt.Sprintf("%v", *p.OrderDesc))
		}
		if p.Page != nil {
			values.Add("page", *p.Page)
		}
		req.URL.RawQuery = values.Encode()
		body := NewListRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("aspect", "list", err)
		}
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the aspect
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad-request" (type *aspect.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *aspect.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *aspect.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *aspect.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *aspect.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *aspect.UnauthorizedT): http.StatusUnauthorized
//   - "unsupported-content-type" (type *aspect.UnsupportedContentType): http.StatusUnsupportedMediaType
//   - error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "list", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "list", err)
			}
			res := NewListAspectListRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body ListBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("aspect", "list", err)
				}
				err = ValidateListBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("aspect", "list", err)
				}
				return nil, NewListBadRequest(&body)
			case "invalid-credential":
				return nil, NewListInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("aspect", "list", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body ListInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "list", err)
			}
			err = ValidateListInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "list", err)
			}
			return nil, NewListInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body ListInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "list", err)
			}
			err = ValidateListInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "list", err)
			}
			return nil, NewListInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body ListNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "list", err)
			}
			err = ValidateListNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "list", err)
			}
			return nil, NewListNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewListNotAuthorized()
		case http.StatusUnsupportedMediaType:
			var (
				body ListUnsupportedContentTypeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "list", err)
			}
			err = ValidateListUnsupportedContentTypeResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "list", err)
			}
			return nil, NewListUnsupportedContentType(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("aspect", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "aspect" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateAspectPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("aspect", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the aspect
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*aspect.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("aspect", "create", "*aspect.CreatePayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		{
			head := p.ContentType
			req.Header.Set("Content-Type", head)
		}
		values := req.URL.Query()
		values.Add("entity", p.Entity)
		values.Add("schema", p.Schema)
		if p.Policy != nil {
			values.Add("policy", *p.Policy)
		}
		req.URL.RawQuery = values.Encode()
		body := p.Content
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("aspect", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the aspect
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeCreateResponse may return the following errors:
//   - "bad-request" (type *aspect.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *aspect.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *aspect.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *aspect.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *aspect.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *aspect.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "create", err)
			}
			err = ValidateCreateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "create", err)
			}
			res := NewCreateAspectIDRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body CreateBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("aspect", "create", err)
				}
				err = ValidateCreateBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("aspect", "create", err)
				}
				return nil, NewCreateBadRequest(&body)
			case "invalid-credential":
				return nil, NewCreateInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("aspect", "create", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body CreateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "create", err)
			}
			err = ValidateCreateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "create", err)
			}
			return nil, NewCreateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body CreateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "create", err)
			}
			err = ValidateCreateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "create", err)
			}
			return nil, NewCreateInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body CreateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "create", err)
			}
			err = ValidateCreateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "create", err)
			}
			return nil, NewCreateNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewCreateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("aspect", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "aspect" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*aspect.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("aspect", "update", "*aspect.UpdatePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateAspectPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("aspect", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the aspect
// update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*aspect.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("aspect", "update", "*aspect.UpdatePayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		{
			head := p.ContentType
			req.Header.Set("Content-Type", head)
		}
		values := req.URL.Query()
		values.Add("entity", p.Entity)
		values.Add("schema", p.Schema)
		req.URL.RawQuery = values.Encode()
		body := p.Content
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("aspect", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the aspect
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUpdateResponse may return the following errors:
//   - "bad-request" (type *aspect.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *aspect.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *aspect.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *aspect.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *aspect.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *aspect.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "update", err)
			}
			err = ValidateUpdateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "update", err)
			}
			res := NewUpdateAspectIDRTOK(&body)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body UpdateBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("aspect", "update", err)
				}
				err = ValidateUpdateBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("aspect", "update", err)
				}
				return nil, NewUpdateBadRequest(&body)
			case "invalid-credential":
				return nil, NewUpdateInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("aspect", "update", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body UpdateInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "update", err)
			}
			err = ValidateUpdateInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "update", err)
			}
			return nil, NewUpdateInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body UpdateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "update", err)
			}
			err = ValidateUpdateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "update", err)
			}
			return nil, NewUpdateInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body UpdateNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "update", err)
			}
			err = ValidateUpdateNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "update", err)
			}
			return nil, NewUpdateNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewUpdateNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("aspect", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildRetractRequest instantiates a HTTP request object with method and path
// set to call the "aspect" service "retract" endpoint
func (c *Client) BuildRetractRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*aspect.RetractPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("aspect", "retract", "*aspect.RetractPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RetractAspectPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("aspect", "retract", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRetractRequest returns an encoder for requests sent to the aspect
// retract server.
func EncodeRetractRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*aspect.RetractPayload)
		if !ok {
			return goahttp.ErrInvalidType("aspect", "retract", "*aspect.RetractPayload", v)
		}
		{
			head := p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeRetractResponse returns a decoder for responses returned by the aspect
// retract endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeRetractResponse may return the following errors:
//   - "bad-request" (type *aspect.BadRequestT): http.StatusBadRequest
//   - "invalid-credential" (type *aspect.InvalidCredentialsT): http.StatusBadRequest
//   - "invalid-parameter" (type *aspect.InvalidParameterValue): http.StatusUnprocessableEntity
//   - "invalid-scopes" (type *aspect.InvalidScopesT): http.StatusForbidden
//   - "not-implemented" (type *aspect.NotImplementedT): http.StatusNotImplemented
//   - "not-authorized" (type *aspect.UnauthorizedT): http.StatusUnauthorized
//   - error: internal error
func DecodeRetractResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "bad-request":
				var (
					body RetractBadRequestResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("aspect", "retract", err)
				}
				err = ValidateRetractBadRequestResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("aspect", "retract", err)
				}
				return nil, NewRetractBadRequest(&body)
			case "invalid-credential":
				return nil, NewRetractInvalidCredential()
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("aspect", "retract", resp.StatusCode, string(body))
			}
		case http.StatusUnprocessableEntity:
			var (
				body RetractInvalidParameterResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "retract", err)
			}
			err = ValidateRetractInvalidParameterResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "retract", err)
			}
			return nil, NewRetractInvalidParameter(&body)
		case http.StatusForbidden:
			var (
				body RetractInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "retract", err)
			}
			err = ValidateRetractInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "retract", err)
			}
			return nil, NewRetractInvalidScopes(&body)
		case http.StatusNotImplemented:
			var (
				body RetractNotImplementedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("aspect", "retract", err)
			}
			err = ValidateRetractNotImplementedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("aspect", "retract", err)
			}
			return nil, NewRetractNotImplemented(&body)
		case http.StatusUnauthorized:
			return nil, NewRetractNotAuthorized()
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("aspect", "retract", resp.StatusCode, string(body))
		}
	}
}

// unmarshalLinkTResponseBodyToAspectLinkT builds a value of type *aspect.LinkT
// from a value of type *LinkTResponseBody.
func unmarshalLinkTResponseBodyToAspectLinkT(v *LinkTResponseBody) *aspect.LinkT {
	res := &aspect.LinkT{
		Rel:  *v.Rel,
		Type: *v.Type,
		Href: *v.Href,
	}

	return res
}

// unmarshalAspectListItemRTResponseBodyToAspectAspectListItemRT builds a value
// of type *aspect.AspectListItemRT from a value of type
// *AspectListItemRTResponseBody.
func unmarshalAspectListItemRTResponseBodyToAspectAspectListItemRT(v *AspectListItemRTResponseBody) *aspect.AspectListItemRT {
	res := &aspect.AspectListItemRT{
		ID:          *v.ID,
		Entity:      *v.Entity,
		Schema:      *v.Schema,
		Content:     v.Content,
		ContentType: *v.ContentType,
	}

	return res
}
