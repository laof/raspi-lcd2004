

### Install latest golang
https://golang.google.cn/dl/

arm64是针对于64位的树莓派系统

下载arm6l
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
