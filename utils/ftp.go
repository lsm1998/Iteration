package utils

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"iteration/client/config"
	"net"
	"os"
	"path"
	"time"
)

func Transfer(C *config.Config) {
	var (
		err        error
		sftpClient *sftp.Client
	)
	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	sftpClient, err = connect(C.User, C.Password, C.Ip, 22)
	if err != nil {
		fmt.Println(C)
		panic(err)
	}
	defer sftpClient.Close()
	srcFile, err := os.Open(fmt.Sprintf("%s\\%s", C.Path, C.JarName))
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	filePath := fmt.Sprintf("%s\\%s", C.Path, C.JarName)
	var remoteFileName = path.Base(GetFileName(&filePath))
	dstFile, err := sftpClient.Create(path.Join(C.RemoteDir, remoteFileName))
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	stat, _ := srcFile.Stat()
	total := stat.Size() / 1024 / 1024
	fmt.Println("文件大小=", total, "MB")
	curr := 0
	maxSize := 1024 * 1024 * 10
	buf := make([]byte, maxSize)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf[0:n])
		curr += n / (1024 * 1024)
		temp := float64(curr) / float64(total) * 100
		fmt.Println("传输一次，当前进度=", temp, "%")
	}
	fmt.Println("文件上传完成...")
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}
