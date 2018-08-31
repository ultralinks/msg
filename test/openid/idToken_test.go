package openid

import (
    "testing"
    "msg/openid"
)

const (
    publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDdlatRjRjogo3WojgGHFHYLugd
UWAY9iR3fy4arWNA1KoS8kVw33cJibXr8bvwUAUparCwlvdbH6dvEOfou0/gCFQs
HUfQrSDv+MuSUMAe8jzKE4qW+jK+xQU9a03GUnKHkkle+Q0pX/g6jXZ7r1/xAK5D
o2kQ+X5xK9cipRgEKwIDAQAB
-----END PUBLIC KEY-----`;
    idTokenStr = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3d3cuZ2FwdHJlZS5jb20iLCJzdWIiOiIxMjM0NTY3ODkwIiwiaWF0IjoxNTE2MjM5MDIyfQ.q8g0usPhtrJxGXUlBGvHsAH9DCmVJwNB0M03wFlh-EZM_vntp8V5rv-NFnY1i1noIY7Li7cajc6sQpx9_w9Y_BnmX9RTRQTf9IbGMQ1Pi-nozh_1RauoB3LH7g6wnOWnUdD9FeexUb8as6lA6nkh4nczdKVgnHGq9qVS3ksR1Qc`
)

// jwt.io

/** header
{
  "alg": "RS256",
  "typ": "JWT"
}
*/

/** payload
{
  "iss": "www.gaptree.com",
  "sub": "1234567890",
  "iat": 1516239022
}
*/

func TestIdToken(t *testing.T) {
    idToken, err := openid.ParseIdToken(idTokenStr, publicKey)
    if err != nil {
        t.Error(err)
    }
    if idToken.StandardClaims.Issuer != "www.gaptree.com" {
        t.Error(idToken)
    }
}
