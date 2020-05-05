package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/souvikhaldar/gomail"
)

func main() {
	configPath := flag.String("config", "config.json", "path for config file")
	flag.Parse()
	conf := ParseConfig(*configPath)
	ip, _ := exec.Command("curl", "ifconfig.co").Output()
	if len(ip) > 20 {
		log.Println("IP could not be traced")
		return
	}
	out, _ := exec.Command("df", "-P").Output()
	used, total, err := parseDfOutput(string(out))
	if err != nil {
		log.Println(err)
		return
	}
	if conf == nil {
		fmt.Println("nil config")
		return
	}
	if conf.SourceMail == "" || conf.SourcePassword == "" || len(conf.TargetMail) == 0 {
		log.Println("nil config")
		return
	}
	e, config := gomail.New(conf.SourceMail, conf.SourcePassword)
	if e != nil {
		fmt.Print(fmt.Errorf("Error in creating config %v", e))
		return
	}

	ratio := (used / total)
	switch {
	case ratio < 0.75:
		log.Println("Disk space is fine")
		body := fmt.Sprintf("Total: %f GB\nUsed: %f GB", total, used)
		log.Println(body)
		return
	case ratio >= 0.75 && ratio < 0.85:
		body := fmt.Sprintf("Disk space on %s is low. Total: %f GB\nUsed: %f GB", ip, total, used)
		log.Println(body)
		if e := config.SendMail(conf.TargetMail, "Disk space alert", body); e != nil {
			log.Print(fmt.Errorf("Error in sending mail %v", e))
			return
		}
		return
	case ratio >= 0.85:
		body := fmt.Sprintf("Disk space on %s is critically low. Total: %f GB\nUsed: %f GB", ip, total, used)
		log.Println(body)
		if e := config.SendMail(conf.TargetMail, "Disk space alert", body); e != nil {
			log.Print(fmt.Errorf("Error in sending mail %v", e))
			return
		}
		return
	}

}

func parseDfOutput(out string) (float64, float64, error) {
	outlines := strings.Split(out, "\n")
	l := len(outlines)
	var total, used float64 = 0, 0
	for _, line := range outlines[1 : l-1] {
		parsedLine := strings.Fields(line)
		t, err := strconv.ParseFloat(parsedLine[1], 64)
		if err != nil {
			return 0, 0, err
		}
		u, err := strconv.ParseFloat(parsedLine[2], 64)
		if err != nil {
			return 0, 0, err
		}
		total += t
		used += u
	}

	return used / (1024 * 1024), total / (1024 * 1024), nil
}
