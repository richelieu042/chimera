## 选型

-[Go 验证库 2026 终极对决：go-playground/validator、ozzo-validation 谁是王者](https://mp.weixin.qq.com/s/kA0OdJegMe8sTNkiYpBR9Q)

## !!! v11

go-playground/validator 升级到v11后，需要大改 copy_baked_in.go 、 copy_util.go（先看内置tag有没有新增，没的话在最新源码上修改）.

## 参考

Golang/Golang - 1.wps

## tag name

* 默认: "validate";
* Gin: "binding";
* 可以通过 Validate.SetTagName() 修改tag name.
