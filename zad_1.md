# Rad tima 24 iz predmeta Razvoj bezbednog softver

## Vežbe 1
### Zadatak A

 1. Najkorišćeniji algoritmi za enkripciju lozinki su bcrypt, SHA-2, PBKDF2.

	 * **Bcrypt:** 
	 Bcrypt je posebno dizajniran za sigurno čuvanje lozinki. Otporniji je na brute-force i rainbow table napade od ostalih algoritama. Automatski se dodaje salt što povećava sigurnost. Bcrypt omogućava postavljanje work factor-a, odnosno broja iteracija za dobijanje heša što dodatno otežava napadačima pokušaj pronalaska odgovarajuće lozinke.
	 
	* **Porodica SHA-2 algoritma (SHA-256, SHA-512):**
	SHA-2 se češće koriste u sigurnosnim protokolima, aplikacijama, digitalnim potpisima, SSL certifikatima, HMAC-u. Algoritmi su brzi i efikasni. Nisu dizajnirani specifično za čuvanje lozinki, stoga su manje otporni na napade brute-force i rainbow table. Pri čuvanje lozinki kako bi se sprečili napadi rainbow table potrebno je koristiti salt.
	
	* **PBKDF2 (Password-Based Key Derivation Function 2):**
	PBKDF2 je algoritam za generisanje ključa iz lozinke, što lozinku čini jačom u kriptografskim protokolima. Otporniji je na napade brute-force i rainbow table zbog dodatnih parametara, poput broja iteracija. Takođe podržava salt.
	
	Bcrypt se smatra najboljim izborom za enkripciju lozinki zbog svojih specifičnih karakteristika. Odabir algoritma naravno zavisi od konkretnih sigurnosnih zahteva sistema.

 2. Konfiguracioni parametar Bcrypt algoritma je work factor (salt rounds) čija preporučena vrednost je izmedju 10 i 14 (default 12). Algoritam izvršava 2 <sup>work factor</sup> iteracija.
 3. Pouzdani provajder Bcrypt algoritma za Spring Boot framework: 
Provajder je implementiran unutar Spring Security-a i automatski se koristi kada konfigurišete `PasswordEncoder` 
> org.springframework.boot:spring-boot-starter-security

4. Najnovija verzija implementacije nema ozbiljnih ranjivosti.

5. Zahtevi
   * Koristiti Bcrypt algoritam umesto drugih heš algoritama
   * Odabrati odgovarajući radni faktor
   * Generisati nasumični salt za svaku lozinku
   * Potrebno je redovno ažuriranje svih zavisnosti, uključujući Spring Boot, Spring Security i bilo koje druge biblioteke koje se koriste
   * Konfigurisati Spring Security u skladu sa najboljim sigurnosnim praksama.

### Zadatak B
1. Logovi:
	 * **Svrha log datoteka:**
	Logovi predstavljaju osnovni mehanizam za postizanje neporecivosti, kao i za dobijanje informacija o greškama i problemima u sistemu.
	 * **Struktura log datoteke:**
	Svaka interakcija između našeg sistema i njegove spoljašnjosti treba da bude zabeležena, kao što je komunikacija sa klijentom, sa bazom podataka itd. <br> Log datoteka ne bi smela da sadrži poverljive informacije kao što su šifre, ključevi, access tokeni i hash vrednosti. <br> Svaki zapis u log datoteci bi trebalo da sadrži:
		* Timestamp akcije/operacije koja je izvršena
		* Log level -  jedna vrednost iz hijerarhije koja predstavlja ozbiljnost log zapisa (DEBUG, INFO, WARNING, ERROR, CRITICAL, ...)
  		* Podatke o izvršiocu akcije
  		* Opis događaja - pisanje objekta u bazu, objekat koji je vraćen na front...
		* Ulaznu tačku u sistem - kontroler, repozitorijum...
		* Dodatne informacije kao što su IP adresa sa koje je poslat zahtev, vreme odziva ako se prate performanse sistema...			 	 	 
	 * **Čuvanje logova:**
		* Lokalno - na uređaju na kom se izvršava aplikacija, u posebnim datotekama kojima običan korisnik ne bi trebalo da pristupi. Jednostavniji i jeftiniji pristup, pogodan za manju količinu logova kojima se ne pristupa često.
	 	* Centralizovano - na posebnom serveru ili na cloudu. Pogodnije za analizu i filtriranje logova u velikim količinama.
	 * **Logovanje u Java aplikacijama:**
	Najpoznatiji radni okviri za logovanje u Java aplikacijama su `Logback` i `Apache Log4j2`. Log4j2 predstavlja unapređenje Logback-a u smislu performansi i filtriranja logova, međutim do kraja 2021. godine (verzije do 2.14) je imao veoma ozbiljnu ranjivost (severity score 10/10). Ova ranjivost kao i mnoge manje ozbiljne su uspešno mitigirane. Stabilne verzije koje se najviše koriste su 2.19.0 i 2.20.0 koje imaju par ranjivosti manje ozbiljnosti (severity score oko 5). Logback framework je podložan XML injection ranjivosti.
	 * **Upravljanje logovima:**
	Kako datoteke sa logovima vremenom postaju sve veće i veće potrebno je definisati automatski proces rotacije logova koji predstavlja kompresiju, arhiviranje ili brisanje zastarelih logova na određen vremenski period ili nakon što datoteke dostignu određenu veličinu.



### Zadatak C
1. **Rotacija lozinki:** 
 - Implementiran je mehanizam za redovno obnavljanje lozinki, što je dobra praksa za održavanje sigurnosti. Takođe, provera vremena od poslednjeg obnavljanja prilikom svake prijave je adekvatan pristup. 
 - Dodatna sigurnosna mera sprečavanja mogućnosti obnavljanja na n (implementaciji n=5) prethodnih lozinki doprinosi boljoj zaštiti sistema.
 - Po ASVS standardu v4.0.3-2.1.10 rotacija lozinke ne bi trebalo da postoji i ovo nije ispoštovano.
 2. **HTTPS komunikacija:** 
 - Korišćenje root sertifikata za obe aplikacije (frontend i backend) je dobra praksa koja obezbeđuje bezbednu komunikaciju. 
 - Korišćenje odvojenog sertifikata za komunikaciju između baze i backend-a takođe predstavlja dobru praksu.
 - Po ASVS standardu v4.0.3-9.1.1 zabranjuje se korišćenje HTTP komunikacije ukoliko nije moguće uspostaviti HTTPS vezu što je ispoštovano.
 - Po ASVS standardu v4.0.3-9.1.2 treba koristiti aktuelne verzije algoritama, šifara i protokola. Ovo je ispoštovano.
 - Po ASVS standardu v4.0.3-9.1.3 treba verifikovati da su starije verzije onemogućene. Ovo je ispoštovano.
 3. **Dvofaktorska autentifikacija:** 
 - Zahtevanje dodatne autentifikacije putem verifikacionog koda na drugom uređaju (email ili telefon) je snažna bezbednosna mera. 
 - Po ASVS standardu v4.0.3-3.3.2 ukoliko korisnik ostaje ulogovan, periodično ga treba naterati da se uloguje ponovo po isteku sesije korišćenjem dvofaktorske autentikacije. Ovo je ispoštovano.
 4. **ReCAPTCHA za forme:** 
 - Implementacija mehanizma ReCAPTCHA-e zaštite od "spamovanja" predstavlja dobru praksu.
 -  Validacija tokena na obe strane (klijentskoj i serverskoj) je ključna za efikasnu zaštitu.
 - Po ASVS standardu v4.0.3-2.2.1  potrebno je implementirati restrikciju broja upita i CAPTCHA se navodi kao koristan alat. Ovo je ispoštovano samo za bitne forme i korišćena je reCAPTCHA v4.
 5. **Validacija podataka:** 
  - Sprečavanje Injection, XSS i Path Traversal napada su ključne bezbednosne kontrole. 
  -  Validacija podataka prema najboljim praksama doprinosi održavanju bezbednog koda. 
  -  Provera tipova datoteka i ograničavanje veličine prilikom upload-a dodatno obezbeđuje sistem.
  -  Po ASVS standardu v4.0.3-5.#.# potrebno je sanitizovati sve inpute naše aplikacije što je realizovano korišćenjem popularnih frontend framework-a koji to automatski radi i popravljaju propuste.
  6.   **Monitoring uz logove:**
  - Potrebno je razviti strategiju za efikasno praćenje logova sistema i odgovoriti na incidente u realnom vremenu.
  - Po ASVS standardu v4.0.3-7.1.1,7.1.2 aplikacije ne treba da loguje kredencijale ili druge osetljive podatke. Ovo je ispoštovano.
  - Po ASVS standardu v4.0.3-7.1.3 aplikacija treba logovati podatke relevantne za sigurnost sistema kao što su uspešna i neuspešna prijava, nevalidni inputi, pristup resursima bez adekvatnih privilegije. Ovo je ispoštovano.
  - Po ASVS standardu v4.0.3-7.1.4 log treba da sadrži podatke na osnovu kojih je moguće rekreirati i detaljno analizirati tok događaja. Ovo je ispoštovano upotrebom timestamp-a i odgovarajućeg log level-a.
  7. **IAM:**
  - Za AWS aplikaciju korišćen je IAM servis za autentifikaciju korisnika i dodele prava nad sistemom. IAM je realizovan i nad komponentama sistema ( lambda,queue,..) zarad ograničavanja pristupa drugim komponentama kojima nemaju potrebe pristupati. 
  - Po ASVS standardu v4.0.3-1.4.1 svaki element sistema treba da zahteva odgovarajuća prava pristupa da bi se mogli koristiti. Ovo je ispoštovano.
  - Po ASVS standardu v4.0.3-1.4.5 sistem treba da proverava da li korisnik ima privilegije da pristupi nekom resursu, a ne samo rolu koju korisnik ima. IAM servis ovo podrazumevano podržava.

Ovo su neke dodatne bezbednosne kontrole, a naravno koristili smo i osnove zaštite podataka.
