goqr
====

A fast qrcode generate write with google golang.
基于Go语言的快速二维码批量生成器

//English is bad , more english readme will come soon.

###安装测试

1.安装Go语言编译环境

2.安装扩展包

	go get github.com/freezestart/goqr/pkg

3.编译

	go build main.go

4.批量生成

	goqr -data=sometext,anothertext,moretext

png图片将按照输入顺序按序号生成在main目录下的output目录

5.服务端接口

	goqr -server

6.指定服务启动端口

	goqr -server -port=8888

默认8889

###测试结果

测试环境 MacBook Pro / i5 2.5G / 4G / Go 1.0.1

生成并输出1000个二维码 耗时8秒 平均每个耗时8毫秒
