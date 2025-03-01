###

切换到 root 用户
```
sudo su -
```

### 0, Enable I2C

```
sudo raspi-config
Interface Options > I2C > Enable
```

### 1，Update Resource

```
sudo apt update
sudo apt upgrade
```

### 2，Install latest Golang

arm64 for raspberry 64 bit system, we will to download xx.arm6l file, for more version please to see https://golang.google.cn/dl 

```
wget https://dl.google.com/go/go1.20.6.linux-armv6l.tar.gz
```

installing
```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.6.linux-armv6l.tar.gz
```

editing env

```
sudo nano ~/.profile
```

```
...
# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/.local/bin" ] ; then
    PATH="$HOME/.local/bin:$PATH"
fi

# set go path here
export PATH=$PATH:/usr/local/go/bin
```

updating source
```
source ~/.profile
```

check

```
go version
```

setup proxy
```
go env -w GOPROXY=https://goproxy.cn,direct
```

### 3, Install  raspi-lcd2004

```
git clone https://github.com/laof/raspi-lcd2004.git
```

### build package
```
go build -o lcd2004

// or
CGO_ENABLED=0 go build -o lcd2004 // fix: gcc: error: unrecognized command-line option '-marm'

```
### chmod package
```
chmod 777 lcd2004
chmod +x lcd2004
```


### setup autostart


```
sudo nano /etc/rc.local
```

```
...
...
sudo /xxxxx &
cd /home/pi/my-fs && sudo ./fs &
sudo /home/pi/raspi-lcd2004/lcd2004
exit 0
```

---------- end ------------------

### find PID
```
$ ps aux   

$ sudo kill xxx
```

### Setup WiFi for lite img
new a file wpa_supplicant.conf in root folder，adding as below

```
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
country=CN
 
network={
	ssid="wifiname"
	psk="psw"
}
```

detect connected devices

```
sudo apt-get install i2c-tools
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





### readme
https://github.com/laof/raspi-lcd2004


### root ssh

```
sudo nano /etc/ssh/sshd_config
```
add permission
```
#PermitRootLogin prohibit-password
#add here
PermitRootLogin yes
```

setup new passwd for root

```
sudo passwd root
```
reboot
```
sudo reboot
```

disk info
```
df -h
```
