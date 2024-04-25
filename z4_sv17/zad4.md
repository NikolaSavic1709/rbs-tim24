## Analiza koda projekta po izboru

### Opis projekta

- Projekat predstavlja backend aplikaciju koristeći Spring Boot radni okvir koja omogućava korisnicima da dobiju i plate
  prevoz slično postojećem Uber-u. Uzimajući u obzir nedostatak današnjeg taksi prevoza, ideja je da se maksimalno
  olakša trasnsport korisnika uz redukciju interakcije sa prevoznikom kako bi se ceo proces ubrzao, bio konzistentniji i
  sigurniji.

### Članovi tima

- Uroš Stanić
- Nemanja Majstorović
- Lazar Magazin

### Pronađeni defekti

#### Statička automatska analiza koda (github)
- Za statičku analizu je korišćen Owasp i izgenersan report se nalaz u ovom direktorijumu.
- Ranjivosti na dependacy-jima su uglavnom uzrokovane korišćenjem zastarelih verzija biblioteka. Ovo je lako rešiv problem.

#### Ručna analiza koda

- Validacija
    - Pri validaciji id polja, nedostaje gornja granica što može rezultovati problemima sa performansama i servera i
      baze podataka usled slanja ogromnih vrednosti. Ovo se može sprečiti dodavanjem gornje granice u konsultaciji sa
      ostalim učesnicima u razvoju softvera.

- Obrada slika
    - Određivanje formata slike se vrši na osnovu prvog bajta u nizu, što nije pouzdano. Bolje rešenje je korišćenjem
      određene biblioteke poput `Apache Commons Imaging`.
      ```
      if ((int)imageBytes[0]==0){
            formatType = "png";
        } else {
            formatType = "jpeg";
        }
      ```
    - Pri konverziji Base64 Stringa u niz bajtova koeišćenjem `ByteArrayOutputStream` klase, tok podataka nije zatvoren
      što može rezultovati curenjem memorije. Rešava se zatvaranjem toka.
    - Napadač bi mogao da izvrši DoS napad slanjem velikih Base64 stringova gde bi dekodiranje trajalo dugo. Promena
      koraka pri validaciji stringa bi bila od pomoći.
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
    - Takođe, čuvanje skila u ovom formatu zahteva preveliko opterećenje pa je potrebno ograničiti veličinu slike koje se mogu uploadovati, 
      da ne bismo doš 
    - Kada se pogodi putanja /swagger servira se Swagger fajl, te bi napadač mogao videti svaki endpoint i njegove parametre. 
      Ovo se može rešiti tako što bi se sklonila ova putanja.
- WebSockets
    - Neodstatak autorizacije pri pretplaćivanju na WebSocket može rezultovati da neko prima podatke koji nisu njemu
      namenjeni
    - WebSocket endpointi nisu zaštićeni od CSRF napada
    - U slučaju klijentske web aplikacije, svima je dozvoljen pristup sa svih lokacija.
      ```
      registry.addEndpoint("/api/socket").setAllowedOrigins("http://localhost:4200").withSockJS();
      registry.addEndpoint("/api/socket/android").setAllowedOrigins("*").withSockJS();
      ```
    - Korišćenje WebSocket protokola umesto WebSocket Secure (WebSocket + TLS/SSL)
    - Prilikom pretrage koisnika sa kojima bismo podelili vožnju može doći do toga da napadač dobije email-ove i profilne svih krosnika.
      Ovo se može rešiti tako što bi se sklonio autocomplete ili dodala funkcionalnost da se dodaju kontakti na neki način.

- Ostalo
    - Metode za čitanje poruka iz baze i grupisanje i sortiranje u jednu konverzaciju, kreiranje raznih dijagrama za
      statistički prikaz mogu biti meta DoS napada pošto se izvršavaju iznova i iznova na korisnički klik. Ukoliko bi
      bilo mnogo poruka između učesnika ili podataka na grafiku izvršavanje može potrajati. Rešenje ovog problema može
      biti korišćenje keš memorije.
    - JWT se nalazi u appliaction.properties. Ovaj ključ bi uz drugačiju vrstu skladištenja, mogao biti
      dinamički (rotacija, promena na određen period). Takođe, ključ bi trebao biti složeniji.
    - Kredencijali za email i bazu podataka se takođe nalaze u application.properties. Ovo je loša praksa jer se
      kredencijali mogu lako pročitati. Bolje rešenje je korišćenje environment varijabli.
    - Čuvanje poruka kao plain tekst, umesto enrkiptovanih vrednosti.

Vreme provedeno je oko 6h, 2h na analizu i 4h na pokretanje OWASP dependency checker-a.
