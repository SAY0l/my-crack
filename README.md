# My-Crack
更适合中国宝宝的弱口令扫描器

## Help
```
NAME:
   My-Crack - Weak password crack

USAGE:
   main.exe [global options] command [command options] [arguments...]

VERSION:
   1.1

AUTHOR:
   sayol <github@sayol.com>

COMMANDS:
   scan     let's crack weak password
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d                    debug mode
   --timeout value, -t value      timeout (default: 3)
   --scan_num value, -n value     thread num (default: 5000)
   --ip_list value, -i value      ip_list (default: "./input/ip_list.txt")
   --user_dict value, -u value    user_dict (default: "./input/user.dic")
   --pass_dict value, -p value    pass_dict (default: "./input/pass.dic")
   --output_file value, -o value  scan result file (default: "./output/my_crack.txt")
   --help, -h                     show help
   --version, -v                  print the version
```

## Use
当前核心支持ftp/mongodb/mysql/mssql/postgre/redis/ssh的弱口令扫描

提供了编译后的版本。

后续版本开放代码

## Other
思路借鉴于赵海峰老师的《白帽子安全开发实战》

顺带回答赵海峰老师的小问题：
```
1. 扫到一个弱口令后，如何取消相同IP\port和用户名请求，避免扫描效率低下
2. 对于FTP匿名访问，如何只记录一个密码，而不是把所有用户名都记录下来
3. 对于Redis这种没有用户名的服务，如何只记录一次密码，而不是记录所有的所有用户及正常的密码的组合
4. 对于不支持设置超时的扫描插件，如何统一设置超时时间
```

1. 利用hash来完成的。在redis，ftp这种可以仅记录密码的就以（ip-port-protocol）进行hash，其他带用户名的以（ip-port-username）进行hash，查hash来作为规避条件，一旦结果集的hash已经存在和当前扫描产生的hash值相同，就continue。

2. 基于上述原理，就可以实现redis、ftp 只记录一个密码，不会把所有用户名都记录下来，因为是用ip-port-protocol 进行hash。

4. 设计了一个WaitTimeout函数。
```
c:= make(chan struct{})
	go func(){
		defer close(c)
		wg.Wait()
	}()
	select{
	case <-c:
		return false //c仅仅作为一个flag，没有实际作用
	case <-time.After(timeout):
		return true
	}
```
使用一个管道作为flag，time.After(timeout)也返回一个管道。一旦c关闭，就会触发select第一个case，返回false，表示正常关闭，wg.Wait正常完成。一旦超时时间到达，会触发select第二个case，就会返回true，表示已经超时。

## Ver 1.1 优化点
1. 连接核心中，更替了一些不能用的版本、第三方库，使用了维护性高，连接稳定的版本库。

例如mongodb的连接核心：
> "go.mongodb.org/mongo-driver/mongo"

> "go.mongodb.org/mongo-driver/mongo/options"

抛弃了之前的个人维护库。

2. 修复了一些问题
```
ipListFile, err := os.Open(fileName)
	if err != nil {
		logger.Log.Fatalf("Open ip list file err, %v", err)
	}

	defer ipListFile.Close()

减少大量代码冗余。上述代码为例，发生错误会直接终止，defer不会入栈，无需再做defer前的判断。
```
3. 错误处理修复

之前版本大量复用err，会出现err覆盖问题，有必要的err取消覆盖写法。

## Ver 1.2 待优化点
代码上线

并发算法需要优化

更多服务的支持 smb,elastic ...