package main

import (
	"fmt"
	"net"
)

func main() {
	//Kullanıcının Girmiş oldugu WebSitenin URL Adresini almak için bir string tipinde değişken oluşturuyoruz
	var siteurl string
	//Kullanıcıdan Bilgi Alıyoruz
	fmt.Print("Web Sitesi Adress'i Giriniz (örn:google.com):")
	fmt.Scan(&siteurl)

	//oluşturdugumuz func kullanmak için bir değişkene atadık
	nameserverlist := getNameServerList(siteurl)

	//eğer getnameserverlist func'dan gelen boş ise
	if nameserverlist == nil {
		//bu yazı out olacak
		fmt.Println("NameServer Listesi Çekilirken Hata Oluştu...")
	} else {
		//getnameserverlist dolu geliyorsa  bilgiler ekrana yazılıcak
		fmt.Println("Name Server Listesi Yazdırılıyor ....")
		// bir döngü kuruyoru ve 0'dan başlatıyoruz i değeri nameserver değerinin uzunluguna ulaşana kadar 1 arttırıyoruz ve gelen verileri output olarak veriyoruz
		for i := 0; i < len(nameserverlist); i++ {
			fmt.Println(nameserverlist[i])
		}
	}

	ipserverlist := getIpList(siteurl)

	if ipserverlist == nil {
		fmt.Println("İp Listesi Çekilirken Hata Oluştu...")
	} else {
		fmt.Println("İp Listesi Yazdırılıyor ....")
		for i := 0; i < len(ipserverlist); i++ {
			fmt.Println(ipserverlist[i])
		}
	}

	txtlist := getTextList(siteurl)

	if txtlist == nil {
		fmt.Println("TXT Listesi Çekilirken Hata Oluştu...")
		fmt.Println("CNAME GEÇİLİYOR...")
	} else {
		fmt.Println("TXT Listesi Yazdırılıyor ....")
		for i := 0; i < len(txtlist); i++ {
			fmt.Println(txtlist[i])
		}
	}

	fmt.Println("CNAME Listesi Yazdırılıyor...")
	fmt.Println(getCname(siteurl))

}

//NAME LİST FUNC
func getNameServerList(link string) []string {
	//boş bir dizi oluşturduk string tipinde
	var servers []string

	//lookupNS func bize nameserver değerlerini vericek
	nameserver, _ := net.LookupNS(link)

	//nameserver'dan gelen değerleri
	for _, ns := range nameserver {
		//yukarıda tanımladıgımız boş diziye appendfunc ile ekliyoruz
		servers = append(servers, ns.Host)
	}
	return servers
}

//IP LİST FUNC
func getIpList(link string) []string {
	//boş bir dizi oluşturduk string tipinde
	var ips []string
	//lookupIP func bize IP değerlerini vericek
	iplist, _ := net.LookupIP(link)

	//iplist'dan gelen değerleri
	for _, ip := range iplist {
		ips = append(ips, ip.String()) //ip değişkenini string tipine dönüştürdük dizi olarak geldiği için
	}

	return ips
}

func getCname(link string) string {
	cname, _ := net.LookupCNAME(link)
	return cname
}

//TXT LİST FUNC
func getTextList(link string) []string {
	//boş bir dizi oluşturduk string tipinde
	var txts []string
	//lookupTXT func bize TXT değerlerini vericek
	txtList, _ := net.LookupTXT(link)

	for _, txt := range txtList {
		txts = append(txts, txt)
	}

	return txts
}
