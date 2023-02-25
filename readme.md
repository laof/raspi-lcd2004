

### Install latest golang
https://golang.google.cn/dl/

arm64是针对于64位的树莓派系统

下载arm6l
```
wget https://golang.google.cn/dl/go1.17.6.linux-armv6l.tar.gz
```


安装
```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.6.linux-armv6l.tar.gz
```


查看
```
go version
```

### Enable I2C

```
sudo raspi-config
```

```
i2cdetect -y 1
```

### display

![image](https://github.com/laof/laof.github.io/raw/main/img/pi/golang.png)


### I2C connection

![image](https://github.com/laof/laof.github.io/raw/main/img/pi/lcd.png)

### gpio

![image](https://github.com/laof/laof.github.io/raw/main/img/pi/gpio.png)

