## Analiza koda projekta po izboru
### Opis projekta
 - Projekat predstavlja backend aplikaciju koristeći Spring Boot radni okvir koja omogućava korisnicima da dobiju i plate prevoz slično postojećem Uber-u. Uzimajući u obzir nedostatak današnjeg taksi prevoza, ideja je da se maksimalno olakša trasnsport korisnika uz redukciju interakcije sa prevoznikom kako bi se ceo proces ubrzao, bio konzistentniji i sigurniji.
### Članovi tima
  - Nikola Savić
  - Jovan Najdovski
  - Jovan Šerbedžija
### Pronađeni defekti

  #### Statička automatska analiza koda (github)
  - Potencijalan gubitak podataka pri dodeljivanju vrednosti korišćenjem operatora += (x+=y umesto x=x+y). Ako je tip x "uži" od tipa y, korišćenjem += rezultat postaje užeg tipa što može dovesti do gubitka podataka 
    ```
    if(workHours.getEndTime()!=null)
        sum+=ChronoUnit.MINUTES.between(workHours.getStartTime(), workHours.getEndTime());
    else
        sum+=ChronoUnit.MINUTES.between(workHours.getStartTime(), LocalDateTime.now());
  - Isključena zaštita od CSRF (Cross-Site Request Forgery) napada što može dovesti do toga da treća strana iskoristi autorizaciju korisnika i izvrši izmenjen zahtev umesto njega. Rešava se uvođenjem CSRF tokena (uglavnom jednokratan) ili proveravanjem odakle je stigao zahtev.
    ```
    public SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception {
        http.cors().and().csrf().disable()
  #### Ručna analiza koda

  - Validacija
    - Nedostatak zaštite od SQL injection i XSS napada na nivou svih kontrolera. Korišćenjem sanitizacije, filtriranja i validacije ovi problemi bi se rešili. 
    - Validacija polja kao što su email adresa, lozinka, adresa, broj telefona je ranjiva na specijalne karaktere i može biti meta XSS i SQL injection napada. Ovo se može sprečiti poboljšavanjem regularnih izraza i dodatnim proverama.
    - Pri validaciji id polja, nedostaje gornja granica što može rezultovati problemima sa performansama i servera i baze podataka usled slanja ogromnih vrednosti. Ovo se može sprečiti dodavanjem gornje granice u konsultaciji sa ostalim učesnicima u razvoju softvera.
      
  - Obrada slika
    - Određivanje formata slike se vrši na osnovu prvog bajta u nizu, što nije pouzdano. Bolje rešenje je korišćenjem određene biblioteke poput `Apache Commons Imaging`.
      ```
      if ((int)imageBytes[0]==0){
            formatType = "png";
        } else {
            formatType = "jpeg";
        }
      ```
    - Pri konverziji Base64 Stringa u niz bajtova koeišćenjem `ByteArrayOutputStream` klase, tok podataka nije zatvoren što može rezultovati curenjem memorije. Rešava se zatvaranjem toka.
    - Napadač bi mogao da izvrši DoS napad slanjem velikih Base64 stringova gde bi dekodiranje trajalo dugo. Promena koraka pri validaciji stringa bi bila od pomoći.
      ```
      public static void checkProfilePictureValidity(String profilePicture) throws DataFormatException {
        if (!profilePicture.startsWith("data:image/jpeg;base64,") && !profilePicture.startsWith("data:image/png;base64,")){
            throw new DataFormatException("Field profilePicture format is not valid!");
        }

        String imageBase64 = profilePicture.split(",")[1];
        if (!imageBase64.matches("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$")){
            throw new DataFormatException("Field profilePicture format is not valid!");
        }

        if (ImageConverter.decodeToImage(profilePicture).length>1048576){
            throw new DataFormatException("Field profilePicture cannot be longer than 786432 characters (1mb)!");
        }
      }
  - WebSockets
    - Neodstatak autorizacije pri pretplaćivanju na WebSocket može rezultovati da neko prima podatke koji nisu njemu namenjeni
    - WebSocket endpointi nisu zaštićeni od CSRF napada
    - U slučaju klijentske web aplikacije, pristup je dozvolejen samo sa adrese localhost:4200, dok je pristup android aplikacija dozvoljen sa svih adresa.
      ```
      registry.addEndpoint("/api/socket").setAllowedOrigins("http://localhost:4200").withSockJS();
      registry.addEndpoint("/api/socket/android").setAllowedOrigins("*").withSockJS();
      ```
    - Korišćenje WebSocket protokola umesto WebSocket Secure (WebSocket + TLS/SSL)
      
  - Ostalo
    - Metode za čitanje poruka iz baze i grupisanje i sortiranje u jednu konverzaciju, kreiranje raznih dijagrama za statistički prikaz mogu biti meta DoS napada pošto se izvršavaju iznova i iznova na korisnički klik. Ukoliko bi bilo mnogo poruka između učesnika ili podataka na grafiku izvršavanje može potrajati. Rešenje ovog problema može biti korišćenje keš memorije.
    - Pri korišćenju OpenStreetMap API-ja, ključ bi trebalo da se nalazi u keystore-ovima, promenljivima okruženja ili konfiguracionim fajlovima, a ne u kodu.
    - Ključ za potpisivanje JSON Web tokena se nalazi u kodu. Ovaj ključ bi uz drugačiju vrstu skladištenja, mogao biti dinamički (rotacija, promena na određen period).
    - Nedostatak autorizacije pri pristupu statističkim podacima.
    - Čuvanje poruka kao plain tekst, umesto enrkiptovanih vrednosti.
    - Pri proveravanju da li korisnik ima prava pristupa nekom resursu, iz zaglavlja se izdvaja JWT i iz njega izvlače uloga korisnika i id, ali se ne proverava da li je token validan


Vreme provedeno na analizi koda je oko 3h, otprilike pet puta manje nego na neuspešnoj instalaciji i konfiguraciji ostalih alata za statičku analizu koda poput SonarQube-a.
