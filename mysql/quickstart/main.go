package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Info struct {
	Webname      string
	Mysqlhost    string
	Mysqldb      string
	Mysqluname   string
	Mysqlupasswd string
}

type WebInfo Info

func exportDB(info WebInfo, webname string) error {
	argv := []string{"--complete-insert", "--skip-comments", "--compact", "--add-drop-table", "-h" + info.Mysqlhost, "-u" + info.Mysqluname, "-p" + info.Mysqlupasswd, info.Mysqldb}
	//argv := []string{"--complete-insert","--skip-comments","--compact","--add-drop-table","-h"+info.Mysqlhost,"-u"+info.Mysqluname,"-p"+info.Mysqlupasswd,info.Mysqldb}
	cmd := exec.Command("mysqldump", argv...)
	f, err := os.OpenFile(webname+".sql", os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		fmt.Println("打开sql文件失败")
		return err
	}
	defer f.Close()
	cmd.Stdout = f
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Run()
	cmd.Wait()
	return nil
}

func importToLocal(info Info) error {
	//argv := []string{"--reconnect","--default-character-set=utf8","-h"+info.Mysqlhost,"-u"+info.Mysqluname,"-p"+info.Mysqlupasswd,"--database="+info.Mysqldb}
	argv := []string{"--reconnect", "--default-character-set=utf8", "-h" + info.Mysqlhost, "-u" + info.Mysqluname, "-p" + info.Mysqlupasswd, "--database=" + info.Mysqldb}
	cmd := exec.Command("mysql", argv...)
	f, err := os.Open(info.Webname + ".sql")
	if err != nil {
		fmt.Println("读取sql文件失败")
		return err
	}
	defer f.Close()
	cmd.Stdin = f
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Run()
	cmd.Wait()
	return nil
}

func main() {
	fmt.Printf("start")
	webinfo := WebInfo{
		Webname:      "webname",
		Mysqlhost:    "10.12.32.151",
		Mysqldb:      "xinfracloud",
		Mysqluname:   "root",
		Mysqlupasswd: "root",
	}

	err := exportDB(webinfo, webinfo.Webname)
	if err != nil {
		fmt.Println(err)
	}
}
