# memcheck
This is light-weight tool using which you can check what is current status of amount of disk space and used notify usviaing email to relevant authority/authorities to take necceassy steps. Low disk space can cause multiple issues, but the prime one is that, if the disk space becomes very low on server (it mostly happens due to logs) it may crash. Hence it is very important to be aware of disk space and take urgent steps whenever the need be.  
# Installation
1)  `git clone https://github.com/souvikhaldar/memcheck.git`  
2)  [Install Go](https://golang.org/doc/install)  
3)  `cd memcheck`  
4)  Fill details in `config.json`  
```
{
    "SourceMail": "abc@gmail.com",
    "SourcePassword": "xyz",
    "TargetMail":["xyz@typ.com"]
}
```
4)  `go build`  

# Usage  
Now you can run the `memcheck` binary generated inside the directory. Ideally one is supposed to run this binary using a cron job at a suitable interval.  
