## SECURE DEPLOYMENT ENVIRONMENT
### Opis okruženja i aplikacije
 - Za zadatak je korišćena virtuelna mašina. Kao softver za virtuelizaciju korišćen je VirtualBox. Za operativni sistem korišćena je Ubuntu 22.04 LTS verzija slike skinuta sa sajta. Virtuelna mašina je podešena tako da je host mašina može pristupiti preko otvorenih portova 22 (SSH) i 80 (za nginx server). Za SSH pristup je podešeno da se koriste ključevi:
     ```
     cat ~/.ssh/id_rsa.pub | ssh user@address 'cat >> ~/.ssh/authorized_keys'
     ```
 -  Za veb server projekta je korišćen nginx. Sama aplikacija je bazirana na PHP-u i MySql bazi podataka. Aplikacija je vrlo jednostavna i ima za cilj da prikaže TODO listu stavki koje su prethodno dodate u bazu. Pored OpenSSH u firewall-u smo dodali i rule koji omogućava komunikaciju na portu 80 za nginx:
    ```
    sudo ufw allow 'Nginx HTTP'
    ```
    
### SYSTEM REVIEW

  #### OPERATING SYSTEM
   - Ubuntu 22.04.4 LTS je trenutno stabilna verzija. Za sve LTS verzije garantovano je minimum 5 godina standardnog bezbednosnog odžavanja (standard security maintenance) na sve pakete iz ‘Main’ repozitorijuma. Za LTS 22.04 datum isteka standardnog support-a je April 2027, dok je istek legacy support-a April 2034. (https://ubuntu.com/about/release-cycle) 
   
  #### KERNEL
   - Ova verzija kernela je dosta nova pa nema naznaka da je potreban upgrade vezano za ovu distribuciju. Trenutno nije pronadjen bezbednosni vulnerability za ovu verziju.  
   
#### TIME MANAGEMENT
   - Ovde možemo videti da je vremenska zona podešena na UTC. Naša sistem ne radi sa mnogo vremenski senzitivnim procesima pa dodatne konfiguracije nisu potrebne.   

  #### PACKAGES INSTALLED
   - Većina paketa je instalirana prilikom same instalacije sistema. Pošto se radi o server a ne o desktop verziji Ubuntu-a nemamo nepotrebnih paketa vezanih za GUI. Jedan od paketa koji su sada nepotrebni jeste cloud-init. Ovaj paket je bio neophodan samo prilikom konfiguracije za instaliranje na virtuelnoj mašini. Pošto su pored paketa vezanih za minimalnu distribuciju sistema korišćeni i standradni paketi za nginx, mysql i php-fpm (PHP fastCGI process manager) nismo pronašli ranjivosti.
   
   #### LOGGING
   - Za logovanje na našem sistemu koristi se rsyslog
   - Konfiguracija za upravljanje fajlovima i direktorijumima je ograničena permisijama: 

### NETWORK REVIEW

  #### GENERAL INFORMATION
   - Redom smo istražili mrežne interfejse, routing tabelu i DNS konfiguraciju:

  #### FIREWALL RULES
   - Koristeći ufw omogućili smo ulazni saobraćaj samo na portovima 80 (HTTP) i 22 (SSH). Portu 80 treba biti omogućen pristup sa svih adresa jer na njemu trči nginx server, dok bi se saobraćaj na portu 22 mogao ograničiti samo na odredjene adrese. Izlazni saobraćaj je omogućen na adresama koje su već uspostavili konekciju na server. Trebalo bi ograničiti izlazni saobraćaj servera samo na DNS servere i HTTP servere vezane za update softvera. Ručno je provereno da će se ufw pravila uvek primenjivati pri startu servera.

#### IPV6
   - Sva pravila fw koja se primenjuju na IPV4 se primenjuju na IPV6, ali zbog ne korišćenja ove verzije protokola, treba se razmotriti ukidanje IPV6 konekcija.
   
   - Isključena zaštita od CSRF (Cross-Site Request Forgery) napada što može dovesti do toga da treća strana iskoristi autorizaciju korisnika i izvrši izmenjen zahtev umesto njega. Rešava se uvođenjem CSRF tokena (uglavnom jednokratan) ili proveravanjem odakle je stigao zahtev.

### FILESYSTEM REVIEW

  #### MOUNTED PARTITIONS
   - Imamo dve glavne particije i jedan swap koji je prazan. Ne koristimo noatime u nijednom slučaju tako da ćemo uvek znati vreme pristupa.

 #### SENSITIVE FILES
   - Sve bitne konfiguracije vezane za lozinke usera, mysql i nginx servera su zaštićene permisijama. Treba voditi računa o simboličnim linkovima kao što je /etc/mysql/my.cnf -> /etc/mysql/mysql.cnf.
   
   #### SETUID
   - Svi fajlovi iz ove liste imaju dobro podešene permisije.
   
   #### NORMAL FILES
   - Iako app armor modul nije toliko značajan trebalo bi obezbediti bolju permisiju za njega jer je potrebno da ga koriste samo administratori. Fajlovi vezani za veb server su dobro zaštićeni.

   #### BACKUP
   - Ne koristimo posebne backup fajlove za osetljive podatke. Pored toga svi fajlovi su dobro zaštičeni pa tako i backup fajlovi.

### USERS REVIEW

  #### REVIEWING THE PASSWD FILE
   - Samo root i jovan user-i imaju shell, odnosno ostali ne mogu samostalno da izvršavaju komande na sistemu.

 #### REVIEWING THE SHADOW FILE
   - Pošto je \$6$ na početku hash-a lozinke, korišćen je SHA-512 algoritam. Što znači da je naša lozinka lepo hash-ovana. Možemo videti da se po default-u koristi yesscript.
   - Pomoću john the reaper alata uspeli smo da nadjemo lozinku usera za 15min. Korišćen je rockyou rečnik sa predifinisanim algoritmom SHA-512.

#### REVIEWING THE SUDO CONFIGURATION
   - Vidimo da su svi useri zaštićeni:

### SERVICES REVIEW

  #### IDENTIFYING RUNNING SERVICES
 

 #### OPENSSH
   - Zabranjen je direktan login na root nalog sistema. Trebalo bi podesiti AllowTcpForwarding parametar u konfiguraciji na no kako bi onemogućili da se ovaj sistem koristi za pristup drugih sistema. Razmotriti promenu porta kako bi se otežali napadi.

#### MYSQL
   - Videli smo da mysql nije potpuno spolja otvoren. Bind-ovan je localhost kako bi sprečio pristup sa spolja.
   - Ako pogledamo sve usere baze, imamo example_user koji koristimo za našu aplikaciju. Pomoću alata John-The-Ripper možemo pronaću dekriptovanu lozinku, iako se koristi noviji algoritam enkripcije za mysql (*9DFF74CEA805F4A29364492BF22F82DE01005E9A). 
   - Mozemo videti da example_user nema fajl privilegije, što znači da neće moći izvršiti te napade preko mysql-a.


#### NGINX CONFIGURATION
   - Nginx se koristi kao web server i proxy. U fajlu /etc/nginx/nginx.conf možemo videti da je user koji pokrece server www-data. Ostale konfiguracije mogu se naci u conf.d ili sites-enabled folderu. Konfiguracija za todo se nalazi u sites-available a sve to je linkovano na sites-enabled. 
   - Preporuka je ipak koristiti HTPPS sa sertifikatom koji bi se nalazio na mašini. Svi pozivi ka portu 80 (HTTP) u tom slučaju trebju da se redirektuju na 443 (HTTPS). Ostale konfiguracije su minimalne i ne može se pronaći ranjivost,
   - Fajlovi koji su dostupni preko web server-a su zažtićeni permisijama:
   

#### PHP CONFIGURATION
   - Koristimo PHP fastCGI process manager modul koji nam omogućuje vezu izmedju nginx-a i PHP-a. Za razliku od apache, nginx ne ubacuje PHP interpreter u svaki request pa mu je potreban ovaj dodatan modul. Neke od konfiguracija:
   

#### CRONTAB
