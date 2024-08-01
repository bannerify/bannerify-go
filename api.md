# V1

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1LivenessResponse">V1LivenessResponse</a>

Methods:

- <code title="get /v1/liveness">client.V1.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1Service.Liveness">Liveness</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1LivenessResponse">V1LivenessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Keys

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeysVerifyKeyResponse">V1KeysVerifyKeyResponse</a>

### GetKey

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyGetKeyGetResponse">V1KeyGetKeyGetResponse</a>

Methods:

- <code title="get /v1/keys.getKey">client.V1.Keys.GetKey.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyGetKeyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyGetKeyGetParams">V1KeyGetKeyGetParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyGetKeyGetResponse">V1KeyGetKeyGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### DeleteKey

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyDeleteKeyNewResponse">V1KeyDeleteKeyNewResponse</a>

Methods:

- <code title="post /v1/keys.deleteKey">client.V1.Keys.DeleteKey.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyDeleteKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyDeleteKeyNewParams">V1KeyDeleteKeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyDeleteKeyNewResponse">V1KeyDeleteKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### CreateKey

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyCreateKeyNewResponse">V1KeyCreateKeyNewResponse</a>

Methods:

- <code title="post /v1/keys.createKey">client.V1.Keys.CreateKey.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyCreateKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyCreateKeyNewParams">V1KeyCreateKeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyCreateKeyNewResponse">V1KeyCreateKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### VerifyKey

Methods:

- <code title="post /v1/keys.verifyKey">client.V1.Keys.VerifyKey.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyVerifyKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyVerifyKeyNewParams">V1KeyVerifyKeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeysVerifyKeyResponse">V1KeysVerifyKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### UpdateKey

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyUpdateKeyNewResponse">V1KeyUpdateKeyNewResponse</a>

Methods:

- <code title="post /v1/keys.updateKey">client.V1.Keys.UpdateKey.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyUpdateKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyUpdateKeyNewParams">V1KeyUpdateKeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyUpdateKeyNewResponse">V1KeyUpdateKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### UpdateRemaining

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyUpdateRemainingNewResponse">V1KeyUpdateRemainingNewResponse</a>

Methods:

- <code title="post /v1/keys.updateRemaining">client.V1.Keys.UpdateRemaining.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyUpdateRemainingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyUpdateRemainingNewParams">V1KeyUpdateRemainingNewParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyUpdateRemainingNewResponse">V1KeyUpdateRemainingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### GetVerifications

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyGetVerificationListResponse">V1KeyGetVerificationListResponse</a>

Methods:

- <code title="get /v1/keys.getVerifications">client.V1.Keys.GetVerifications.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyGetVerificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyGetVerificationListParams">V1KeyGetVerificationListParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1KeyGetVerificationListResponse">V1KeyGetVerificationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## APIs

### GetAPI

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1APIGetAPIGetResponse">V1APIGetAPIGetResponse</a>

Methods:

- <code title="get /v1/apis.getApi">client.V1.APIs.GetAPI.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1APIGetAPIService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1APIGetAPIGetParams">V1APIGetAPIGetParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1APIGetAPIGetResponse">V1APIGetAPIGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### CreateAPI

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1APICreateAPINewResponse">V1APICreateAPINewResponse</a>

Methods:

- <code title="post /v1/apis.createApi">client.V1.APIs.CreateAPI.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1APICreateAPIService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1APICreateAPINewParams">V1APICreateAPINewParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#V1APICreateAPINewResponse">V1APICreateAPINewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# APIs

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#APIDeleteAPIResponse">APIDeleteAPIResponse</a>

Methods:

- <code title="post /v1/apis.deleteApi">client.APIs.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#APIService.DeleteAPI">DeleteAPI</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#APIDeleteAPIParams">APIDeleteAPIParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#APIDeleteAPIResponse">APIDeleteAPIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Keys

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#APIKeyListResponse">APIKeyListResponse</a>

Methods:

- <code title="get /v1/apis/{apiId}/keys">client.APIs.Keys.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#APIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#APIKeyListParams">APIKeyListParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#APIKeyListResponse">APIKeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Ratelimits

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#RatelimitLimitResponse">RatelimitLimitResponse</a>

Methods:

- <code title="post /v1/ratelimits.limit">client.Ratelimits.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#RatelimitService.Limit">Limit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#RatelimitLimitParams">RatelimitLimitParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#RatelimitLimitResponse">RatelimitLimitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Keys

Response Types:

- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#KeyNewResponse">KeyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#KeyVerifyResponse">KeyVerifyResponse</a>

Methods:

- <code title="post /v1/keys">client.Keys.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#KeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#KeyNewParams">KeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#KeyNewResponse">KeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/keys/verify">client.Keys.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#KeyService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#KeyVerifyParams">KeyVerifyParams</a>) (<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go">bannerify</a>.<a href="https://pkg.go.dev/github.com/bannerify/bannerify-go#KeyVerifyResponse">KeyVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
