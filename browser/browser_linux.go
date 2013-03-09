 
// +build linux

package browser
import (
	"os/exec"
	"log"
)
type Browser struct{
}
func (p * Browser) OpenBrowserSync(url string){
	openBrowser(url)
}
func (p * Browser) OpenBrowserAsync(url string){
	go openBrowser(url)
}
func openBrowser(url string){
	cmd := exec.Command("start",url);
	 err := cmd.Start()
	 if err!=nil{
	 	log.Fatal(err)
	 }
}
