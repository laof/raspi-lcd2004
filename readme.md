### 无桌面版配置WiFi
在boot盘中新建 wpa_supplicant.conf 文件，内容如下：

```
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
country=CN
 
network={
	ssid="wifiname"
	psk="psw"
}
```


### 自启动

build
```
go build lcd2004
chmod 777 lcd2004
```

编辑文件 sudo nano /etc/rc.local
```
...
sudo /home/pi/raspi-lcd2004/lcd2004
exit 0
```

### meadme
https://github.com/laof/raspi-lcd2004

### Install latest Golang
https://golang.google.cn/dl

arm64是针对于64位的树莓派系统, 下载arm6l


```
wget https://golang.google.cn/dl/go1.20.1.linux-armv6l.tar.gz
```

安装
```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.1.linux-armv6l.tar.gz
```

查看
```
go version
```

### Enable I2C

```
sudo raspi-config
```

侦测连接设备

```
i2cdetect -y 1
```
```

     0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
00:                         -- -- -- -- -- -- -- --       
10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --   
20: -- -- -- -- -- -- -- 27 -- -- -- -- -- -- -- --   
30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --   
40: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --   
50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --   
60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --   
70: -- -- -- -- -- -- -- --
```

### display

![image](https://github.com/laof/laof.github.io/raw/main/img/pi/golang.png)


### I2C connection

![image](https://github.com/laof/laof.github.io/raw/main/img/pi/lcd.png)

### gpio

![image](https://github.com/laof/laof.github.io/raw/main/img/pi/gpio.png)

