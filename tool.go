package gotool

import (
        "net/http"
        "os"
        "log"
        "os/exec"
        )
func EndServices(req *http.Request){
        if req.FormValue("e")=="e"{
                os.Exit(0)
}
}
func OpenBrowser(url string){
        cmd := exec.Command("cmd", "/c","start",url);
        err := cmd.Start()
        if err != nil {
                log.Fatal(err)
                cmd:=exec.Command("start",url);
                err := cmd.Start()
                if err != nil {
                        log.Fatal(err)
                }
        }
}
