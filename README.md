MWeb to Azure Storage Service
======
`MWeb`是一款Mac 下不错的markdown笔记软件，但是支持图床太少。经过
我一段时间的筛选，我觉得Azure的Storage 服务不错（如果能搞到Azure国际版的帐号的话）。

这个服务运行在服务器后台，可以将`Mweb`的图床切换到Azure，目前七牛不绑定域名的话，请求次数是有限制的，
而域名绑定需要备案，觉得麻烦。所以就花时间写了这个。但目前Azure并不支持https的CNAME，略遗憾。

如果可能，未来也会支持更多的图床，欢迎给我提issus，PR。

我个人提供了一个入口，如果你不放心的话，你也可以自己架设:
配置方式如下图所示(看不到图片请直接点击链接 https://qhsong.blob.core.windows.net/qhsong-blog/2016/08/intro.png ):

![](https://qhsong.blob.core.windows.net/qhsong-blog/2016/08/intro.png)


或者可以使用如下配置参数：
* API地址：https://cimg.sqh.me/api/upload/azure
* Post file name: file
* 图片URL路径: azureURL

同时添加如下HTTP body：
* `azureKey`: Azure Storage 的AccessKey
* `azureName`: Azure Storage 的用户名称
* `appendDate`: True/False 是否在文件名之前添加日期
* `azureContainer`: Azure Storage Container名称，如果不存在会新建


当然，你可以folk我的代码再架设一份在自己私有的空间。
