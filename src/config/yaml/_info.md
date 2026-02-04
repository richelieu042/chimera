## 使用场景

结构体有yaml tag，
(1) 使用 yamlKit包;
(2) 否则使用 k8sYamlKit包（注意缺点）.

## github.com/goccy/go-yaml

适用场景：反序列化 → 修改 → 再序列化，全程保留注释.

## 暂不支持: 序列化时保留注释

![_img.png](_img.png)
