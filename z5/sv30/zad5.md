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
   - Ubuntu 22.04.4 LTS je trenutno stabilna verzija. Za sve LTS verzije garantovano je minimum 5 godina standardnog bezbednosnog održavanja (standard security maintenance) na sve pakete iz ‘Main’ repozitorijuma. Za LTS 22.04 datum isteka standardnog support-a je April 2027, dok je istek legacy support-a April 2034. (https://ubuntu.com/about/release-cycle)
     
   ![Screenshot from 2024-05-16 23-06-47](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/3ea9c092-e68f-4df0-a531-be2b833df872)
   
  #### KERNEL
   - Ova verzija kernela je dosta nova pa nema naznaka da je potreban upgrade vezano za ovu distribuciju. Trenutno nije pronadjen bezbednosni vulnerability za ovu verziju.  

![Screenshot from 2024-05-16 23-17-05](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/7aebb7c4-45a7-46c1-832c-4e5d1326f738)

#### TIME MANAGEMENT
   - Ovde možemo videti da je vremenska zona podešena na UTC. Naš sistem ne radi sa mnogo vremenski senzitivnim procesima pa dodatne konfiguracije nisu potrebne.

![Screenshot from 2024-05-16 23-30-47](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/c1767bac-9ddf-4b8b-9da9-4744f458311a)

  #### PACKAGES INSTALLED
   - Većina paketa je instalirana prilikom same instalacije sistema. Pošto se radi o server a ne o desktop verziji Ubuntu-a nemamo nepotrebnih paketa vezanih za GUI. Jedan od paketa koji su sada nepotrebni jeste cloud-init. Ovaj paket je bio neophodan samo prilikom konfiguracije za instaliranje na virtuelnoj mašini. Pošto su pored paketa vezanih za minimalnu distribuciju sistema korišćeni i standradni paketi za nginx, mysql i php-fpm (PHP fastCGI process manager) nismo pronašli ranjivosti.
     
![Screenshot from 2024-05-16 23-36-12](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/9f60a227-8fba-4e0b-a889-88ca7e133d70)
   
   #### LOGGING
   - Za logovanje na našem sistemu koristi se rsyslog
   - Konfiguracija za upravljanje fajlovima i direktorijumima je ograničena permisijama: 

![Screenshot from 2024-05-17 00-46-43](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/6ba043a2-f125-49ad-9eca-6677a12b8862)

![Screenshot from 2024-05-17 00-49-46](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/5afb55ec-f360-4d6e-b5b3-c811960cb83a)

### NETWORK REVIEW

  #### GENERAL INFORMATION
   - Redom smo istražili mrežne interfejse, routing tabelu i DNS konfiguraciju:

![Screenshot from 2024-05-17 00-59-08](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/59d1d228-0d2a-4457-92eb-86893592a5e4)

  #### FIREWALL RULES
   - Koristeći ufw omogućili smo ulazni saobraćaj samo na portovima 80 (HTTP) i 22 (SSH). Portu 80 treba biti omogućen pristup sa svih adresa jer na njemu trči nginx server, dok bi se saobraćaj na portu 22 mogao ograničiti samo na odredjene adrese. Izlazni saobraćaj je omogućen na adresama koje su već uspostavili konekciju na server. Trebalo bi ograničiti izlazni saobraćaj servera samo na DNS servere i HTTP servere vezane za update softvera. Ručno je provereno da će se ufw pravila uvek primenjivati pri startu servera.

![Screenshot from 2024-05-17 01-09-02](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/f00fd508-34a5-4669-8e50-be8e23b97ba9)

#### IPV6
   - Sva pravila fw koja se primenjuju na IPV4 se primenjuju na IPV6, ali zbog ne korišćenja ove verzije protokola, treba se razmotriti ukidanje IPV6 konekcija.
   
### FILESYSTEM REVIEW

  #### MOUNTED PARTITIONS
   - Imamo dve glavne particije i jedan swap koji je prazan. Ne koristimo noatime u nijednom slučaju tako da ćemo uvek znati vreme pristupa.

![Screenshot from 2024-05-17 01-35-19](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/ee155ba5-2225-475e-8513-1e8575775f3c)

 #### SENSITIVE FILES
   - Sve bitne konfiguracije vezane za lozinke usera, mysql i nginx servera su zaštićene permisijama. Treba voditi računa o simboličnim linkovima kao što je /etc/mysql/my.cnf -> /etc/mysql/mysql.cnf.

![Screenshot from 2024-05-17 01-53-22](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/988a8b90-aaf7-4395-8801-0b19e7ea0ff9)
   
   #### SETUID
   - Svi fajlovi iz ove liste imaju dobro podešene permisije.

![Screenshot from 2024-05-17 02-00-16](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/d0a1f3aa-70ba-4a8d-8431-5f7617b7e4c5)

   #### NORMAL FILES
   - Iako app armor modul nije toliko značajan trebalo bi obezbediti bolju permisiju za njega jer je potrebno da ga koriste samo administratori. Fajlovi vezani za veb server su dobro zaštićeni.

![Screenshot from 2024-05-17 02-06-25](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/a64a9248-3c4d-4c3d-bd3a-3146b6b24cd1)

   #### BACKUP
   - Ne koristimo posebne backup fajlove za osetljive podatke. Pored toga svi fajlovi su dobro zaštičeni pa tako i backup fajlovi.

### USERS REVIEW

  #### REVIEWING THE PASSWD FILE
   - Samo root i jovan user-i imaju shell, odnosno ostali ne mogu samostalno da izvršavaju komande na sistemu.
     
![Screenshot from 2024-05-17 02-25-55](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/453fedbe-6892-485d-b7b5-2a72f1ead8c5)

 #### REVIEWING THE SHADOW FILE
   - Pošto je `$6$` na početku hash-a lozinke, korišćen je SHA-512 algoritam. Što znači da je naša lozinka lepo hash-ovana. Možemo videti da se po default-u koristi yesscript.

![Screenshot from 2024-05-17 02-44-23](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/e4e492ae-a045-4527-9d1f-7fb67bb0ddfe)

   - Pomoću John-the-ripper alata uspeli smo da nadjemo lozinku usera za 15min. Korišćen je rockyou rečnik sa predifinisanim algoritmom SHA-512.
     
![Screenshot from 2024-05-17 03-18-27](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/0cdf98de-a6e2-46e4-8e46-17e4579fb788)


#### REVIEWING THE SUDO CONFIGURATION
   - Vidimo da su svi useri zaštićeni:
     
![Screenshot from 2024-05-17 03-12-41](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/1bfeff26-bce4-4dad-b072-5f66c983ac8f)

### SERVICES REVIEW

  #### IDENTIFYING RUNNING SERVICES
 
![Screenshot from 2024-05-17 03-21-38](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/60858e96-9299-4528-b3aa-75452a27dfb4)

 #### OPENSSH
   - Zabranjen je direktan login na root nalog sistema. Trebalo bi podesiti AllowTcpForwarding parametar u konfiguraciji na no kako bi onemogućili da se ovaj sistem koristi za pristup drugih sistema. Razmotriti promenu porta kako bi se otežali napadi.

#### MYSQL
   - Videli smo da mysql nije potpuno spolja otvoren. Bind-ovan je localhost kako bi sprečio pristup sa spolja.
   - Ako pogledamo sve usere baze, imamo example_user koji koristimo za našu aplikaciju. Pomoću alata John-The-Ripper možemo pronaći dekriptovanu lozinku, iako se koristi noviji algoritam enkripcije za mysql (*9DFF74CEA805F4A29364492BF22F82DE01005E9A).

![Screenshot from 2024-05-17 04-07-10](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/865bc5c0-8749-4acb-a1e9-9dcf3e7184d7)

   - Mozemo videti da example_user nema fajl privilegije, što znači da je preko njegovog naloga nemoguće izvršiti napade menjanjem fajlova preko mysql-a.

![Screenshot from 2024-05-17 04-12-11](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/f3c9089c-8943-4634-af9f-b6f223071954)

#### NGINX CONFIGURATION
   - Nginx se koristi kao web server i proxy. U fajlu /etc/nginx/nginx.conf možemo videti da je user koji pokrece server www-data. Ostale konfiguracije mogu se naci u conf.d ili sites-enabled folderu. Konfiguracija za todo se nalazi u sites-available a sve to je linkovano na sites-enabled.

![Screenshot from 2024-05-17 04-34-34](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/636cc66c-0ef4-480e-8c8b-943ce65b3196)
     
   - Preporuka je ipak koristiti HTPPS sa sertifikatom koji bi se nalazio na mašini. Svi pozivi ka portu 80 (HTTP) u tom slučaju trebju da se redirektuju na 443 (HTTPS). Ostale konfiguracije su minimalne i ne može se pronaći ranjivost.

![Screenshot from 2024-05-17 04-40-45](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/73e61d15-2f02-4686-b7a8-34f1fbc3ba32)

   - Fajlovi koji su dostupni preko web server-a su zažtićeni permisijama:

   ![Screenshot from 2024-05-17 04-46-16](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/b74e63d2-5831-4211-9af1-b9534c4f44b6)

#### PHP CONFIGURATION
   - Koristimo PHP fastCGI process manager modul koji nam omogućuje vezu izmedju nginx-a i PHP-a. Za razliku od apache, nginx ne ubacuje PHP interpreter u svaki request pa mu je potreban ovaj dodatan modul. Neke od konfiguracija:

   ![Screenshot from 2024-05-17 04-52-53](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/c3e52d25-2acf-4fff-a9f6-3ce05f692db5)


#### CRONTAB

![Screenshot from 2024-05-17 04-16-49](https://github.com/NikolaSavic1709/rbs-tim24/assets/100165980/f0e6a09e-7750-42ac-91e0-155ee05669ac)


