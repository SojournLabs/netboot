/*
 * Program that serves configuration settings and OS images for server infrastructure.
 * @author jonathan lung
 * Don't judge. This is my first go program to test out the language. It's also a quick
 * hack.
 */
package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

/*
 *  Fill in a template for the boot process for a machine.
 */
func templateHandler(w http.ResponseWriter, r *http.Request) {
	re, _ := regexp.Compile(`^/template/([^/]*)(?:/(.*))?`)
	matches := re.FindStringSubmatch(r.URL.Path)
	templateFile := matches[1]

	// Lowercase, no-colon representation of MAC address.
	macAddress := strings.Replace(strings.ToLower(matches[2]), ":", "", -1)

	// Get the remote IP address.
	remoteIP := getRemoteIP(r)

	serverConfig := make(map[string]string)

	// Load master configuration data first.
	bootData, err := ioutil.ReadFile("config/global")
	if err == nil {
		yaml.Unmarshal(bootData, &serverConfig)
	}

	// Load IP-specific configuration data next.
	bootData, err = ioutil.ReadFile("config/" + remoteIP)
	if err == nil {
		// Read file into a temporary map and then merge.
		var tempMap map[string]string
		yaml.Unmarshal(bootData, &tempMap)

		for k, v := range tempMap {
			serverConfig[k] = v
		}
	}

	// Lastly, load MAC-address-specific configuration.
	bootData, err = ioutil.ReadFile("config/" + macAddress)
	if err == nil {
		// Read file into a temporary map and then merge.
		var tempMap map[string]string
		yaml.Unmarshal(bootData, &tempMap)

		for k, v := range tempMap {
			serverConfig[k] = v
		}
	}

	// Populate default IP address and MAC address.
	_, ok := serverConfig["ipAddress"]
	if !ok {
		serverConfig["ipAddress"] = remoteIP
	}

	_, ok = serverConfig["macAddress"]
	if !ok {
		serverConfig["macAddress"] = macAddress
	}

	templateData, _ := ioutil.ReadFile("templates/" + templateFile)
	templ, _ := template.New("templ").Parse(string(templateData))
	templ.Execute(w, serverConfig)

	return
}

/*
 * Retrieve a file (static image) directly from the images directory.
 */
func imageHandler(w http.ResponseWriter, r *http.Request) {
	imageFilename := r.URL.Path[len("/images/"):]
	if strings.Contains(imageFilename, "/") {
		return
	}

	stats, _ := os.Lstat("images/" + imageFilename)
	w.Header().Set("Content-Length", strconv.FormatInt(stats.Size(), 10))

	content, _ := ioutil.ReadFile("images/" + imageFilename)
	w.Write(content)
}

/*
 * Get the IP of the connection.
 */
func getRemoteIP(r *http.Request) string {
	if regexp.MustCompile(`^127\..*:.*|\[::1\]:.*$`).MatchString(r.RemoteAddr) {
		return r.Header.Get("X-Forwarded-For")
	} else {
		// Trim off port
		return regexp.MustCompile(`^(.*):.*$`).FindAllStringSubmatch(r.RemoteAddr, -1)[0][1]
	}
}

func main() {
	http.HandleFunc("/template/", templateHandler)
	http.HandleFunc("/images/", imageHandler)
	http.ListenAndServe(":8800", nil)
}
