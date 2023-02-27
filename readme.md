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

update resource
```
sudo apt update
sudo apt upgrade
```


### autostart

build
```
go build -o lcd2004
chmod 777 lcd2004
chmod +x lcd2004
```

edit
```
sudo nano /etc/rc.local
```
add
```
...
sudo /xxxxx &
sudo /home/pi/raspi-lcd2004/lcd2004
exit 0
```


### Install latest Golang

arm64 for raspberry 64 bit system, we will to download xx.arm6l file, for more version please to see https://golang.google.cn/dl 

```
wget https://golang.google.cn/dl/go1.20.1.linux-armv6l.tar.gz
```

installing
```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.1.linux-armv6l.tar.gz
```

editing configuration

```
sudo nano ~/.profile
```

adding path

```
...
# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/.local/bin" ] ; then
    PATH="$HOME/.local/bin:$PATH"
fi

# set go path here
export PATH=$PATH:/usr/local/go/bin
```

updating
```
source ~/.profile
```

check
```
go version
go env -w GOPROXY=https://goproxy.cn,direct

```

### Enable I2C

```
sudo raspi-config
```

detect connected devices

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



### clone 
```
git clone https://github.com/laof/raspi-lcd2004.git
```

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