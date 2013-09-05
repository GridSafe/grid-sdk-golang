/*
* author: lycheng
* email: cheng@grid-safe.com
*/

package grid_sdk

import "../config"

import "regexp"
import "net/url"
import "net/http"
import "io/ioutil"
import "encoding/json"

const ApiUrl = "https://www.cdnzz.com/api/json"

type Response struct {
    result string
    error_code int
    msg string
}

// 预加载
// url string: 想要进行预加载的url
//
// @return (result bool, msg string): 前者代表操作是否成功，后者表示具体的信息
func Preload(url string) (result bool, msg string){
    return postRequest("AddPreload", url)
}


// 通过正则来判断URL 是否正确
// url string: 需要判断的url
//
// @return: bool 是否正确
func urlValidate(url string) (bool) {
    r := `^(((http|https|ftp):(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|
            (?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??
            (?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?`
    result, _:= regexp.MatchString(r, url)
    return result
}

// 清除缓存
// url string: 想要清除缓存的url
//
// @return (result bool, msg string): 前者代表操作是否成功，后者表示具体的信息
func PurgeCache(url string) (result bool, msg string){
    return postRequest("PurgeCache", url)
}

// 对api 接口进行post 请求
// @param method string: 操作的行为，清缓存或者预加载
// @param requestUrl string: 请求的具体url
//
// @return (result bool, msg string): 前者代表操作是否成功，后者表示具体的信息
func postRequest(method, requestUrl string)(result bool, msg string) {

    if urlValidate(requestUrl) == false {
        return false, "URL syntax is wrong"
    }

    var postForm = url.Values{"user": {config.UserName},
            "signature": {config.UserSignature},
            "method": {method},
            "url": {requestUrl},
        }

    resp, err := http.PostForm(ApiUrl, postForm)
    if err != nil {
        return false, "network error"
    }

    defer resp.Body.Close()
    body, er := ioutil.ReadAll(resp.Body)
    if er != nil {
        return false, "返回的内容格式有误"
    }

    var dat map[string]interface{}

    r := json.Unmarshal(body, &dat)
    if r != nil {
        return false, "返回的内容格式有误"
    }

    if dat["result"] == "error" {
        return false, dat["msg"].(string)
    }
    return true, dat["msg"].(string)
}
