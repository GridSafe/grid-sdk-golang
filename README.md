# 格安Golang SDK

## Usage

1. 修改./src/config/config.go 文件

    - 修改UserName 和 UserSignature
    - 在[账户页面](https://www.cdnzz.com/account) 可以看到你的接口标识(signature), UserName 则是你的电邮地址

2. import 需要的模块

  ```
  import "./grid-sdk-golang/src/sdk"
  ```

3. 调用相应的操作

  ```
  func main() {
      result, msg := grid_sdk.Preload("http://www.example.com/img.js")
      fmt.Println(result)
      fmt.Println(msg)
  }
  ```

## Change Log

Please see 'CHANGE_LOG.md'



