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

### Zadatak C
