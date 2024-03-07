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
