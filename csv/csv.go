package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Domain struct {
	option1true,
	option1false,
	option2true,
	option2false int
}

var dir string
var domains = make(map[string]*Domain)

func main() {
	flag.StringVar(&dir, "dir", "", "The input directory")
	flag.Parse()
	cdir, err := os.Open(dir)
	if err != nil {
		log.Fatal("Cannot open directory...")
		return
	}
	files, err := cdir.Readdir(-1)
	if err != nil {
		log.Fatal("Cannot open directory...")
		return
	}

	out, err := os.OpenFile("report.csv", os.O_RDWR|os.O_CREATE, 0644)
	defer out.Close()
	if err != nil {
		log.Fatalf("Error writing CSV!%v\n", err)
	}

	for f := range files {
		log.Printf("Starting reading the file %s at %s\n", files[f].Name(), time.Now().Local().Format(time.RFC1123Z))
		cfilename := files[f].Name()
		cfile, err := os.OpenFile(filepath.Join(dir, cfilename), os.O_RDONLY, 0600)
		defer cfile.Close()
		if err != nil {
			log.Printf("Error reading file %s\n%v", cfilename, err)
			continue
		}
		parseFile(cfile)
		log.Printf("Ending parsing the file %s at %s\n", files[f].Name(), time.Now().Local().Format(time.RFC1123Z))
	}

	log.Printf("Got Domain MAP %v\n", domains)
	log.Println("Writing CSV file")
	err = writeCsv(out, domains)
	if err != nil {
		log.Fatalf("Error writing CSV!%v\n", err)
	}

}

func parseFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		request := fields[6]
		status := fields[8]
		if status == "200" {
			reqfields := strings.Split(request, "/")
			site := reqfields[2]
			if len(reqfields) >= 5 {
				//			log.Printf("%v %v %v %v %v\n", reqfields[0], reqfields[1], reqfields[2], reqfields[3], reqfields[4])
				option := reqfields[3]
				val, err := strconv.ParseBool(reqfields[4])
				if err != nil {
					log.Printf("Error parsing Bool value from file %s in line\n%v\n", file.Name(), line)
					continue
				}
				if list, ok := domains[site]; ok {
					list.incrementOption(option, val)
					continue
				} else {
					domain := new(Domain)
					domain.incrementOption(option, val)
					domains[site] = domain
				}

			}
		}
	}
}

// OPTION1 view
// OPTION2 session
// true = adblocked
// false = non ad blocked

func (d *Domain) incrementOption(option string, val bool) {
	switch option {
	case "option1":
		if val {
			d.option1true++
		} else {
			d.option1false++
		}
	case "option2":
		if val {
			d.option2true++
		} else {
			d.option2false++
		}
	default:
		break
	}
}

func writeCsv(outfile *os.File, siteMap map[string]*Domain) error {
	write := csv.NewWriter(outfile)
	write.UseCRLF = true
	//	err := write.Write([]string{"Domain", "PV AdBlocked", "PV NotAdBlocked", "Sessions AdBlocked", "Sessions NotAdBlocked"})
	for site, domain := range siteMap {
		log.Printf("Writing on CSV the domain %s\n", site)
		viewsblocked := strconv.Itoa(domain.option1true)
		viewsnonblocked := strconv.Itoa(domain.option1false)
		sessblocked := strconv.Itoa(domain.option2true)
		sessnonblocked := strconv.Itoa(domain.option2false)
		err := write.Write([]string{site, viewsblocked, viewsnonblocked, sessblocked, sessnonblocked})
		if err != nil {
			log.Fatalf("%v", err)
			return err
		}
		write.Flush()
	}
	return nil
}
