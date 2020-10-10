package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"unicode/utf8"
)

/*
* PreSignV2 : Base64( HMAC-SHA1( UTF-8-Encoding-Of(YourSecretAccessKey), UTF-8-Encoding-Of( StringToSign ) ) )
* baseURL   : http://10.51.201.201 or http://172.16.31.201 or https://s3.harix.iamidata.com
* output    : http://10.51.201.201/bucket/object?AWSAccessKeyId=XXX&Expires=4743019496&Signature=XXX
* default objectURL expire date is: Sat Apr 20 09:24:56 CST 2120
 */
func PreSignV2(accKey, secKey, bucket, object string) (string, error) {
	var (
		StringToSign string
		params       = url.Values{}
		baseURL      = "https://s3.harix.iamidata.com"
		expireTTL    = "4743019496"
	)
	if !utf8.ValidString(bucket) || !utf8.ValidString(object) {
		return "", errors.New("Invalid utf-8 string")
	}
	StringToSign = fmt.Sprintf("GET\n\n\n%s\n/%s/%s", expireTTL, bucket, object)
	keyForSign := []byte(secKey)
	h := hmac.New(sha1.New, keyForSign)
	h.Write([]byte(StringToSign))
	sig := base64.StdEncoding.EncodeToString(h.Sum(nil))
	params.Add("AWSAccessKeyId", accKey)
	params.Add("Expires", expireTTL)
	params.Add("Signature", sig)
	return fmt.Sprintf("%s/%s/%s?%s", baseURL, bucket, object, params.Encode()), nil
}

func main() {
	var (
		acc    = "5373OR9D1ZA5UD6FWE6O"
		sec    = "xPfIfXBjqMnt62dZA9c2wXXCmLVPaMUOmMBt3M6H"
		bucket = "dms"
		obj    = "vbn-cloudpepper-v3-uit-mcs-mysql/20201009234312/vbn-cloudpepper-v3-uit-mcs-mysql.binlog-on.dump.gz"
	)
	url, err := PreSignV2(acc, sec, bucket, obj)
	fmt.Println(url, err)
}
